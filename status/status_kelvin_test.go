package status

import (
	"fmt"
	"github.com/louis0113/firy/scales"
	"testing"
)

func TestStatusKelvin1(t *testing.T) {

	k := new(scales.Kelvin)
	k.KelvinStats(0.00)
	want := fmt.Sprintf("%s%s\n", msg_status, points[0])
	r, err := StatusKelvin(373.00, k)

	if want != r || err != nil {
		t.Errorf("TestCelsius1 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
func TestStatusKelvin2(t *testing.T) {

	k := new(scales.Kelvin)
	k.KelvinStats(0.00)
	want := fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, 82.50, k.GetSymbol(), points[1])
	r, err := StatusKelvin(290.50, k)
	if want != r || err != nil {
		t.Errorf("TestCelsius2 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
func TestStatusKelvin3(t *testing.T) {

	k := new(scales.Kelvin)
	k.KelvinStats(0.00)
	want := fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, -180.50, k.GetSymbol(), points[2])
	r, err := StatusKelvin(180.50, k)
	if want != r || err != nil {
		t.Errorf("TestCelsius3 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}
}
func TestStatusKelvin4(t *testing.T) {

	k := new(scales.Kelvin)
	k.KelvinStats(0.00)
	want := fmt.Sprintf("%s%s\n", msg_status, points[3])
	r, err := StatusKelvin(0.00, k)
	if want != r || err != nil {
		t.Errorf("TestCelsius4 = %s\n and want match for %s\n\n, nil\n%v\n", r, want, err)
	}

}
