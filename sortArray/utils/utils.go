package utils

import (
	"fmt"
	"log"
	"math"
	"time"
)

func RetryBackOff(attempt float64) {
	wait := math.Pow(2, attempt)
	next := math.Pow(2, attempt+1)
	time.Sleep(time.Duration(wait) * time.Millisecond)
	log.Println(fmt.Sprintf("Next connection attempt will be performed in %v seconds", next/1000))
}
