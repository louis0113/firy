package status

import (
	"github.com/louis0113/firy/scales"
	"testing"
)

func TestStatusKelvin(t *testing.T) {

	k := new(scales.Kelvin)
	want := ""
	r, err := StatusKelvin(200, k)

	if want != r || err != nil {

	}
}

func TestStatusCelsius(t *testing.T) {
	c := new(scales.Celsius)
	want := ""
	r, err := StatusCelsius(50.60, c)

	if want != r || err != nil {

	}
}

func TestStatusFahrenheint(t *testing.T) {
	f := new(scales.Fahrenheit)

	want := ""
	r, err := StatusFahrenheit(50.60, f)

	if want != r || err != nil {

	}
}
