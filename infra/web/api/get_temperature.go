package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/NayronFerreira/temperature_challenge_lab/infra/web/model"
)

func (a API) GetTemperature(city, state string) (retVal model.WeatherRes, err error) {
	timeoutSeconds, err := strconv.Atoi(a.Config.TimeoutSeconds)
	if err != nil {
		return retVal, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	location := url.QueryEscape(fmt.Sprintf("%s,%s", city, state))
	weatherUrl := fmt.Sprintf("%s/current.json?key=%s&q=%s", a.Config.WeatherHost, a.Config.WeatherKey, location)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, weatherUrl, nil)
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
