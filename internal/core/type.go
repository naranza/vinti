// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package core

import "time"

type ApiRequest struct {
	Cmd          string `json:"cmd"`
	Folder       string `json:"folder,omitempty"`
	File         string `json:"file,omitempty"`
	Data         string `json:"data,omitempty"`
	GrantType    string `json:"grant_type,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	Role        string `json:"scope,omitempty"`
}

type ApiResponse struct {
  Code  int      			`json:"code"`
  Message string      `json:"message"`
  Files   []string    `json:"files,omitempty"`
	AccessToken string 	`json:"access_token,omitempty"`
	TokenType   string 	`json:"token_type,omitempty"`
	ExpiresIn   int    	`json:"expires_in,omitempty"`
}

type TokenData struct {
	ClientID string    `json:"client_id"`
	Expire   time.Time `json:"expire"`
}

type ClientInfo struct {
	Username     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Role        string `json:"scope"`
}