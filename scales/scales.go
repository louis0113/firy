package scales

import (
	"errors"
	"fmt"
	"math/big"
)

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

var (
	celsius                                             float64
	celsius_max, celsius_term_min, celsius_min          float64 = 100.00, 0.00, -273.00
	kelvin                                              float64
	kelvin_max, kelvin_term_min, kelvin_min             float64 = 373.00, 273.00, 0.00
	fahrenheit                                          float64
	fahrenheit_max, fahrenheit_term_min, fahrenheit_min = 212.00, 32.00, -460.00
	v                                                   bool
	scales                                              = [...]string{"Celsius", "Fahrenheint", "Kelvin"}
)

func CelsiusToKelvin(c float64) (float64, error) {

	v = validation(c, celsius_max, celsius_min)
	if v {

		kelvin = c + 273.00
		return kelvin, nil

	}

	return 0.0, errors.New(msgError(scales[0], celsius_max, scales[0], celsius_min))

}

func KelvinToCelsius(k float64) (float64, error) {

	v = validation(k, kelvin_max, kelvin_min)
	if v {
		celsius = k - 273.00
		return celsius, nil
	}

	return 0.0, errors.New(msgError(scales[2], kelvin_max, scales[2], kelvin_min))

}

func CelsiusToFahrenheit(c float64) (float64, error) {

	v = validation(c, celsius_max, celsius_min)

	if v {
		fahrenheit = (c/5.00)*9.00 + 32.00
		f := truncate(fahrenheit, 0.000001)

		return f, nil
	}

	return 0.0, errors.New(msgError(scales[0], celsius_max, scales[0], celsius_min))

}

func FahrenheitToCelsius(f float64) (float64, error) {

	v = validation(f, fahrenheit_max, fahrenheit_min)
	if v {
		celsius = (f - 32.00) * 5.00 / 9.00
		c := truncate(celsius, 0.000001)
		return c, nil
	}

	return 0.0, errors.New(msgError(scales[1], fahrenheit_max, scales[1], fahrenheit_min))

}

func KelvinToFahrenheit(k float64) (float64, error) {

	v = validation(k, kelvin_max, kelvin_min)
	if v {
		fahrenheit = ((k-273.00)/5.00)*9.00 + 32.00
		f := truncate(fahrenheit, 0.000001)
		return f, nil
	}

	return 0.0, errors.New(msgError(scales[2], kelvin_max, scales[2], kelvin_min))

}

func FahrenheitToKelvin(f float64) (float64, error) {

	v = validation(f, fahrenheit_max, fahrenheit_min)
	if v {
		kelvin = ((f-32.00)/9.00)*5.00 + 273.00
		k := truncate(kelvin, 0.000001)
		return k, nil

	}

	return 0.0, errors.New(msgError(scales[1], fahrenheit_max, scales[1], fahrenheit_min))
}
