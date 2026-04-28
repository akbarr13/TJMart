package handler

import (
	"TJMart/models"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method tidak valid!", http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		Username string `json:"username"`
		Password string `json:"password:"`
	}

	json.NewDecoder(r.Body).Decode(&body)

	var users = models.GetUsers()
	for i := range users {
		if body.Username == users[i].Username && body.Password == users[i].Password {
			http.SetCookie(w, &http.Cookie{
				Name:  "session",
				Value: users[i].Username,
				Path:  "/",
			})
			json.NewEncoder(w).Encode(map[string]any{
				"success": true,
				"message": "Login berhasil!",
			})
			return
		}
	}

}

func Register(w http.ResponseWriter, r *http.Request) {

}
