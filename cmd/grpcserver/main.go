package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/abshkbh/chv-lambda/out/protos"
	"github.com/abshkbh/chv-lambda/pkg/server"
	"google.golang.org/grpc"
)

const (
	numBootVcpus    = 1
	memorySizeBytes = 512 * 1024 * 1024
	// Case sensitive.
	serialPortMode = "Tty"
	// Case sensitive.
	consolePortMode = "Off"
	chvBinPath      = "/home/maverick/projects/chv-lambda/resources/bin/cloud-hypervisor"

	bridgeName              = "br0"
	bridgeIP                = "10.20.1.1/24"
	bridgeSubnet            = "10.20.1.0/24"
	numNetDeviceQueues      = 2
	netDeviceQueueSizeBytes = 256
	netDeviceId             = "_net0"
	stateDir                = "/run/chv-lambda"
	reapVmTimeout           = 20 * time.Second
)

var (
	kernelPath = "resources/bin/compiled-vmlinux.bin"
	rootfsPath = "out/ubuntu-ext4.img"
	initPath   = "/usr/bin/tini -- /opt/custom_scripts/guestinit"
)

func main() {
	apiServer, err := server.NewServer(stateDir, bridgeName, bridgeIP, bridgeSubnet)
	if err != nil {
		log.WithError(err).Fatalf("failed to create api server")
	}
	lis, err := net.Listen("tcp", "127.0.0.1:6000")
	if err != nil {
		log.WithError(err).Fatalf("failed to listen")
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	s := grpc.NewServer()

	// Set up signal handler.
	go func() {
		log.Infof("waiting for signal")
		sig := <-sigChan
		log.Infof("received signal: %v", sig)
		s.GracefulStop()
		log.Infof("gracefully stopped")
	}()

	protos.RegisterVMManagementServiceServer(s, apiServer)
	log.Printf("server Pid:%d listening at %v", os.Getpid(), lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.WithError(err).Fatalf("failed to serve")
	}
	apiServer.DestroyAllVMs(context.Background(), &protos.DestroyAllVMsRequest{})
	log.WithError(err).Info("server exited")
}