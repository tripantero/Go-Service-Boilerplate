package models

import (
	"os"

	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect provides abstraction to get access to MongoDB
func Connect() error {
	return mgm.SetDefaultConfig(nil, os.Getenv("DBNAME"), options.Client().ApplyURI(os.Getenv("DBURI")))
}

// Create simplify insert document to collection
// but please pass reference parameter with reference of object that already created
func Create(reference interface{}) error {
	convertedRef := reference.(mgm.Model)
	return mgm.Coll(convertedRef).Create(convertedRef)
}
