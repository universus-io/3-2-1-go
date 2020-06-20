package calculator

import (
	"testing"
)

const Delta = 0.0001

func TestExp(t *testing.T) {

	calc := &Calculator{}

	assertEquals(t, 0.1224560, calc.Exp(-2.1), Delta)
	assertEquals(t, 0.3678790, calc.Exp(-1.0), Delta)
	assertEquals(t, 1.0000000, calc.Exp(0.0), Delta)
	assertEquals(t, 2.7182820, calc.Exp(1.0), Delta)
	assertEquals(t, 8.1661700, calc.Exp(2.1), Delta)

}

func BenchmarkExp(b *testing.B) {

	calc := &Calculator{}

	for i := 0; i < b.N; i++ {
		calc.Exp(1.137)
	}

}
