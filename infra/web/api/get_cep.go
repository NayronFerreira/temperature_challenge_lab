package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/NayronFerreira/temperature_challenge_lab/infra/web/model"
)

func (a API) GetCEP(cep string) (retVal model.ViaCepRes, err error) {
	timeoutSeconds, err := strconv.Atoi(a.Config.TimeoutSeconds)
	if err != nil {
		return retVal, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s/json/", a.Config.ViaCepHost, cep), nil)
	if err != nil {
		return retVal, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return retVal, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return retVal, err
	}

	if err = json.Unmarshal(body, &retVal); err != nil {
		return retVal, err
	}

	return retVal, nil
}
