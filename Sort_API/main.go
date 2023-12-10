// main.go
package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"sync"
	"time"
)

type RequestPayload struct {
	ToSort [][]int `json:"to_sort"`
}

type ResponsePayload struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNS       int64   `json:"time_ns"`
}

func processSingleHandler(w http.ResponseWriter, r *http.Request) {
	var payload RequestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	startTime := time.Now()

	var sortedArrays [][]int
	for _, arr := range payload.ToSort {
		sorted := make([]int, len(arr))
		copy(sorted, arr)
		sort.Ints(sorted)
		sortedArrays = append(sortedArrays, sorted)
	}

	response := ResponsePayload{
		SortedArrays: sortedArrays,
		TimeNS:       time.Since(startTime).Nanoseconds(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func processConcurrentHandler(w http.ResponseWriter, r *http.Request) {
	var payload RequestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	startTime := time.Now()

	var sortedArrays [][]int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, arr := range payload.ToSort {
		wg.Add(1)
		go func(arr []int) {
			defer wg.Done()

			sorted := make([]int, len(arr))
			copy(sorted, arr)
			sort.Ints(sorted)

			mu.Lock()
			sortedArrays = append(sortedArrays, sorted)
			mu.Unlock()
		}(arr)
	}

	wg.Wait()

	response := ResponsePayload{
		SortedArrays: sortedArrays,
		TimeNS:       time.Since(startTime).Nanoseconds(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/process-single", processSingleHandler)
	http.HandleFunc("/process-concurrent", processConcurrentHandler)

	http.ListenAndServe(":8000", nil)
}
