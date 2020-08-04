package v1

import (
	"errors"

	"github.com/Kamva/mgm/v3"
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	"github.com/Satssuki/Go-Service-Boilerplate/services/api/validation"
	"go.mongodb.org/mongo-driver/bson"
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
		count, err := currentUser.GetCollection().CountDocuments(mgm.Ctx(), bson.M{
			"name": user.User.Name,
			"age":  user.User.Age,
		})
		if count > 0 {
			return errors.New("User already there")
		} else {
			err = currentUser.GetCollection().Create(currentUser)
		}
		return err
	}
	return err
}
