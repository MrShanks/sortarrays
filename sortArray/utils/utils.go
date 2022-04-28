package utils

import (
	"fmt"
	"log"
	"math"
	"time"
)

var maxWaitTime = 8000.0

func RetryBackOff(attempt float64) {
	var wait float64
	var next float64

	wait = math.Pow(2, attempt)
	next = math.Pow(2, attempt+1)
	if wait > maxWaitTime {
		wait = maxWaitTime
		next = maxWaitTime
	}
	time.Sleep(time.Duration(wait) * time.Millisecond)
	log.Println(fmt.Sprintf("Next connection attempt will be performed in %v seconds", next/1000))
}
