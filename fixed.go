package main
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
const fixedTimeout = 100 * time.Millisecond
func transaction(id int) bool {
	workTime := time.Duration(rand.Intn(200)) * time.Millisecond
	time.Sleep(workTime)
	return workTime < fixedTimeout
}
func snapshotIsolation(transactionID int, wg *sync.WaitGroup) {
	defer wg.Done()
.	startTime := time.Now()
	retries := 0
.	for {
		if time.Since(startTime) > fixedTimeout {
			break
		}
.			if transaction(transactionID) {
			fmt.Printf("Transaction %d succeeded after %d retries\n", transactionID, retries)
			return
		} else {
			retries++
		}
	}
.	fmt.Printf("Transaction %d failed after %d retries due to timeout\n", transactionID, retries)
}
func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go snapshotIsolation(i, &wg)
	}
	wg.Wait()
}
