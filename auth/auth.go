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
		fmt.Println("Invalid token, it must have 3 parts: ", len(parts))
		return false, nil, "Invalid token"
	}

	part1 := string(parts[1])
	fmt.Println("Part 1", part1, "a")
	userInfo, err := base64.StdEncoding.DecodeString(part1)
	fmt.Println("User info: ", userInfo)
	if err != nil {
		fmt.Println("The token cannot be decoded: ", err.Error())
		return false, err, err.Error()
	}
	fmt.Println("User info: ", userInfo)

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("Cannot decode json structure: ", err.Error())
		return false, err, err.Error()
	}
	fmt.Println("Token JSON: ", tkj)

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
