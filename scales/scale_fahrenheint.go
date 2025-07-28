package scales

import "errors"

type Fahrenheit struct {
	Temp          float64
	Max_temp      float64
	Term_min_temp float64
	Min_temp      float64
	Symbol        rune
	Name          string
}

func (f *Fahrenheit) FahrenheitStats(temp float64) {

	f.Temp = temp
	f.Max_temp = 212.00
	f.Term_min_temp = 32.00
	f.Min_temp = -459.00
	f.Symbol = '℉'
	f.Name = `Fahrenheit`
}

func (f Fahrenheit) GetTemp() float64 {

	return f.Temp
}

func (f *Fahrenheit) SetTemp(temp float64) {

	f.Temp = temp
}

func (f Fahrenheit) GetMaxTemp() float64 {

	return f.Max_temp
}

func (f Fahrenheit) GetTermMinTemp() float64 {

	return f.Term_min_temp
}

func (f Fahrenheit) GetMinTemp() float64 {

	return f.Min_temp
}

func (f Fahrenheit) GetSymbol() rune {

	return f.Symbol
}

func (f Fahrenheit) GetName() string {

	return f.Name
}

func (c *Celsius) FahrenheitToCelsius(f *Fahrenheit) (float64, error) {

	v = validation(f.Temp, f.Max_temp, f.Min_temp)
	if v {
		c.Temp = (f.Temp - f.Term_min_temp) * 5.00 / 9.00
		c := truncate(c.Temp, 0.000001)
		return c, nil
	}

	return 0.0, errors.New(msgError(f.Name, f.Max_temp, f.Name, f.Min_temp))

}

func (k *Kelvin) FahrenheitToKelvin(f *Fahrenheit) (float64, error) {

	v = validation(f.Temp, f.Max_temp, f.Min_temp)
	if v {
		k.Temp = ((f.Temp-f.Term_min_temp)/9.00)*5.00 + k.Term_min_temp
		k := truncate(k.Temp, 0.000001)
		return k, nil

	}

	return 0.0, errors.New(msgError(f.Name, f.Max_temp, f.Name, f.Min_temp))
}
