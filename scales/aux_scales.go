package scales

import (
	"fmt"
	"math/big"
)

var v bool

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

func truncate(f float64, unit float64) float64 {
	bf := big.NewFloat(0).SetPrec(1000).SetFloat64(f)
	bu := big.NewFloat(0).SetPrec(1000).SetFloat64(unit)

	bf.Quo(bf, bu)

	i := big.NewInt(0)
	bf.Int(i)
	bf.SetInt(i)

	f, _ = bf.Mul(bf, bu).Float64()
	return f
}

func validation(temp, maxi, minm float64) bool {

	if temp > maxi || temp < minm {
		return false
	}
	return true
}

func msgError(tStringMax string, tValueMax float64, tStringMin string, tValueMin float64) (msg string) {

	msg = fmt.Sprintf("%s max-temp has %.2f and %s min-temp has %.2f\n", tStringMax, tValueMax, tStringMin, tValueMin)
	return
}
