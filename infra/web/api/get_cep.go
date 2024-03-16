package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	model "github.com/NayronFerreira/temperature_challenge_lab/infra/web/model/response"
)

func (a API) GetLocationByCEP(cep string) (locality, UF string, err error) {
	timeoutSeconds, err := strconv.Atoi(a.Config.TimeoutSeconds)
	if err != nil {
		return locality, UF, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s/json/", a.Config.ViaCepHost, cep), nil)
	if err != nil {
		return locality, UF, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return locality, UF, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locality, UF, err
	}
	var viaCepRes model.ViaCepRes
	if err = json.Unmarshal(body, &viaCepRes); err != nil {
		return locality, UF, err
	}

	return viaCepRes.Localidade, viaCepRes.Uf, nil
}
