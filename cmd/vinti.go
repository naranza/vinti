// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package main

import (
    "fmt"
    "log"
    "net/http"
    // "os"
    "vinti/internal/core"
    "vinti/internal/api"
)

var config string

func main() {
  config, err := core.ConfigLoad("config/config.cogo")
  if err != nil {
      log.Fatalf("Failed to load config: %v", err)
  }
  fmt.Println(config)

  // Wrap the handler so config can be passed
  http.HandleFunc("/run", func(w http.ResponseWriter, r *http.Request) {
      api.APIHandler(config, w, r)
  })

  log.Println("API server running on :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
