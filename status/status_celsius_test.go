package status

import (
	"fmt"
	"github.com/louis0113/firy/scales"
	"testing"
)

func TestStatusCelsius1(t *testing.T) {
	c := new(scales.Celsius)
	c.CelsiusStats(0.00)
	want := fmt.Sprintf("%s%s\n", msg_status, points[0])
	r, err := StatusCelsius(100.00, c)

	if want != r || err != nil {
		t.Errorf("TestCelsius1 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
func TestStatusCelsius2(t *testing.T) {
	c := new(scales.Celsius)
	c.CelsiusStats(0.00)

	want := fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, 49.40, c.GetSymbol(), points[1])

	r, err := StatusCelsius(50.60, c)

	if want != r || err != nil {
		t.Errorf("TestCelsius2 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)

	}
}
func TestStatusCelsius3(t *testing.T) {
	c := new(scales.Celsius)
	c.CelsiusStats(0.00)

	r, err := StatusCelsius(-100.90, c)

	want := fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, -172.25, c.GetSymbol(), points[2])
	if want != r || err != nil {
		t.Errorf("TestCelsius3 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
func TestStatusCelsiusi4(t *testing.T) {
	c := new(scales.Celsius)
	c.CelsiusStats(0.00)

	want := fmt.Sprintf("%s%s\n", msg_status, points[3])
	r, err := StatusCelsius(-273.15, c)

	if want != r || err != nil {
		t.Errorf("TestCelsius4 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
