package models

import (
	"os"

	"github.com/Kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BaseModel declare the common action that model can do
type BaseModel interface {
	GetCollection() *mgm.Collection
}

// SetConfig provides abstraction to get access to MongoDB
func SetConfig() error {
	return mgm.SetDefaultConfig(nil, os.Getenv("DBNAME"), options.Client().ApplyURI(os.Getenv("DBURI")))
}

// DatabasePing ping to database server whether still connected or not
func DatabasePing() error {
	client, _ := mgm.NewClient()
	err := client.Ping(mgm.Ctx(), nil)
	return err
}
