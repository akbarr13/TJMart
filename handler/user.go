package handler

import (
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.Error(w, "Method tidak valid!", http.StatusMethodNotAllowed)
		return
	}
	
	var body struct {
		Username string `json:"username"`
		Password string `json:"password:"`
	}
	
	json.NewDecoder(r.Body).Decode(&body)
	
	
}