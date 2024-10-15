package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type response struct {
	Message string  `json:"message"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)

		data, _ := json.Marshal(response{
			Message: "An error occurred",
		})
		w.Write(data)
        return
    }

    if creds.Username != "admin" || creds.Password != "password" {
        w.WriteHeader(http.StatusUnauthorized)
		data, _ := json.Marshal(response{
			Message: "Incorrect username or password",
		})
		w.Write(data)
        return
    }
	token, err := generateToken(creds.Username)
	if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
		data, _ := json.Marshal(response{
			Message: "An error occurred",
		})
		w.Write(data)
        return
    }
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
        Value:   token,
        Expires: time.Now().Add(1 * time.Hour),

	})

}
