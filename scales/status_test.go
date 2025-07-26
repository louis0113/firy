package scales

import "testing"

var m1, m2, m3, m4 string
var val bool

const boilingPoint = "The status is: exact boiling point of water"
const freezingPoint = "The status is: exact freezing point of water"
const m_error = "Could not determine the temperature point"

func TestStatusKelvin(t *testing.T) {
	m1, m2, m3, m4 = "The status is: There is {num} Kº from the boiling point of water", "The status is: There is {num} Kº from the freezing point of water", boilingPoint, freezingPoint

	msg, err := StatusKelvin(170.89)
	val = msg != m1 || msg != m2 || msg != m3 || msg != m4
	if val != true || err != nil {
		t.Errorf("%s:  %v", m_error, err)

	}
}

func TestStatusFahrenheit(t *testing.T) {
	m1, m2, m3, m4 = "The status is: There is {num} Fº from the boiling point of water", "The status is: There is {num} Fº from the freezing point of water", boilingPoint, freezingPoint
	msg, err := StatusFahrenheit(300.89)
	val = msg != m1 || msg != m2 || msg != m3 || msg != m4
	if val != true || err != nil {
		t.Errorf("%s:  %v", m_error, err)

	}
}

func TestStatusCelsius(t *testing.T) {
	m1, m2, m3, m4 = "The status is: There is {num} Cº from the boiling point of water", "The status is: There is {num} Cº from the freezing point of water", boilingPoint, freezingPoint
	msg, err := StatusCelsius(70.89)
	val = msg != m1 || msg != m2 || msg != m3 || msg != m4
	if val != true || err != nil {
		t.Errorf("%s:  %v", m_error, err)

	}
}
