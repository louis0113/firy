package status

import (
	"fmt"
	"github.com/louis0113/firy/scales"
	"testing"
)

func TestStatusFahrenheint1(t *testing.T) {
	f := new(scales.Fahrenheit)

	f.FahrenheitStats(0.00)

	want := fmt.Sprintf("%s%s\n", msg_status, points[0])
	r, err := StatusFahrenheit(212.00, f)

	if want != r || err != nil {
		t.Errorf("TestCelsius1 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
func TestStatusFahrenheint2(t *testing.T) {
	f := new(scales.Fahrenheit)
	f.FahrenheitStats(0.00)

	want := fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, 92.00, f.GetSymbol(), points[1])
	r, err := StatusFahrenheit(120.00, f)

	if want != r || err != nil {
		t.Errorf("TestCelsius2 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
func TestStatusFahrenheint3(t *testing.T) {
	f := new(scales.Fahrenheit)

	f.FahrenheitStats(0.00)

	want := fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, -359.00, f.GetSymbol(), points[2])
	r, err := StatusFahrenheit(-100.00, f)

	if want != r || err != nil {
		t.Errorf("TestCelsius3 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
func TestStatusFahrenheint4(t *testing.T) {
	f := new(scales.Fahrenheit)

	f.FahrenheitStats(0.00)

	want := fmt.Sprintf("%s%s\n", msg_status, points[3])
	r, err := StatusFahrenheit(-459.00, f)

	if want != r || err != nil {
		t.Errorf("TestCelsius4 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
