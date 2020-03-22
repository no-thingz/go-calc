package models

import (
	"errors"
)

type Number struct {
	NumA      float64 `json:"a"`
	NumB      float64 `json:"b"`
	Result    float64 `json:"result"`
	ErrorText string  `json:"error_text"`
}

var (
	numbers []*Number
)

func CalcSum(n Number) Number {
	n.Result = n.NumA + n.NumB
	return n
}
func CalcSub(n Number) Number {
	n.Result = n.NumA - n.NumB
	return n
}
func CalcMul(n Number) Number {
	n.Result = n.NumA * n.NumB
	return n
}

func CalcDiv(n Number) (Number, error) {
	if n.NumB == 0 {
		n.ErrorText = "Divine by Zero"
		return n, errors.New("error")
	}
	n.Result = n.NumA / n.NumB
	return n, nil
}
