// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package api

import (
  "os"
  "encoding/json"
  "log"
  "net/http"
	"vinti/internal/core"
	"vinti/internal/command"
)

type ApiRequest struct {
	Cmd          string `json:"cmd"`
	Folder       string `json:"folder,omitempty"`
	File         string `json:"file,omitempty"`
	Data         string `json:"data,omitempty"`
	GrantType    string `json:"grant_type,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

// Standard Response
type ApiResponse struct {
  Code  int      `json:"code"`           // e.g. "ok", "error", "invalid_token", etc.
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

func writeHttpResponse(w http.ResponseWriter, response ApiResponse) {
  w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to write JSON response: %v", err)
	}
}

func APIHandler(config *core.Config, w http.ResponseWriter, r *http.Request) {
  var request ApiRequest
	var response ApiResponse
	
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&request)
  if err != nil {
  	response.Code = http.StatusBadRequest
		response.Message = "Invalid params"
    writeHttpResponse(w, response)
    return
  }
  if !allowedCommands[request.Cmd] {
    response.Code = http.StatusBadRequest
		response.Message = "Invalid command"
    writeHttpResponse(w, response)
    return
  }

  switch request.Cmd {
  case "add":
    filename, err := command.Add(config, request.Folder, request.Data)
    if err != nil {
      response.Code = http.StatusInternalServerError
      response.Message = "Failed to add data"
    } else {
      response.Code = http.StatusOK
      response.Message = filename
      log.Printf("[add] folder=%q filename=%q", request.Folder, filename)
    }
    
  case "get":
    result, err := command.Get(config, request.Folder, request.File)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = "Cannot create folder"
  	} else {
      log.Printf("[mkd] folder=%q", request.Folder)
      response.Code = http.StatusOK
      response.Message = result
    }
  case "del":
    err := command.Del(config, request.Folder, request.File)
	  if err != nil {
  		if os.IsNotExist(err) {
  			response.Code = http.StatusNotFound
  			response.Message = "File not found"
  		} else {
  			response.Code = http.StatusInternalServerError
  			response.Message = "Failed to delete file"
  		}
  	} else {
  		log.Printf("[del] folder=%q file=%q", request.Folder, request.File)
  		response.Code = http.StatusOK
  		response.Message = "done"
    }
  case "arc":
    // to do
  case "sto":
    err := command.Sto(config, request.Folder, request.File, request.Data) 
    if err != nil {
      response.Code = http.StatusInternalServerError
      response.Message = "Failed to sto data"
    } else {
      response.Code = http.StatusOK
      response.Message = "done"
      log.Printf("[sto] folder=%q filename=%q", request.Folder, request.File)
    }
  case "all":
    files, err := command.All(config, request.Folder)
  	if err != nil {
  		response.Code = http.StatusInternalServerError
  		response.Message = "Failed to list files"
  	} else {
  		response.Code = http.StatusOK
  		response.Files = files
  		response.Message = "done"
  		log.Printf("[all] folder=%q numfiles=%d", request.Folder, len(files))
  	}
  case "mkd":
  	err := command.MakeDir(config, request.Folder)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = "Cannot create folder"
		} else {
  		log.Printf("[mkd] folder=%q", request.Folder)
  		response.Code = http.StatusOK
  		response.Message = "done"
		}
  case "ddi":
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
  writeHttpResponse(w, response)
  

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(statusCode)
	// if err := json.NewEncoder(w).Encode(response); err != nil {
	// 	log.Printf("Failed to write JSON response: %v", err)
	// }
}