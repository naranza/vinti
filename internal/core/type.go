// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package core

type ApiRequest struct {
  Cmd string `json:"cmd"`
  Folder string `json:"folder,omitempty"`
  File string `json:"file,omitempty"`
  Data string `json:"data,omitempty"`
  Role string `json:"scope,omitempty"`
}

type ApiResponse struct {
  Code int `json:"code"`
  Message string `json:"message"`
  Files []string `json:"files,omitempty"`
}

type ApiUser struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Role string `json:"role"`
}