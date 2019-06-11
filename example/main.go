package main

import (
	"fmt"

	"github.com/jeffalyanak/ina219"
)

func main() {

	config := ina219.Config(
		ina219.Range32V,
		ina219.Gain320MV,
		ina219.Adc12Bit,
		ina219.Adc12Bit,
		ina219.ModeContinuous,
	)

	myINA219, err := ina219.New(
		0x40, // ina219 address
		0x00, // i2c bus
		0.01, // Shunt resistance in ohms
		config,
	)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	err = ina219.Read(myINA219)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	fmt.Printf(
		"Power: %fw, Current: %fa, Voltage: %fv, Shunt: %fv",
		myINA219.Power,
		myINA219.Current,
		myINA219.Bus,
		myINA219.Shunt,
	)
}
