package scales

import "errors"

type Kelvin struct {
	Temp          float64
	Max_temp      float64
	Term_min_temp float64
	Min_temp      float64
	Symbol        rune
	Name          string
}

func (k *Kelvin) KelvinStats(temp float64) {

	k.Temp = temp
	k.Max_temp = 373.00
	k.Term_min_temp = 273.00
	k.Min_temp = 0.00
	k.Symbol = 'K'
	k.Name = `Kelvin`
}

func (c *Celsius) KelvinToCelsius(k *Kelvin) (float64, error) {

	v = validation(k.Temp, k.Max_temp, k.Min_temp)
	if v {

		c.Temp = k.Temp - k.Term_min_temp
		return c.Temp, nil
	}

	return 0.0, errors.New(msgError(k.Name, k.Max_temp, k.Name, k.Min_temp))

}

func (f *Fahrenheit) KelvinToFahrenheit(k *Kelvin) (float64, error) {

	v = validation(k.Temp, k.Max_temp, k.Min_temp)
	if v {
		f.Temp = ((k.Temp-k.Term_min_temp)/5.00)*9.00 + f.Term_min_temp
		f := truncate(f.Temp, 0.000001)
		return f, nil
	}

	return 0.0, errors.New(msgError(k.Name, k.Max_temp, k.Name, k.Min_temp))

}
