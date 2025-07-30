package status

import (
	"fmt"
	"github.com/louis0113/firy/scales"
)

var (
	msgF   string
	points = [...]string{"exact boiling point of water", "boiling point of water", "boiling point of water", "exact freezing point of water"}
	temp   float64
)

const msg_status = "The status is: "

func StatusKelvin(kel float64, k *scales.Kelvin) (string, error) {

	switch {

	case kel == k.GetMaxTemp():
		msgF = fmt.Sprintf("%s%s\n", msg_status, points[0])
		return msgF, nil

	case kel < k.GetMaxTemp() && kel > k.GetTermMinTemp():
		temp = k.GetMaxTemp() - kel

		msgF = fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, temp, k.GetSymbol(),
			points[1])
		return msgF, nil

	case kel > k.GetMinTemp():
		temp = k.GetMinTemp() - kel
		msgF = fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, temp, k.GetSymbol(),
			points[2])
		return msgF, nil

	default:
		msgF = fmt.Sprintf("%s%s\n", msg_status, points[3])
		return msgF, nil
	}

}

func StatusFahrenheit(fah float64, f *scales.Fahrenheit) (string, error) {
	switch {

	case fah == f.GetMaxTemp():
		msgF = fmt.Sprintf("%s%s\n", msg_status, points[0])
		return msgF, nil

	case fah < f.GetMaxTemp() && fah > f.GetTermMinTemp():
		temp = f.GetMaxTemp() - fah

		msgF = fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, temp, f.GetSymbol(),
			points[1])
		return msgF, nil

	case fah > f.GetMinTemp():
		temp = f.GetMinTemp() - fah
		msgF = fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, temp, f.GetSymbol(),
			points[2])
		return msgF, nil

	default:
		msgF = fmt.Sprintf("%s%s\n", msg_status, points[3])
		return msgF, nil

	}

}

func StatusCelsius(cel float64, c *scales.Celsius) (string, error) {

	switch {

	case cel == c.GetMaxTemp():
		msgF = fmt.Sprintf("%s%s\n", msg_status, points[0])
		return msgF, nil

	case cel < c.GetMaxTemp() && cel > c.GetTermMinTemp():
		temp = c.GetMaxTemp() - cel

		msgF = fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, temp, c.GetSymbol(),
			points[1])
		return msgF, nil

	case cel > c.GetMinTemp():
		temp = c.GetMinTemp() - cel
		msgF = fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, temp, c.GetSymbol(),
			points[2])
		return msgF, nil

	default:
		msgF = fmt.Sprintf("%s%s\n", msg_status, points[3])
		return msgF, nil

	}

}
