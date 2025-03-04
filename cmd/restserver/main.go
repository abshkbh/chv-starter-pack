package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/abshkbh/arrakis/out/gen/serverapi"
	"github.com/abshkbh/arrakis/pkg/config"
	"github.com/abshkbh/arrakis/pkg/server"
)

type restServer struct {
	vmServer *server.Server
}

// Implement handler functions
func (s *restServer) startVM(w http.ResponseWriter, r *http.Request) {
	var req serverapi.StartVMRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	if req.GetVmName() == "" {
		http.Error(w, "Empty vm name", http.StatusBadRequest)
		return
	}

	resp, err := s.vmServer.StartVM(r.Context(), &req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to start VM: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *restServer) destroyVM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vmName := vars["name"]

	// Create request object with the VM name
	req := serverapi.VMRequest{
		VmName: &vmName,
	}

	resp, err := s.vmServer.DestroyVM(r.Context(), &req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to destroy VM: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *restServer) destroyAllVMs(w http.ResponseWriter, r *http.Request) {
	resp, err := s.vmServer.DestroyAllVMs(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to destroy all VMs: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *restServer) listAllVMs(w http.ResponseWriter, r *http.Request) {
	resp, err := s.vmServer.ListAllVMs(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to list all VMs: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *restServer) listVM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vmName := vars["name"]
	resp, err := s.vmServer.ListVM(r.Context(), vmName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to list VM: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *restServer) snapshotVM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vmName := vars["name"]

	var req struct {
		SnapshotId string `json:"snapshotId,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	resp, err := s.vmServer.SnapshotVM(r.Context(), vmName, req.SnapshotId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create snapshot: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *restServer) updateVMState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vmName := vars["name"]

	var req serverapi.VmsNamePatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	status := req.GetStatus()
	if status != "stopped" && status != "paused" && status != "resume" {
		http.Error(w, "Status must be either 'stopped' or 'paused'", http.StatusBadRequest)
		return
	}

	vmReq := serverapi.VMRequest{
		VmName: &vmName,
	}

	var resp *serverapi.VMResponse
	var err error
	if status == "stopped" {
		resp, err = s.vmServer.StopVM(r.Context(), &vmReq)
	} else if status == "paused" {
		resp, err = s.vmServer.PauseVM(r.Context(), &vmReq)
	} else { // status == "resume"
		resp, err = s.vmServer.ResumeVM(r.Context(), &vmReq)
	}

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update VM state: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *restServer) vmCommand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vmName := vars["name"]

	var req serverapi.VmCommandRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	if req.GetCmd() == "" {
		http.Error(w, "Command cannot be empty", http.StatusBadRequest)
		return
	}

	resp, err := s.vmServer.VMCommand(r.Context(), vmName, req.GetCmd())
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to execute command: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *restServer) vmFileUpload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vmName := vars["name"]

	var req serverapi.VmFileUploadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	if len(req.GetFiles()) == 0 {
		http.Error(w, "No files provided", http.StatusBadRequest)
		return
	}

	resp, err := s.vmServer.VMFileUpload(r.Context(), vmName, req.GetFiles())
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to upload files: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (s *restServer) vmFileDownload(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vmName := vars["name"]

	paths := r.URL.Query().Get("paths")
	if paths == "" {
		http.Error(w, "Missing 'paths' query parameter", http.StatusBadRequest)
		return
	}

	resp, err := s.vmServer.VMFileDownload(r.Context(), vmName, paths)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to download files: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	var serverConfig *config.ServerConfig
	var configFile string

	app := &cli.App{
		Name:  "arrakis-restserver",
		Usage: "A daemon for spawning and managing cloud-hypervisor based microVMs.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Path to config file",
				Destination: &configFile,
				Value:       "./config.yaml",
			},
		},
		Action: func(ctx *cli.Context) error {
			var err error
			serverConfig, err = config.GetServerConfig(configFile)
			if err != nil {
				return fmt.Errorf("server config not found: %v", err)
			}
			log.Infof("server config: %v", serverConfig)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.WithError(err).Fatal("server exited with error")
	}

	// At this point `serverConfig` is populated.
	// Create the VM server
	vmServer, err := server.NewServer(*serverConfig)
	if err != nil {
		log.Fatalf("failed to create VM server: %v", err)
	}

	// Create REST server
	s := &restServer{vmServer: vmServer}
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/vms", s.startVM).Methods("POST")
	r.HandleFunc("/vms/{name}", s.updateVMState).Methods("PATCH")
	r.HandleFunc("/vms/{name}", s.destroyVM).Methods("DELETE")
	r.HandleFunc("/vms", s.destroyAllVMs).Methods("DELETE")
	r.HandleFunc("/vms", s.listAllVMs).Methods("GET")
	r.HandleFunc("/vms/{name}", s.listVM).Methods("GET")
	r.HandleFunc("/vms/{name}/snapshots", s.snapshotVM).Methods("POST")
	r.HandleFunc("/vms/{name}/cmd", s.vmCommand).Methods("POST")
	r.HandleFunc("/vms/{name}/files", s.vmFileUpload).Methods("POST")
	r.HandleFunc("/vms/{name}/files", s.vmFileDownload).Methods("GET")

	// Start HTTP server
	srv := &http.Server{
		Addr:    serverConfig.Host + ":" + serverConfig.Port,
		Handler: r,
	}

	go func() {
		log.Printf("REST server listening on: %s:%s", serverConfig.Host, serverConfig.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	vmServer.DestroyAllVMs(context.Background())
	log.Println("Server stopped")
}
