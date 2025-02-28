package app

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type UseCaseInterface interface {
	UseCase(url string, requests, calls int) (int, error)
}

type UseCase struct {
	client Client
}

func NewUseCase(client Client) *UseCase {
	return &UseCase{
		client: client,
	}
}

type Result struct {
	StatusCode int
	Duration   time.Duration
	Error      error
}

func (s *UseCase) StressUseCase(url *string, requests, concurrency *int) {
	start := time.Now()
	results := make(chan Result, *requests)
	var wg sync.WaitGroup

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go s.worker(*url, *requests, *concurrency, results, &wg)
	}

	wg.Wait()
	close(results)

	s.generateReport(results, time.Since(start), *requests)
}

func (s *UseCase) worker(url string, requests, concurrency int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < requests/concurrency; i++ {
		s.makeRequest(url, results)
	}
}

func (s *UseCase) makeRequest(url string, resultado chan<- Result) {
	start := time.Now()
	resp, err := s.client.RequestClient(url)
	duration := time.Since(start)

	if err != nil {
		resultado <- Result{Error: err}
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		resultado <- Result{Error: err}
		return
	}

	resultado <- Result{StatusCode: resp.StatusCode, Duration: duration}
}

func (s *UseCase) generateReport(resultado <-chan Result, tempoTotal time.Duration, totalRequests int) {
	status := make(map[int]int)
	var total time.Duration
	quantidadeOk := 0

	for r := range resultado {
		if r.Error != nil {
			fmt.Printf("Erro ao fazer request: %v\n", r.Error)
			continue
		}

		if r.StatusCode == http.StatusOK {
			quantidadeOk++
		}
		status[r.StatusCode]++
		total += r.Duration
	}

	s.printReport(tempoTotal, totalRequests, quantidadeOk, status, total)
}

func (s *UseCase) printReport(tempoTotal time.Duration, totalRequests, quantidadeOk int, status map[int]int, total time.Duration) {
	fmt.Printf("Tempo total gasto na execução: %v\n", tempoTotal)
	fmt.Printf("Total de requests realizados: %d\n", totalRequests)
	fmt.Printf("Requests com status HTTP 200: %d\n", quantidadeOk)
	fmt.Println("Distribuição de outros códigos de status HTTP")
	for statusCode, quantidade := range status {
		if statusCode != http.StatusOK {
			fmt.Printf("%d: %d requests\n", statusCode, quantidade)
		}
	}
	fmt.Printf("Tempo médio por request: %v\n", total/time.Duration(totalRequests))
}
