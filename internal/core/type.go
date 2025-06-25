// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package core

type ApiRequest struct {
  Cmd string `json:"cmd"`
  Folder string `json:"folder,omitempty"`
  File string `json:"file,omitempty"`
  Data string `json:"data,omitempty"`
  GrantType string `json:"grant_type,omitempty"`
  Username string `json:"client_id,omitempty"`
  Password string `json:"client_secret,omitempty"`
  Role string `json:"scope,omitempty"`
}

type ApiResponse struct {
  Code int `json:"code"`
  Message string `json:"message"`
  Files []string `json:"files,omitempty"`
}

type ClientInfo struct {
  Username string `json:"client_id"`
  Password string `json:"client_secret"`
  Role string `json:"scope"`
}