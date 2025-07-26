package main

import (
	"fmt"
	"github.com/louis0113/firy/scales"
	"log"
)

func main() {
  


	celsius := 100.00
	kelvin, err := scales.CelsiusToKelvin(celsius)

  if err != nil {
    log.Fatal(err)
  } else{
		fmt.Printf("Celsius: %.2f Cº\nKelvin: %.2f Kº\n ", celsius, kelvin)
  } 

	if err != nil {
		log.Fatal(err)
	} 
	
  msg , err := scales.StatusCelsius(75.60)
  fmt.Println(msg)
}
