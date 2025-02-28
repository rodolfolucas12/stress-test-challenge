package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rodolfolucas12/stress-test-challenge/config"
)

func main() {
	url, requests, concurrency := parseFlags()

	if err := validateFlags(url, requests, concurrency); err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}

	config.Init(&url, &requests, &concurrency)
}

func parseFlags() (string, int, int) {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 0, "Número total de requests")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")

	flag.Parse()

	return *url, *requests, *concurrency
}

func validateFlags(url string, requests int, concurrency int) error {
	if url == "" {
		return fmt.Errorf("URL não pode ser vazia")
	}
	if requests <= 0 {
		return fmt.Errorf("número total de requests deve ser maior que zero")
	}
	if concurrency <= 0 {
		return fmt.Errorf("número de chamadas simultâneas deve ser maior que zero")
	}
	return nil
}
