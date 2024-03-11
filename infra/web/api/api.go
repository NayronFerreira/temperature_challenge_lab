package api

import "github.com/NayronFerreira/temperature_challenge_lab/config"

type API struct {
	Config config.Config
}

func NewAPI(config config.Config) API {
	return API{Config: config}
}
