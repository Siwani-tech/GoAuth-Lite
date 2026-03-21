package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Siwani-tech/GoAuth-Lite.git/internal/middleware"
	"github.com/Siwani-tech/GoAuth-Lite.git/internal/models"
	"github.com/Siwani-tech/GoAuth-Lite.git/internal/services"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GoAuth Lite is running"))
}

func SignpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid request body",
		})
		return
	}
	err = services.Signup(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "user created ",
	})

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	token, err := services.Login(user.Email, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	emails := r.Context().Value(middleware.Useremailkey).(string)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hii there " + "  " + emails))
}
