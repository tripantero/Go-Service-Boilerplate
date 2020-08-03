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

// GetCollection return the collection of current user model
func (user *User) GetCollection() *mgm.Collection {
	return mgm.Coll(user)
}
