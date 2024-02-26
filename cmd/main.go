package main

import (
	"flag"
	"github.com/alexander-redmann-teaching/sma-battery-charging-prediction/internal"
	"github.com/magiconair/properties"
	"log"
)

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
