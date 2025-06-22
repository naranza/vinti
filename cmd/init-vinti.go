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
	if len(os.Args) < 2 {
		log.Fatalf("Usage:\n  init-vinti <what> [<config>]")
	}

	what := os.Args[1]
	var inputFile string
	if what == "client_id" {
		if len(os.Args) < 3 {
			log.Fatalf("Missing input file: Usage: vinti-init client_id <config>")
		}
		inputFile = os.Args[2]
	}

	// Load config
	absPath, err := filepath.Abs("config/config.cogo")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		os.Exit(1)
	}
	config, err := core.ConfigLoad(absPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		os.Exit(1)
	}
	
	switch what {
	case "folder":
		// init folder
		for _, folder := range []string{"_client_id", "_token"} {
			if err := command.FolderInsert(config, folder); err != nil {
				log.Fatalf("Failed to create folder %q: %v", folder, err)
			}
		}
		fmt.Println("✓ Folders initialized")
		
	case "client_id":
		// init client_id
		var client core.ClientInfo
		err := cogo.LoadConfig(inputFile, &client)
		if (err != nil) {
			log.Fatalf("Error reading client file")
			os.Exit(1)
		}
		data, err := json.MarshalIndent(client, "", "  ")
		if err != nil {
			log.Printf("Skipping client %q: failed to marshal: %v", client.ClientID, err)
			return
		}
		err = command.FileWrite(config, "_client_id",  client.ClientID, string(data))
		if  err != nil {
			log.Printf("Failed to write client %q: %v", client.ClientID, err)
		} else {
			fmt.Printf("✓ Created client_id: %s\n", client.ClientID)
		}
	default:
		log.Fatalf("Unknown what: %s. Expected 'folder' or 'client_id'", what)
	}
}

