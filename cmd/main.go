package main

import (
	"github.com/magiconair/properties"
)
import "github.com/alexander-redmann-teaching/sma-battery-charging-prediction/internal"

func main() {
	p := properties.MustLoadFile("./config.properties", properties.UTF8)
	apiToken, _ := p.Get("tibber.api.token")
	apiUrl, _ := p.Get("tibber.api.url")
	internal.GetPriceForecast(apiUrl, apiToken)

}
