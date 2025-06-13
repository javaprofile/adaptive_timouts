package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Transaction struct {
	id     int
	status string
	retries int
}

func processTransaction(t *Transaction, timeout int) {
	if rand.Float32() < 0.5 {
		t.status = "Timed Out"
		t.retries++
	} else {
		t.status = "Completed"
	}
}

func runTransactions(numTransactions int, baseTimeout int) {
	retryCounts := make([]int, numTransactions)

	for i := 0; i < numTransactions; i++ {
		t := &Transaction{id: i + 1}
		timeout := baseTimeout + rand.Intn(50)

		for t.status != "Completed" {
			processTransaction(t, timeout)
			if t.status == "Timed Out" {
				fmt.Printf("Transaction %d timed out, retrying...\n", t.id)
			}
		}
		retryCounts[i] = t.retries
		fmt.Printf("Transaction %d completed with %d retries.\n", t.id, t.retries)
	}

	fmt.Println("\nRetry Count Metrics:")
	for i, retries := range retryCounts {
		fmt.Printf("Transaction %d: %d retries\n", i+1, retries)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numTransactions := 5
	baseTimeout := 200

	runTransactions(numTransactions, baseTimeout)
}
