package v1

import (
	"errors"

	"github.com/Satssuki/Go-Service-Boilerplate/models"
)

// UserService struct that wrapper the user model api
type UserService struct {
	User models.User
}

// Insert implementation of function in base interface
func (user *UserService) Insert() error {
	if len(user.User.Name) >= 4 && user.User.Age > 0 {
		return models.Create(user.User)
	}
	return errors.New("Model not valid")
}
