package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	ifname = "eth0"
	ipBin  = "/usr/bin/ip"
)

// parseKeyFromCmdLine parses a key from the kernel command line. Assumes each
// key:val is present like key="val" in /proc/cmdline.
func parseKeyFromCmdLine(prefix string) (string, error) {
	cmdline, err := os.ReadFile("/proc/cmdline")
	if err != nil {
		return "", fmt.Errorf("failed to read /proc/cmdline: %w", err)
	}

	key := prefix + "="
	cmdlineStr := string(cmdline)

	start := strings.Index(cmdlineStr, key)
	if start == -1 {
		return "", fmt.Errorf("key %q not found in kernel command line", key)
	}

	start += len(key)
	value := strings.TrimPrefix(cmdlineStr[start:], "\"")
	end := strings.IndexByte(value, '"')
	if end == -1 {
		return "", fmt.Errorf("unclosed quote for key %q in kernel command line", key)
	}
	return value[:end], nil
}

// parseNetworkingMetadata parses the networking metadata from the kernel command line.
func parseNetworkingMetadata() (string, string, error) {
	guestCIDR, err := parseKeyFromCmdLine("guest_ip")
	if err != nil {
		return "", "", fmt.Errorf("failed to parse guest_ip: %w", err)
	}

	gatewayCIDR, err := parseKeyFromCmdLine("gateway_ip")
	if err != nil {
		return "", "", fmt.Errorf("failed to parse gateway_ip: %w", err)
	}

	if guestCIDR == "" || gatewayCIDR == "" {
		return "", "", fmt.Errorf("guest_ip or gateway_ip not found in kernel command line")
	}

	// gateway's IP needs to be returned without the subnet mask.
	gatewayIP, _, err := net.ParseCIDR(gatewayCIDR)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse gatewayCIDR: %w", err)
	}

	return guestCIDR, gatewayIP.String(), nil
}

// setupNetworking sets up networking inside the guest.
func setupNetworking(guestCIDR string, gatewayIP string) error {
	cmd := exec.Command(ipBin, "l", "set", "lo", "up")
	err := cmd.Run()
	if err != nil {
		log.WithError(err).Fatal("failed to set the lo interface up")
	}

	cmd = exec.Command(ipBin, "a", "add", guestCIDR, "dev", ifname)
	err = cmd.Run()
	if err != nil {
		log.WithError(err).Fatal("failed to add IP address to interface")
	}

	cmd = exec.Command(ipBin, "l", "set", ifname, "up")
	err = cmd.Run()
	if err != nil {
		log.WithError(err).Fatal("failed to set interface up")
	}

	cmd = exec.Command(ipBin, "r", "add", "default", "via", gatewayIP, "dev", ifname)
	err = cmd.Run()
	if err != nil {
		log.WithError(err).Fatal("failed to add default route")
	}

	f, err := os.OpenFile("/etc/resolv.conf", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.WithError(err).Fatal("failed to open /etc/resolv.conf")
	}
	defer f.Close()

	_, err = f.WriteString("nameserver 8.8.8.8\n")
	if err != nil {
		log.WithError(err).Fatal("failed to write nameserver to /etc/resolv.conf")
	}
	return nil
}

func main() {
	log.Infof("starting guestinit")
	err := os.WriteFile("/etc/hostname", []byte("chv-vm"), 0644)
	if err != nil {
		log.WithError(err).Fatal("failed to write hostname")
	}

	guestCIDR, gatewayIP, err := parseNetworkingMetadata()
	if err != nil {
		log.WithError(err).Fatal("failed to parse guest networking metadata")
	}

	// Setup networking.
	err = setupNetworking(guestCIDR, gatewayIP)
	if err != nil {
		log.WithError(err).Fatal("failed to setup networking")
	}

	log.Info("guestinit exiting...")
}
