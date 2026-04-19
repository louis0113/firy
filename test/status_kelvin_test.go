package test

import (
	"fmt"
	"github.com/louis0113/firy/scales"
	"github.com/louis0113/firy/status"
	"testing"
)

func TestStatusKelvin1(t *testing.T) {

	k := new(scales.Kelvin)
	k.KelvinStats(0.00)
	want := fmt.Sprintf("%s%s\n", status.Msg_status, status.Points[0])
	r, err := status.StatusKelvin(373.00, k)

	if want != r || err != nil {
		t.Errorf("TestCelsius1 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
func TestStatusKelvin2(t *testing.T) {

	k := new(scales.Kelvin)
	k.KelvinStats(0.00)
	want := fmt.Sprintf("%s there is %.2f %c from the %s\n", status.Msg_status, 82.50, k.GetSymbol(), status.Points[1])
	r, err := status.StatusKelvin(290.50, k)
	if want != r || err != nil {
		t.Errorf("TestCelsius2 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
func TestStatusKelvin3(t *testing.T) {

	k := new(scales.Kelvin)
	k.KelvinStats(0.00)
	want := fmt.Sprintf("%s there is %.2f %c from the %s\n", status.Msg_status, -180.50, k.GetSymbol(), status.Points[2])
	r, err := status.StatusKelvin(180.50, k)
	if want != r || err != nil {
		t.Errorf("TestCelsius3 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
func TestStatusKelvin4(t *testing.T) {

	k := new(scales.Kelvin)
	k.KelvinStats(0.00)
	want := fmt.Sprintf("%s%s\n", status.Msg_status, status.Points[3])
	r, err := status.StatusKelvin(0.00, k)
	if want != r || err != nil {
		t.Errorf("TestCelsius4 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}

}
