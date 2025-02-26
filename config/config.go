package config

import (
	"github.com/rodolfolucas12/stress-test-challenge/app"
)

func Init(url *string, requests, concurrency *int) {
	client := app.NewClient()
	usecase := app.NewUseCase(*client)
	usecase.StressUseCase(url, requests, concurrency)
}
