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

func (k Kelvin) GetTemp() float64 {

	return k.Temp
}

func (k *Kelvin) SetTemp(temp float64) {

	k.Temp = temp
}

func (k Kelvin) GetMaxTemp() float64 {

	return k.Max_temp
}

func (k Kelvin) GetTermMinTemp() float64 {

	return k.Term_min_temp
}

func (k Kelvin) GetMinTemp() float64 {

	return k.Min_temp
}

func (k Kelvin) GetSymbol() rune {

	return k.Symbol
}

func (k Kelvin) GetName() string {

	return k.Name
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
