package calculator

import (
	"math"
)

type Calculator struct {
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Exp(x float64) float64 {
	return math.Exp(x)
}
