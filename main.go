package main

import (
	"github.com/Satssuki/Go-Service-Boilerplate/helpers"
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	"github.com/Satssuki/Go-Service-Boilerplate/route"
	"github.com/joho/godotenv"
)

func main() {
	helpers.SetupLogger()

	err := godotenv.Load()
	helpers.Warn(err)

	err = models.SetConfig()
	helpers.Warn(err)

	err = models.DatabasePing()
	helpers.Error(err)

	err = route.SetupRouter().Run(helpers.GetPORT())
	helpers.Fatal(err)
}
