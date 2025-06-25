// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package main

import (
  "github.com/naranza/cogo"
  "encoding/json"
  "fmt"
  "log"
  "os"
  "path/filepath"
  "vinti/internal/command"
  "vinti/internal/core"
)

func main() {
  if len(os.Args) < 3 {
    fmt.Println("Usage:\n  init-vinti <cmd> <subject>")
    os.Exit(1)
  }

  cmd := os.Args[1]
  subject := os.Args[2]
  
  // Load config
  absPath, err := filepath.Abs("config/config.cogo")
  if err != nil {
    fmt.Println("Failed to load config: %v", err)
    os.Exit(1)
  }
  config, err := core.ConfigLoad(absPath)
  if err != nil {
    fmt.Println("Failed to load config: %v", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "folder":
    // init folder
    err := command.FolderInsert(config, subject); 
    if err != nil {
      fmt.Println("Failed to create folder %q: %v", subject, err)
    } else {
      fmt.Println("Folders initialized")
    }
      
    
  case "user":
    // init client_id
    var client core.ClientInfo
    err := cogo.LoadConfig(subject, &client)
    if (err != nil) {
      fmt.Println("Error reading client file")
      os.Exit(1)
    }
    data, err := json.MarshalIndent(client, "", "  ")
    if err != nil {
      log.Printf("Skipping client %q: failed to marshal: %v", client.Username, err)
      return
    }
    err = command.FileWrite(config, "_client_id",  client.Username, string(data))
    if  err != nil {
      log.Printf("Failed to write client %q: %v", client.Username, err)
    } else {
      fmt.Printf("✓ Created client_id: %s\n", client.Username)
    }
  default:
    fmt.Println("Unknown command: %s. Expected 'folder' or 'user'", cmd)
  }
}

