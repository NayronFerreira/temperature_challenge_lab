package api

import (
	"encoding/json"
	"net/http"

	"github.com/NayronFerreira/temperature_challenge_lab/infra/web/model"
	"github.com/go-chi/chi"
)

func (a API) GetTemperatureByCEP(w http.ResponseWriter, r *http.Request) {

	cep := chi.URLParam(r, "cep")
	cepRes, err := a.GetCEP(cep)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	weatherRes, err := a.GetTemperature(cepRes.Localidade, cepRes.Uf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tempRes := model.NewTemperatureInfo(weatherRes.Current.TempC)

	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(tempRes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
