package models

import (
	"encoding/json"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password:"`
}

const userFile = "data/user.json"

func GetUsers() []User {
	file, err := os.ReadFile(userFile)
	if err != nil {
		return []User{}
	}
	var users []User
	json.Unmarshal(file, &users)
	return users
}
