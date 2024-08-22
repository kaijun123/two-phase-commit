package utils

import (
	"fmt"
	"math/rand"
)

// produces a true output 95% of the time
func RandOutcome() bool {
	if randomNum := rand.Float64(); randomNum >= RandomProbabiltyThreshold {
		fmt.Println("randOutcome: false")
		return false
	}
	fmt.Println("randOutcome: true")
	return true
}
