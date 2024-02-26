package internal

import (
	"context"
	"github.com/hasura/go-graphql-client"
	"log"
	"net/http"
	"time"
)

type TibberResult struct {
	Viewer struct {
		Homes []struct {
			CurrentSubscription struct {
				PriceInfo struct {
					Current struct {
						Total    float32
						Energy   float32
						Tax      float32
						StartsAt time.Time
					}
					Today []struct {
						Total    float32
						Energy   float32
						Tax      float32
						StartsAt time.Time
					}
					Tomorrow []struct {
						Total    float32
						Energy   float32
						Tax      float32
						StartsAt time.Time
					}
				}
			}
		}
	}
}

func GetPriceForecast(tibberApiUrl string, apiToken string) (TibberResult, bool) {
	log.Default().Println("Get tibber data from: " + tibberApiUrl)

	client := graphql.NewClient(tibberApiUrl, nil)
	client = client.WithRequestModifier(func(r *http.Request) {
		r.Header.Set("Authorization", "Bearer "+apiToken)
	})

	var query TibberResult

	err := client.Query(context.Background(), &query, nil)

	if err != nil {
		log.Fatal("Error on receiving Tibber data", err)
		return query, false
	}
	return query, true
}
