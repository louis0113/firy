package test

import (
	"fmt"
	"github.com/louis0113/firy/scales"
	"github.com/louis0113/firy/status"
	"testing"
)

func TestStatusCelsius1(t *testing.T) {
	c := new(scales.Celsius)
	c.CelsiusStats(0.00)
	want := fmt.Sprintf("%s%s\n", status.Msg_status, status.Points[0])
	r, err := status.StatusCelsius(100.00, c)

	if want != r || err != nil {
		t.Errorf("TestCelsius1 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
func TestStatusCelsius2(t *testing.T) {
	c := new(scales.Celsius)
	c.CelsiusStats(0.00)

	want := fmt.Sprintf("%s there is %.2f %c from the %s\n", status.Msg_status, 49.40, c.GetSymbol(), status.Points[1])

	r, err := status.StatusCelsius(50.60, c)

	if want != r || err != nil {
		t.Errorf("TestCelsius2 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)

	}
}
func TestStatusCelsius3(t *testing.T) {
	c := new(scales.Celsius)
	c.CelsiusStats(0.00)

	r, err := status.StatusCelsius(-100.90, c)

	want := fmt.Sprintf("%s there is %.2f %c from the %s\n", status.Msg_status, -172.25, c.GetSymbol(), status.Points[2])
	if want != r || err != nil {
		t.Errorf("TestCelsius3 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
func TestStatusCelsiusi4(t *testing.T) {
	c := new(scales.Celsius)
	c.CelsiusStats(0.00)

	want := fmt.Sprintf("%s%s\n", status.Msg_status, status.Points[3])
	r, err := status.StatusCelsius(-273.15, c)

	if want != r || err != nil {
		t.Errorf("TestCelsius4 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
