// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package main

import (
    "fmt"
    "log"
    "net/http"
    // "os"
    "path/filepath"
    "vinti/internal/core"
    "vinti/internal/api"
    vlog "vinti/internal/log"
)

var config string

func main() {
  // Load config
  absPath, err := filepath.Abs("config/config.cogo")
  if err != nil {
    log.Fatalf("Failed to load config: %v", err)
  }
  config, err := core.ConfigLoad(absPath)
  if err != nil {
    log.Fatalf("Failed to load config: %v", err)
  }
  
  // Init Vinti Log
  err = vlog.Init(config.LogThreshold);
  
  http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
    api.APIHandler(config, w, r)
  })

  addr := fmt.Sprintf(":%d", config.ServerPort)
  if config.TlsCertPath != "" && config.TlsKeyPath != "" {
    vlog.Log(vlog.INFO, "Starting SSL Vinti server on port %d", config.ServerPort)
    err = http.ListenAndServeTLS(addr, config.TlsCertPath, config.TlsKeyPath, nil)
  } else {
    vlog.Log(vlog.INFO, "Starting Vinti server on port %d", config.ServerPort)
    err = http.ListenAndServe(addr, nil)
  }

  if err != nil {
    log.Fatalf("Server failed: %v", err)
  }
}
