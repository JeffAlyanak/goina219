package main

import (
	"fmt"

	"github.com/jeffalyanak/goina219"
)

func main() {

	config := goina219.Config(
		goina219.Range16V,
		goina219.Gain320MV,
		goina219.Adc12Bit,
		goina219.Adc12Bit,
		goina219.ModeContinuous,
	)

	// i := goina219.CalibrationValue(16, 0.004, float64(goina219.Gain160MV))

	// fmt.Println(i)

	myINA219, err := goina219.New(
		0x40,  // ina219 address
		0x01,  // i2c bus
		0.001, // Shunt resistance in ohms
		config,
		goina219.Gain320MV,
	)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	i := 0
	for i < 100 {
		err = goina219.Read(myINA219)
		if err != nil {
			panic(fmt.Sprintf("%v", err))
		}

		fmt.Printf(
			"Power: %fw, Current: %fa, Voltage: %fv, Shunt: %fv\n",
			(myINA219.Power * goina219.PowerMultiplier),
			(myINA219.Current / goina219.CurrentDivider),
			myINA219.Bus,
			myINA219.Shunt,
		)
	}
}
