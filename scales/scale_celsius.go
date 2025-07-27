package scales

import "errors"

type Celsius struct {
	Temp          float64
	Max_temp      float64
	Term_min_temp float64
	Min_temp      float64
	Symbol        rune
	Name          string
}

func (c *Celsius) CelsiusStats(temp float64) {

	c.Temp = temp
	c.Max_temp = 100.00
	c.Term_min_temp = 0.00
	c.Min_temp = -273.15
	c.Symbol = '℃'
	c.Name = `Celsius`
}

func (k *Kelvin) CelsiusToKelvin(c *Celsius) (float64, error) {

	v = validation(c.Temp, c.Max_temp, c.Min_temp)

	if v {

		k.Temp = c.Temp + k.Term_min_temp
		return k.Temp, nil

	}

	return 0.0, errors.New(msgError(c.Name, c.Max_temp, c.Name, c.Min_temp))

}

func (f *Fahrenheit) CelsiusToFahrenheit(c *Celsius) (float64, error) {

	v = validation(c.Temp, c.Max_temp, c.Min_temp)

	if v {
		f.Temp = (c.Temp/5.00)*9.00 + f.Term_min_temp
		f := truncate(f.Temp, 0.000001)

		return f, nil
	}

	return 0.0, errors.New(msgError(c.Name, c.Max_temp, c.Name, c.Min_temp))

}
