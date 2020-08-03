package v1

import (
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	"github.com/Satssuki/Go-Service-Boilerplate/services/api/validation"
)

// UserService struct that wrapper the user model api
type UserService struct {
	User models.User
}

// CreateUserService just a shorthand create object from struct
func CreateUserService() UserService {
	return UserService{models.User{}}
}

// Insert implementation of function in base interface
func (user *UserService) Insert() error {
	err := validation.ValidateUser(&user.User)
	if err == nil {
		currentUser := &user.User
		return currentUser.GetCollection().Create(currentUser)
	}
	return err
}
