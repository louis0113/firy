package scales

import "testing"

func TestCelsiusToKelvin(t *testing.T) {

	c := new(Celsius)
	cel := 100.00
	c.CelsiusStats(cel)

	want := 373.00
	k := new(Kelvin)
	k.KelvinStats(0.00)
	msg, err := k.CelsiusToKelvin(c)

	if msg != want || err != nil {
		t.Errorf("CelsiusToKelvin(celsius) = %.2f and want match for %.2f, nil\n%v\n", msg, want, err)
	}

}

func TestKelvinToCelsius(t *testing.T) {
	k := new(Kelvin)
	kel := 373.00
	k.KelvinStats(kel)
	want := 100.00
	c := new(Celsius)
	c.CelsiusStats(0.00)
	msg, err := c.KelvinToCelsius(k)

	if msg != want || err != nil {
		t.Errorf("KelvinToCelsius(fahrenheit)= %.2f and want match for %.2f, nil\n%v\n", msg, want, err)
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	c := new(Celsius)
	cel := 40.00
	c.CelsiusStats(cel)

	want := 104.00
	f := new(Fahrenheit)
	f.FahrenheitStats(0.00)
	msg, err := f.CelsiusToFahrenheit(c)

	if msg != want || err != nil {
		t.Errorf("CelsiusToFahrenheit(celsius)= %.2f and want match for %.2f, nil\n%v\n", msg, want, err)
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	f := new(Fahrenheit)
	fa := 104.00
	f.FahrenheitStats(fa)

	want := 40.00
	c := new(Celsius)
	c.CelsiusStats(0.00)
	msg, err := c.FahrenheitToCelsius(f)

	if msg != want || err != nil {
		t.Errorf("FahrenheitToCelsius(fahrenheit)= %.2f anf want match for %.2f, nil\n%v\n", msg, want, err)
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	k := new(Kelvin)

	kel := 310.00
	k.KelvinStats(kel)
	want := 98.60
	f := new(Fahrenheit)
	f.FahrenheitStats(0.00)
	msg, err := f.KelvinToFahrenheit(k)

	if msg != want || err != nil {
		t.Errorf("KelvinToFahrenheit(kelvin)= %.2f and want match for %.2f, nil\n%v\n", msg, want, err)
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	f := new(Fahrenheit)

	fa := 98.60
	f.FahrenheitStats(fa)
	want := 310.00
	k := new(Kelvin)
	k.KelvinStats(0.00)
	msg, err := k.FahrenheitToKelvin(f)

	if msg != want || err != nil {
		t.Errorf("FahrenheitToKelvin(fahrenheit)= %.2f and want match for %.2f, nil\n%v\n", msg, want, err)
	}
}
