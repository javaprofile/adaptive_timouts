package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Transaction struct {
	id        int
	status    string
	startTime time.Time
}

func processTransaction(tx *Transaction, dynamicTimeout time.Duration) bool {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	if time.Since(tx.startTime) > dynamicTimeout {
		tx.status = "Timed Out"
		return false
	}
	return true
}

func runTransactions(transactionCount int, baseTimeout time.Duration) {
	var retryCount int
	var dynamicTimeout time.Duration

	for i := 1; i <= transactionCount; i++ {
		tx := &Transaction{
			id:        i,
			status:    "Started",
			startTime: time.Now(),
		}

		dynamicTimeout = baseTimeout + time.Duration(rand.Intn(50))*time.Millisecond

		for !processTransaction(tx, dynamicTimeout) {
			retryCount++
			fmt.Printf("Transaction %d failed, retrying... (Retry Count: %d)\n", tx.id, retryCount)
			tx.startTime = time.Now()
		}
		fmt.Printf("Transaction %d completed successfully after %d retries.\n", tx.id, retryCount)
	}

	fmt.Printf("Total retries: %d\n", retryCount)
}

func main() {
	transactionCount := 10
	baseTimeout := 200 * time.Millisecond
	runTransactions(transactionCount, baseTimeout)
}
