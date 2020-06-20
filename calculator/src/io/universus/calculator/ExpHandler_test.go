package calculator

import (
	"testing"
)

func TestParse(t *testing.T) {

	h := &ExpHandler{NewCalculator()}

	assertEquals(t, 0.0, h.parseUrlPath("/exp/0.0"), Delta)
	assertEquals(t, 0.01, h.parseUrlPath("/exp/0.01"), Delta)
	assertEquals(t, 1.37, h.parseUrlPath("/exp/1.37"), Delta)

}
