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

	case kel == k.Max_temp:
		msgF = fmt.Sprintf("%s%s\n", msg_status, points[0])
		return msgF, nil

	case kel < k.Max_temp && kel > k.Term_min_temp:
		temp = k.Max_temp - kel

		msgF = fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, temp, k.Symbol,
			points[1])
		return msgF, nil

	case kel > k.Min_temp:
		temp = k.Min_temp - kel
		msgF = fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, temp, k.Symbol,
			points[2])
		return msgF, nil

	default:
		msgF = fmt.Sprintf("%s%s\n", msg_status, points[3])
		return msgF, nil
	}

}

func StatusFahrenheit(fah float64, f *scales.Fahrenheit) (string, error) {
	switch {

	case fah == f.Max_temp:
		msgF = fmt.Sprintf("%s%s\n", msg_status, points[0])
		return msgF, nil

	case fah < f.Max_temp && fah > f.Term_min_temp:
		temp = f.Max_temp - fah

		msgF = fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, temp, f.Symbol,
			points[1])
		return msgF, nil

	case fah > f.Min_temp:
		temp = f.Min_temp - fah
		msgF = fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, temp, f.Symbol,
			points[2])
		return msgF, nil

	default:
		msgF = fmt.Sprintf("%s%s\n", msg_status, points[3])
		return msgF, nil

	}

}

func StatusCelsius(cel float64, c *scales.Celsius) (string, error) {

	switch {

	case cel == c.Max_temp:
		msgF = fmt.Sprintf("%s%s\n", msg_status, points[0])
		return msgF, nil

	case cel < c.Max_temp && cel > c.Term_min_temp:
		temp = c.Max_temp - cel

		msgF = fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, temp, c.Symbol,
			points[1])
		return msgF, nil

	case cel > c.Min_temp:
		temp = c.Min_temp - cel
		msgF = fmt.Sprintf("%s there is %.2f %c from the %s\n", msg_status, temp, c.Symbol,
			points[2])
		return msgF, nil

	default:
		msgF = fmt.Sprintf("%s%s\n", msg_status, points[3])
		return msgF, nil

	}

}
