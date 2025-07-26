package scales

import "fmt"

const msg_status = "The status is: "

var msgF string

var pontos = [...]string{"exact boiling point of water", "boiling point of water", "boiling point of water", "exact freezing point of water"}

func StatusKelvin(k float64) (string, error) {

	if k == kelvin_max {

		msgF = fmt.Sprintf("%s%s\n", msg_status, pontos[0])
		return msgF, nil
	}

	if k < kelvin_max && k > kelvin_term_min {
		temp := kelvin_max - k
		msgF = fmt.Sprintf("%s There is %.2f Kº from the %s", msg_status, temp, pontos[1])
		return msgF, nil
	}

	if k > kelvin_min {
		temp := kelvin_min - k
		msgF = fmt.Sprintf("%s There is %.2f Kº from the %s", msg_status, temp, pontos[2])
		return msgF, nil
	}

	msgF = fmt.Sprintf("%s%s\n", msg_status, pontos[3])
	return msgF, nil

}

func StatusFahrenheit(f float64) (string, error) {

	if f == fahrenheit_max {

		msgF = fmt.Sprintf("%s%s\n", msg_status, pontos[0])
		return msgF, nil
	}

	if f < fahrenheit_max && f > fahrenheit_term_min {
		temp := fahrenheit_max - f
		msgF = fmt.Sprintf("%s There is %.2f Fº from the %s", msg_status, temp, pontos[1])
		return msgF, nil
	}

	if f > fahrenheit_min {
		temp := fahrenheit_min - f
		msgF = fmt.Sprintf("%s There is %.2f Fº from the %s", msg_status, temp, pontos[2])
		return msgF, nil
	}

	msgF = fmt.Sprintf("%s%s\n", msg_status, pontos[3])
	return msgF, nil

}

func StatusCelsius(c float64) (string, error) {

	if c == celsius_max {

		msgF = fmt.Sprintf("%s%s\n", msg_status, pontos[0])
		return msgF, nil
	}

	if c < celsius_max && c > celsius_term_min {
		temp := celsius_max - c
		msgF = fmt.Sprintf("%s There is %.2f Cº from the %s", msg_status, temp, pontos[1])
		return msgF, nil
	}

	if c > celsius_min {
		temp := celsius_min - c
		msgF = fmt.Sprintf("%s There is %.2f Cº from the %s", msg_status, temp, pontos[2])
		return msgF, nil
	}

	msgF = fmt.Sprintf("%s%s\n", msg_status, pontos[3])
	return msgF, nil

}
