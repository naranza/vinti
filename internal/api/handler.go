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

var allowedCommands = map[string]bool{
  "fo-lst": true, 
  "fi-get": true, "fi-set": true, "fi-del": true, "fi-arc": true, "fi-lst": true, 
  "da-ins": true,
  "di-ins": true, "di-del": true,
  "ci-get": true, "ci-set": true, "ci-del": true, "ci-arc": true, "ci-una": true,
  "to-req": true,
}

func writeHttpResponse(w http.ResponseWriter, response core.ApiResponse) {
  w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to write JSON response: %v", err)
	}
}

func APIHandler(config *core.Config, w http.ResponseWriter, r *http.Request) {
  var request core.ApiRequest
	var response core.ApiResponse
	
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
  case "da-ins":
    filename, err := command.DataInsert(config, request.Folder, request.Data)
    if err != nil {
      response.Code = http.StatusInternalServerError
      response.Message = "Failed to add data"
    } else {
      response.Code = http.StatusOK
      response.Message = filename
      log.Printf("[ida] folder=%q filename=%q", request.Folder, filename)
    }
    
  case "fi-get":
    result, err := command.FileRead(config, request.Folder, request.File)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = "Cannot create folder"
  	} else {
      log.Printf("[mkd] folder=%q", request.Folder)
      response.Code = http.StatusOK
      response.Message = result
    }
  case "fi-del":
    err := command.FileDelete(config, request.Folder, request.File)
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
  case "fi-arc":
    // to do
  case "fi-set":
    err := command.FileWrite(config, request.Folder, request.File, request.Data) 
    if err != nil {
      response.Code = http.StatusInternalServerError
      response.Message = "Failed to sto data"
    } else {
      response.Code = http.StatusOK
      response.Message = "done"
      log.Printf("[sto] folder=%q filename=%q", request.Folder, request.File)
    }
  case "fi-lst":
    files, err := command.FileList(config, request.Folder)
  	if err != nil {
  		response.Code = http.StatusInternalServerError
  		response.Message = "Failed to list files"
  	} else {
  		response.Code = http.StatusOK
  		response.Files = files
  		response.Message = "done"
  		log.Printf("[all] folder=%q numfiles=%d", request.Folder, len(files))
  	}
  case "di-ins":
  	err := command.FolderInsert(config, request.Folder)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = "Cannot create folder"
		} else {
  		log.Printf("[mkd] folder=%q", request.Folder)
  		response.Code = http.StatusOK
  		response.Message = "done"
		}
  case "ci-set":
    if request.ClientID == "" || request.ClientSecret == "" || request.Role == "" {
      response.Code = http.StatusBadRequest
      response.Message = "Missing fields for aci"
      break
    }

    client := map[string]string{
      "client_id": request.ClientID,
      "client_secret": request.ClientSecret,
      "Role": request.Role,
    }
    data, _ := json.Marshal(client)
    err := command.FileWrite(config, "_client_id", request.ClientID, string(data))
    if err != nil {
      response.Code = http.StatusInternalServerError
      response.Message = "Failed to store client"
    } else {
      response.Code = http.StatusOK
      response.Message = "done"
      log.Printf("[aci] stored client_id=%q", request.ClientID)
    }
 case "to-req":
    // token-request

    accessToken, err := command.TokenRequest(config, request.ClientID)
    if err != nil {
      response.Code = http.StatusInternalServerError
      response.Message = "Failed to generate token"
    } else {
      response.Code = http.StatusOK
  		response.Message = "done"
      response.AccessToken = accessToken
  		response.TokenType = "Bearer"
  		response.ExpiresIn = config.TokenExpiresIn
      // If no error, response is already filled by O2t
      log.Printf("[to-req] client_id=%q token=%q", request.ClientID, response.AccessToken)
    }

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