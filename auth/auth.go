package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub       string
	Event_id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidateToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		fmt.Println("Invalid token")
		return false, nil, "Invalid token"
	}

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("The token cannot be decoded: ", err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("Cannot decode json structure: ", err.Error())
		return false, err, err.Error()
	}

	// Fecha actual
	now := time.Now()
	// Fecha de vencimiento del token
	expTime := time.Unix(int64(tkj.Exp), 0)
	if expTime.Before(now) {
		fmt.Println("Token expiration date: ", expTime.String())
		fmt.Println("Token expirated!!!")
		return false, err, "Token expirated!!!"
	}

	return true, nil, tkj.Username
}