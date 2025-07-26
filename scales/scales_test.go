package scales

import "testing"

func TestCelsiusToKelvin(t *testing.T) {

	celsius = 100.00

	want := 373.00

	msg, err := CelsiusToKelvin(celsius)

	if msg != want || err != nil {
		t.Errorf("CelsiusToKelvin(celsius) = %.2f and want match for %.2f, nil\n%v\n", msg, want, err)
	}

}

func TestKelvinToCelsius(t *testing.T) {

	kelvin = 373.00

	want := 100.00

	msg, err := KelvinToCelsius(kelvin)

	if msg != want || err != nil {
		t.Errorf("KelvinToCelsius(fahrenheit)= %.2f and want match for %.2f, nil\n%v\n", msg, want, err)
	}
}

func TestCelsiusToFahrenheint(t *testing.T) {

	celsius = 40.00

	want := 104.00

	msg, err := CelsiusToFahrenheit(celsius)

	if msg != want || err != nil {
		t.Errorf("CelsiusToFahrenheit(celsius)= %.2f and want match for %.2f, nil\n%v\n", msg, want, err)
	}
}

func TestFahrenheintToCelsius(t *testing.T) {

	fahrenheit = 104.00

	want := 40.00

	msg, err := FahrenheitToCelsius(fahrenheit)

	if msg != want || err != nil {
		t.Errorf("FahrenheitToCelsius(fahrenheit)= %.2f anf want match for %.2f, nil\n%v\n", msg, want, err)
	}
}

func TestKelvinToFahrenheit(t *testing.T) {

	kelvin = 310.00

	want := 98.60

	msg, err := KelvinToFahrenheit(kelvin)

	if msg != want || err != nil {
		t.Errorf("KelvinToFahrenheit(kelvin)= %.2f and want match for %.2f, nil\n%v\n", msg, want, err)
	}
}

func TestFahrenheintToKelvin(t *testing.T) {

	fahrenheit = 98.60

	want := 310.00

	msg, err := FahrenheitToKelvin(fahrenheit)

	if msg != want || err != nil {
		t.Errorf("FahrenheitToKelvin(fahrenheit)= %.2f and want match for %.2f, nil\n%v\n", msg, want, err)
	}
}
