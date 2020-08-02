package main

import (
	"github.com/Satssuki/Go-Service-Boilerplate/helpers"
	"github.com/Satssuki/Go-Service-Boilerplate/route"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	route.SetupRouter().Run(helpers.GetPORT())
}
