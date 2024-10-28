package calculator

type Calculator struct{}

func (c Calculator) Add(a, b float64) float64 {
    return a + b
}

func (c Calculator) Subtract(a, b float64) float64 {
	return a - b
}

func (c Calculator) Multiply(a, b float64) float64 {
	return a * b
}
