package internal

import "log"

func GetPriceForecast(tibberApiUrl string, apiToken string) {
	log.Default().Println(tibberApiUrl + " : " + apiToken)
}
