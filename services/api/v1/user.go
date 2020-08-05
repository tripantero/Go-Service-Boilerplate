package v1

import (
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
func (user *UserService) Insert() (string, error) {
	err := validation.ValidateUser(&user.User)
	var message string = "Users created"
	if err == nil {
		_ = err
		currentUser := &user.User
		count, Err := currentUser.GetCollection().CountDocuments(mgm.Ctx(), bson.M{
			"name": user.User.Name,
			"age":  user.User.Age,
		})
		err = Err
		if count > 0 {
			err = nil
			message = "Users already defined"
		} else {
			err = currentUser.GetCollection().Create(currentUser)
		}
	}
	return message, err
}
