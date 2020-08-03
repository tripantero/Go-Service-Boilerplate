package models

import (
	"github.com/Kamva/mgm/v3"
)

// User just dummy struct
type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `bson:"name" json:"name"`
	Age              int    `bson:"age"  json: "age"`
}

// Collection return the collection of current user model
func (user *User) Collection() *mgm.Collection {
	return mgm.Coll(user)
}
