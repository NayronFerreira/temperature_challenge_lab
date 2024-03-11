package main

import (
	"fmt"

	"github.com/NayronFerreira/temperature_challenge_lab/config"
	"github.com/NayronFerreira/temperature_challenge_lab/infra/web/controller"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(config)

	controller.InitializeRoutes(config)

}
