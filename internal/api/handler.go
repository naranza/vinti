// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package api

import (
  "os"
  "encoding/json"
  "fmt"
  "net/http"
  "vinti/internal/core"
  "vinti/internal/command"
  "github.com/naranza/bagolo"
  vlog "vinti/internal/log"
)

var allowedCommands = map[string]bool{
  "fi-get": true, "fi-set": true, "fi-del": true, "fi-ren": true, "fi-lst": true, 
  "da-ins": true,
  "fo-ins": true,
  "to-req": true,
}

func writeHttpResponse(w http.ResponseWriter, response core.ApiResponse) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(response.Code)
  if err := json.NewEncoder(w).Encode(response); err != nil {
      vlog.Log(vlog.INFO, "Failed to write JSON response: %v", err)
  }
}

func LogAndSendResponse(
  w http.ResponseWriter, 
  response core.ApiResponse, 
  user core.ApiUser, 
  level int, 
  message string){
  vlog.Log(level, "%s %d %s", user.Username, response.Code, message)
  writeHttpResponse(w, response)
}

func APIHandler(config *core.Config, w http.ResponseWriter, r *http.Request) {
  var request core.ApiRequest
  var response core.ApiResponse
  var user core.ApiUser
  var logLevel int
  var logMessage string

  response.Code = http.StatusOK
  response.Message = "ok"

  user.Username = "vinti"
  user.Role = "system"

  logLevel = vlog.INFO
  logMessage = ""

  // Authorisation and authentication
  username, password, err := bagolo.Auth(r)
  if err != nil {
    response.Code = http.StatusBadRequest
    response.Message = "Authorization header issue"
  } else if username == "" || password == "" {
    response.Code = http.StatusUnauthorized
    response.Message = "Invalid credential"
  } else {
    userData, err := command.FileRead(config, "_user", username)
    if err != nil {
      response.Code = http.StatusInternalServerError
      response.Message = "User error"
    } else {
      err = json.Unmarshal([]byte(userData), &user)
      if err != nil {
        response.Code = http.StatusInternalServerError
        response.Message = "User decode error"
      } else if username != user.Username || password != user.Password {
        response.Code = http.StatusUnauthorized
        response.Message = "Invalid credential"
      }
    }
  }
  if response.Code != http.StatusOK {
    LogAndSendResponse(w, response, user, logLevel, response.Message)
    return
  }  
  
  // request decode
  decoder := json.NewDecoder(r.Body)
  err = decoder.Decode(&request)
  if err != nil {
    response.Code = http.StatusBadRequest
    response.Message = "Invalid request"
    return
  } else  if !allowedCommands[request.Cmd] {
    response.Code = http.StatusBadRequest
    response.Message = "Invalid cmd"
  }
  if response.Code != http.StatusOK {
    LogAndSendResponse(w, response, user, logLevel, response.Message)
    return
  }  

  // command handling
  switch request.Cmd {
  case "fo-ins":
    err := command.FolderInsert(config, request.Folder)
    if err != nil {
      response.Code = http.StatusInternalServerError
      response.Message = "Cannot create folder"
    } else {
      response.Code = http.StatusOK
      response.Message = "done"
    }
    logMessage = fmt.Sprintf("%s %s", request.Cmd, request.Folder)
  case "da-ins":
    filename, err := command.DataInsert(config, request.Folder, request.Data)
    if err != nil {
      response.Code = http.StatusInternalServerError
      response.Message = "Failed to add data"
    } else {
      response.Code = http.StatusOK
      response.Message = filename
    }
    logMessage = fmt.Sprintf("%s %s %s", request.Cmd, request.Folder, filename)
  case "fi-get":
    result, err := command.FileRead(config, request.Folder, request.File)
    if err != nil {
      response.Code = http.StatusInternalServerError
      response.Message = "Cannot read file"
    } else {
      response.Code = http.StatusOK
      response.Message = result
    }
    logMessage = fmt.Sprintf("%s %s %s", request.Cmd, request.Folder, request.File)
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
      logMessage = response.Message
    } else {
      response.Code = http.StatusOK
      response.Message = "done"
    }
    logMessage = fmt.Sprintf("%s %s", request.Cmd, request.File)
  case "fi-ren":
    err := command.FileRename(config, request.Folder, request.File, request.To) 
    if err != nil {
    response.Code = http.StatusInternalServerError
    response.Message = "Failed to rename file"
    } else {
      response.Code = http.StatusOK
      response.Message = "done"
    }
    logMessage = fmt.Sprintf("%s %s %s %s", request.Cmd, request.Folder, request.File, request.To)
  case "fi-set":
    err := command.FileWrite(config, request.Folder, request.File, request.Data) 
    if err != nil {
    response.Code = http.StatusInternalServerError
    response.Message = "Failed to store data"
    } else {
      response.Code = http.StatusOK
      response.Message = "done"
    }
    logMessage = fmt.Sprintf("%s %s %s", request.Cmd, request.Folder, request.File)
  case "fi-lst":
    files, err := command.FileList(config, request.Folder)
    if err != nil {
      response.Code = http.StatusInternalServerError
      response.Message = "Failed to list files"
    } else {
      response.Code = http.StatusOK
      response.Files = files
      response.Message = "done"
    }
    logMessage = fmt.Sprintf("%s %s %d", request.Cmd, request.Folder, len(response.Files))
  default:
    response.Code = http.StatusInternalServerError
    response.Message = "Developers left something behind"
  }
  LogAndSendResponse(w, response, user, logLevel, logMessage)
}