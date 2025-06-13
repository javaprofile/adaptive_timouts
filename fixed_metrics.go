package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numTransactions = 10
	timeout         = 100 * time.Millisecond
)

type Transaction struct {
	ID        int
	RetryCount int
}

func transaction(t *Transaction) bool {
	sleepTime := time.Duration(rand.Intn(200))* time.Millisecond
	time.Sleep(sleepTime)

	if sleepTime > timeout {
		return false
	}
	return true
}

func snapshotIsolation(t *Transaction, wg *sync.WaitGroup) {
	defer wg.Done()

	retries := 0
	for {
		if transaction(t) {
			break
		} else {
			retries++
		}
	}

	t.RetryCount = retries
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	var transactions []Transaction

	for i := 0; i < numTransactions; i++ {
		t := Transaction{ID: i}
		wg.Add(1)
		go snapshotIsolation(&t, &wg)
		transactions = append(transactions, t)
	}

	wg.Wait()

	for _, t := range transactions {
		fmt.Printf("Transaction %d - Retry Count: %d\n", t.ID, t.RetryCount)
	}
}
