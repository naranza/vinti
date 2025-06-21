// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package api

import (
  "encoding/json"
  "log"
  "net/http"
	"vinti/internal/core"
)

// Common base request struct
type Request struct {
  Cmd    string          `json:"cmd"`
  Params json.RawMessage `json:"params"` // to unmarshal dynamically per command
}

// Add Command Params
type FolderDataParams struct {
  Folder string `json:"folder"`
  Data   string `json:"data"`
}

// Get Command Params
// Del Command Params
// Arc Command Params
type FolderFileParams struct {
  Folder string `json:"folder"`
  File   string `json:"file"`
}

// Sto Command Params
type FolderFileDataParams struct {
  Folder string `json:"folder"`
  File   string `json:"file"`
  Data   string `json:"data"`
}

// MkD Command Params
// All Command Params
// Ddi Command Params
type FolderParams struct {
  Folder string `json:"folder"`
}

// OAuth2 Token Command Params (02t)
type O2tParams struct {
  GrantType    string `json:"grant_type"`
  ClientID     string `json:"client_id"`
  ClientSecret string `json:"client_secret"`
  Scope        string `json:"scope"`
}

// Standard Response
type Response struct {
    Status  string      `json:"status"`           // e.g. "ok", "error", "invalid_token", etc.
    Message string      `json:"message"`          // description or content
    Files   []string    `json:"files,omitempty"`  // used for 'all' command
    O2      *O2tResponse `json:"o2,omitempty"`    // used for oauth token response
}

// OAuth2 Token Response payload
type O2tResponse struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
    ExpiresIn   int    `json:"expires_in"`
    Scope       string `json:"scope"`
}

var allowedCommands = map[string]bool{
  "add": true,
  "get": true,
  "arc": true,
  "del": true,
  "sto": true,
  "all": true,
  "o2t": true,
  "mkd": true,
  "ddi": true,
}

func APIHandler(config *core.Config, w http.ResponseWriter, r *http.Request) {
  var req Request
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&req); err != nil {
    writeError(w, http.StatusBadRequest, "Invalid JSON")
    return
  }

  if !allowedCommands[req.Cmd] {
    writeError(w, http.StatusBadRequest, "Invalid command")
    return
  }

  switch req.Cmd {
  case "add":
    var params FolderDataParams
    if err := json.Unmarshal(req.Params, &params); err != nil {
      writeError(w, http.StatusBadRequest, "Invalid params for add")
      return
    }
    filename, err := core.Add(config, params.Folder, params.Data)
	  if err != nil {
	    writeError(w, http.StatusInternalServerError, "Failed to add data: "+err.Error())
	  } else {
			log.Printf("[add] folder=%q filename=%q", params.Folder, filename)
			writeOK(w, filename)
		}
  case "get", "del", "arc":
    var params FolderFileParams
    if err := json.Unmarshal(req.Params, &params); err != nil {
      writeError(w, http.StatusBadRequest, "Invalid params for "+req.Cmd)
      return
    }
    // TODO: implement get/del/arc logic here
    log.Printf("[%s] folder=%q file=%q", req.Cmd, params.Folder, params.File)

    writeOK(w, req.Cmd+" executed successfully")

  case "sto":
    var params FolderFileDataParams
    if err := json.Unmarshal(req.Params, &params); err != nil {
      writeError(w, http.StatusBadRequest, "Invalid params for sto")
      return
    }
    // TODO: implement store logic here
    log.Printf("[sto] folder=%q file=%q data=%q", params.Folder, params.File, params.Data)

    writeOK(w, "Data stored successfully")

  case "all", "mkd", "ddi":
    var params FolderParams
    if err := json.Unmarshal(req.Params, &params); err != nil {
      writeError(w, http.StatusBadRequest, "Invalid params for "+req.Cmd)
      return
    }
    // TODO: implement all/mkd/ddi logic here
    log.Printf("[%s] folder=%q", req.Cmd, params.Folder)

    // Example for "all": return some dummy file list
    if req.Cmd == "all" {
      resp := Response{
        Status:  "ok",
        Message: "",
        Files:   []string{"file1.txt", "file2.txt", "file3.txt"},
      }
      writeJSON(w, resp)
      return
    }

    writeOK(w, req.Cmd+" executed successfully")

  case "o2t":
    var params O2tParams
    if err := json.Unmarshal(req.Params, &params); err != nil {
      writeError(w, http.StatusBadRequest, "Invalid params for o2t")
      return
    }
    // TODO: implement OAuth2 token logic here
    log.Printf("[o2t] client_id=%q scope=%q", params.ClientID, params.Scope)

    // Dummy token response example
    resp := Response{
      Status:  "ok",
      Message: "",
      O2: &O2tResponse{
        AccessToken: "dummy-access-token",
        TokenType:   "Bearer",
        ExpiresIn:   3600,
        Scope:       params.Scope,
      },
    }
    writeJSON(w, resp)

  default:
    writeError(w, http.StatusNotImplemented, "Command not implemented")
  }
}

// Helper to write JSON OK response with message
func writeOK(w http.ResponseWriter, msg string) {
  writeJSON(w, Response{
    Status:  "ok",
    Message: msg,
  })
}

// Helper to write JSON error response
func writeError(w http.ResponseWriter, statusCode int, msg string) {
  w.WriteHeader(statusCode)
  writeJSON(w, Response{
    Status:  "error",
    Message: msg,
  })
}

// Helper to write any JSON response
func writeJSON(w http.ResponseWriter, v interface{}) {
  w.Header().Set("Content-Type", "application/json")
  if err := json.NewEncoder(w).Encode(v); err != nil {
    log.Printf("Failed to write JSON response: %v", err)
  }
}