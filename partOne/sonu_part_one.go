package partOne

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

// accepts the value in ns
func CalculateTAT(execTime []time.Duration) time.Duration {
	var finalTAT time.Duration
	fmt.Println("Single goroutine going for sleep")
	for i := 0; i < len(execTime); i++ {
		// sleeping for the execTime in ms
		timeInMs := 1000000 * execTime[i]
		time.Sleep(timeInMs)
		finalTAT += timeInMs
	}
	return finalTAT
}

func main() {
	// since you mentioned that only a single CORE
	runtime.GOMAXPROCS(1)
	// N  = number of scripts
	var N int
	fmt.Println("Enter the number of scripts: ")
	fmt.Scan(&N)
	fmt.Println("Enter the executation time for every script: ")
	execTime := make([]time.Duration, N)
	for i := 0; i < N; i++ {
		if _, err := fmt.Scan(&execTime[i]); err != nil {
			log.Fatal(err)
		}
	}
	totalTAT := CalculateTAT(execTime)
	fmt.Println(totalTAT)
	parseTAT := totalTAT.String()
	u, _ := time.ParseDuration(parseTAT)
	fmt.Println("Total turn-around-time in milliseconds :", u.Milliseconds())
}
