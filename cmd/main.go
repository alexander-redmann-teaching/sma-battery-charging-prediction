package main

import (
	"flag"
	"github.com/magiconair/properties"
	"log"
)
import "github.com/alexander-redmann-teaching/sma-battery-charging-prediction/internal"

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.properties", "file where to read properties from")
	flag.Parse()
	log.Println("Parse config file from: " + configFile)
	p := properties.MustLoadFile(configFile, properties.UTF8)
	apiUrl, _ := p.Get("tibber.api.url")
	apiToken, _ := p.Get("tibber.api.token")
	internal.GetPriceForecast(apiUrl, apiToken)

}
