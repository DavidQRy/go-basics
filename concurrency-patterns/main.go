package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s %v\n", url, err)
		return
	}

	defer resp.Body.Close()

	duration := time.Since(start)
	fmt.Printf("Fetched %s in %v - Status %s\n", url, duration, resp.Status)
}

func main() {

	urls := []string{
		"https://rickandmortyapi.com/api/character",   // Lista de personajes
		"https://rickandmortyapi.com/api/character/1", // Detalle de un personaje (Rick Sanchez)
		"https://rickandmortyapi.com/api/location",    // Lista de ubicaciones
		"https://rickandmortyapi.com/api/location/1",  // Detalle de una ubicación (Earth - C-137)
		"https://rickandmortyapi.com/api/episode",     // Lista de episodios
	}

	start := time.Now()

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg)
	}

	wg.Wait()

	durationTotal := time.Since(start)

	fmt.Printf("Total sequence time:  %s", durationTotal)
}
