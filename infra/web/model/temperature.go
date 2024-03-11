package model

func NewTemperatureInfo(celsius float64) TemperatureInfo {
	return TemperatureInfo{
		Celsius:    celsius,
		Fahrenheit: getFarenheitByCelsius(celsius),
		Kelvin:     getKelvinByCelsius(celsius),
	}
}

type TemperatureInfo struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

func getFarenheitByCelsius(celsius float64) float64 {
	return celsius*1.8 + 32
}

func getKelvinByCelsius(celsius float64) float64 {
	return celsius + 273.15
}
