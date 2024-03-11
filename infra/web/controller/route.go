package controller

import (
	"net/http"

	"github.com/NayronFerreira/temperature_challenge_lab/config"
	"github.com/NayronFerreira/temperature_challenge_lab/infra/web/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func InitializeRoutes(config config.Config) {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/{cep}", api.NewAPI(config).GetTemperatureByCEP)
	http.ListenAndServe(":8080", router)
}
