package partTwo

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

// accepts the value in ns
func calculateTAT(execTime []time.Duration, data chan time.Duration) {
	// fmt.Println("Single goroutine going for sleep")
	for i := 0; i < len(execTime); i++ {
		// sleeping for the execTime in ms
		timeInMs := 1000000 * execTime[i]
		time.Sleep(timeInMs)
		data <- timeInMs
	}
}

func main() {
	// since you mentioned that 4 CPU CORE are to be utilized
	runtime.GOMAXPROCS(4)
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
	// creating a channel
	data := make(chan time.Duration)
	// running the code asynchronously
	go calculateTAT(execTime, data)
	// reading the data from the channel
	totalTAT := <-data
	// converting the time to string to pass in ParseDuration later
	parseTAT := totalTAT.String()
	u, _ := time.ParseDuration(parseTAT)
	fmt.Println("Total turn-around-time in milliseconds :", u.Milliseconds())
}
