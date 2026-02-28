package repository

import (
	"errors"

	"github.com/Siwani-tech/GoAuth-Lite.git/internal/models"
)

var users = make(map[string]models.User)

func Saveuser(user models.User) error {
	users[user.Email] = user
	return nil
}

func GetUserByEmail(email string) (models.User, error) {
	user, exist := users[email]
	if !exist {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}
