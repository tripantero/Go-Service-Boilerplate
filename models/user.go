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

// InsertUser single function to insert User
func InsertUser(name string, age int) error {
	user := &User{Name: name, Age: age}
	return Create(user)
}
