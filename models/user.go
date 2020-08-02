package models

import (
	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	// collStruct yep again this is just dummy objet for initializing the collection naame
	collStruct = &User{}
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

// FindAllUsers return user list from database
func FindAllUsers() (*[]User, error) {
	users := []User{}
	cursor, err := mgm.Coll(collStruct).Find(mgm.Ctx(), bson.M{})
	cursor.All(mgm.Ctx(), &users)
	return &users, err
}
