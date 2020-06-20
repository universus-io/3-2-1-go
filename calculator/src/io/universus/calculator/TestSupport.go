package calculator

import (
	"math"
	"testing"
)

func assertEquals(t *testing.T, expected float64, actual float64, delta float64) {

	if math.Abs(expected-actual) > delta {
		t.Fatalf("expected=%f, actual=%f", expected, actual)
	}

}
