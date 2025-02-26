package main

import (
	"flag"

	"github.com/rodolfolucas12/stress-test-challenge/config"
)

func main() {

	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 0, "Número total de requests")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")

	flag.Parse()

	if *url == "" || *requests == 0 || *concurrency == 0 {
		flag.Usage()
		return
	}

	config.Init(url, requests, concurrency)
}
