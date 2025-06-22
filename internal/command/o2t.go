// Naranza Vinti, Copyright 2025 Andrea Davanzo and contributors, License AGPLv3

package command

import (
	"encoding/json"
	"math/rand"
	"time"
	"vinti/internal/core"
)

func O2t(config *core.Config, clientID string) ( string, error) {

	expireTime := time.Now().Add(time.Second * time.Duration(config.TokenExpiresIn))

	token := generateToken(32)

	tokenData := core.TokenData{
		ClientID: clientID,
		Expire:   expireTime,
	}

	tokenByte, err := json.Marshal(tokenData)
	if err == nil {
		err = Sto(config, "_token", token, string(tokenByte))
	}

	if err != nil {
		token = ""
	}

	return token, err
}

func generateToken(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

