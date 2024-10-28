package calculator

import "fmt"

type Operation struct {
    Operation string
    A         float64
    B         float64
    Result    float64
}

type History struct {
    operations []Operation
}

func (h *History) AddOperation(op Operation) {
    h.operations = append(h.operations, op)
}

func (h *History) PrintHistory() {
    for _, op := range h.operations {
        fmt.Printf("%s(%f, %f) = %f\n", op.Operation, op.A, op.B, op.Result)
    }
}
