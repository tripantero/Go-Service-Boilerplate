package models

import (
	"os"

	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BaseModel declare the common action that model can do
type BaseModel interface {
	Collection() *mgm.Collection
}

// Connect provides abstraction to get access to MongoDB
func Connect() error {
	return mgm.SetDefaultConfig(nil, os.Getenv("DBNAME"), options.Client().ApplyURI(os.Getenv("DBURI")))
}
