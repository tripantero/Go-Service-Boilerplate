package main

import (
	"github.com/Satssuki/Go-Service-Boilerplate/helpers"
	"github.com/Satssuki/Go-Service-Boilerplate/models"
	"github.com/Satssuki/Go-Service-Boilerplate/route"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helpers.PanicIfErrNotNil(err)

	err = models.Connect()
	helpers.PanicIfErrNotNil(err)

	err = route.SetupRouter().Run(helpers.GetPORT())
	helpers.PanicIfErrNotNil(err)
}
