// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package api

import (
  "encoding/json"
  "log"
  "net/http"
	"vinti/internal/core"
	"vinti/internal/command"
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

func writeHttpResponse(w http.ResponseWriter, statusCode int, response interface{}) {
  w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to write JSON response: %v", err)
	}
}

func APIHandler(config *core.Config, w http.ResponseWriter, r *http.Request) {
  var req Request
	var response Response
	
  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&req); err != nil {
  	response.Status = "error"
		response.Message = "Invalid params"
    writeHttpResponse(w, http.StatusBadRequest, response)
    return
  }
  if !allowedCommands[req.Cmd] {
    response.Status = "error"
		response.Message = "Invalid command"
    writeHttpResponse(w, http.StatusBadRequest, response)
    return
  }

  log.Printf("Received request: Cmd=%q Params=%s", req.Cmd, string(req.Params))
  
  switch req.Cmd {
  case "add":
    var params FolderDataParams
    if err := json.Unmarshal(req.Params, &params); err != nil {
			response.Status = "error"
			response.Message = "Invalid params"
      writeHttpResponse(w, http.StatusBadRequest, response)
      return
		}
    
		filename, err := command.Add(config, params.Folder, params.Data)
		if err != nil {
			response.Status = "error"
			response.Message = "Failed to add data"
      writeHttpResponse(w, http.StatusInternalServerError, response)
      return
		}

	  response.Status = "ok"
	  response.Message = filename
    writeHttpResponse(w, http.StatusOK, response)
		log.Printf("[add] folder=%q filename=%q", params.Folder, filename)
    
  case "get", "del", "arc":
    // var params FolderFileParams
    // if err := json.Unmarshal(req.Params, &params); err != nil {
    //   writeError(w, http.StatusBadRequest, "Invalid params for "+req.Cmd)
    //   return
    // }
    // // TODO: implement get/del/arc logic here
    // log.Printf("[%s] folder=%q file=%q", req.Cmd, params.Folder, params.File)

    // writeOK(w, req.Cmd+" executed successfully")

  case "sto":
    // var params FolderFileDataParams
    // if err := json.Unmarshal(req.Params, &params); err != nil {
    //   writeError(w, http.StatusBadRequest, "Invalid params for sto")
    //   return
    // }
    // // TODO: implement store logic here
    // log.Printf("[sto] folder=%q file=%q data=%q", params.Folder, params.File, params.Data)

    // writeOK(w, "Data stored successfully")

  case "all", "mkd", "ddi":
    var params FolderParams
    if err := json.Unmarshal(req.Params, &params); err != nil {
			response.Status = "error"
			response.Message = "Invalid params"
      writeHttpResponse(w, http.StatusBadRequest, response)
      return
    }
		if req.Cmd == "mkd" {
			err := command.MakeDir(config, params.Folder)
			if err != nil {
				response.Status = "error"
				response.Message = "Cannot create folder"
        writeHttpResponse(w, http.StatusInternalServerError, response)
        return
  		}
			log.Printf("[mkd] folder=%q", params.Folder)
			response.Status = "ok"
			response.Message = "created"
      writeHttpResponse(w, http.StatusOK, response)
		}
    // // Example for "all": return some dummy file list
    // if req.Cmd == "all" {
    //   resp := Response{
    //     Status:  "ok",
    //     Message: "",
    //     Files:   []string{"file1.txt", "file2.txt", "file3.txt"},
    //   }
    //   writeJSON(w, resp)
    //   return
    // }

    // writeOK(w, req.Cmd+" executed successfully")

  case "o2t":
    // var params O2tParams
    // if err := json.Unmarshal(req.Params, &params); err != nil {
    //   writeError(w, http.StatusBadRequest, "Invalid params for o2t")
    //   return
    // }
    // // TODO: implement OAuth2 token logic here
    // log.Printf("[o2t] client_id=%q scope=%q", params.ClientID, params.Scope)

    // // Dummy token response example
    // resp := Response{
    //   Status:  "ok",
    //   Message: "",
    //   O2: &O2tResponse{
    //     AccessToken: "dummy-access-token",
    //     TokenType:   "Bearer",
    //     ExpiresIn:   3600,
    //     Scope:       params.Scope,
    //   },
    // }
    // writeJSON(w, resp)

  default:
    // writeError(w, http.StatusNotImplemented, "Command not implemented")
  }
  
  

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(statusCode)
	// if err := json.NewEncoder(w).Encode(response); err != nil {
	// 	log.Printf("Failed to write JSON response: %v", err)
	// }
}
