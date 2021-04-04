package test

import (
	partOne "sonu_challenge_SKIDOS/partOne"
	"testing"
	"time"
)

func TestCalculateTAT(t *testing.T) {
	// time in ns
	execTime := []time.Duration{100, 200, 300, 400, 500}
	finalTAT := partOne.CalculateTAT(execTime)
	// converting time in ms
	timeInMS := 1000000 * time.Duration(1500)
	// testing both the outputs
	if finalTAT != timeInMS {
		t.Error("finalTAT is incorrect")
	}

	execTimeTwo := []time.Duration{200, 300, 400, 500, 600, 700}
	finalTATTwo := partOne.CalculateTAT(execTimeTwo)
	timeInMSTwo := 1000000 * time.Duration(2700)
	if finalTATTwo != timeInMSTwo {
		t.Error("finalTATTwo is incorrect")
	}
}
