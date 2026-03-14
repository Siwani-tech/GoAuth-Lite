package services

import (
	"errors"

	"github.com/Siwani-tech/GoAuth-Lite.git/internal/models"
	"github.com/Siwani-tech/GoAuth-Lite.git/internal/repository"
	"github.com/Siwani-tech/GoAuth-Lite.git/internal/utils"
)

func Signup(user models.User) error {

	if user.Email == "" || user.Password == "" {

		return errors.New("email and password are required")
	}
	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	HashPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = HashPass
	err = repository.Saveuser(user)
	if err != nil {
		return err
	}
	return nil
}

func Login(email string, password string) (string, error) {

	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	user, err := repository.GetUserByEmail(email)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = utils.CheckPasshash(password, user.Password)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.Email)

	if err != nil {
		return "", err
	}

	return token, nil
}
