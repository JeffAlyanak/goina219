package ina219

import (
	"math"

	"github.com/NeuralSpaz/i2c"
)

const configRegister = 0x00 // See datasheet for further details on these registers
const shuntVoltRegister = 0x01
const busVoltRegister = 0x02
const powerRegister = 0x03
const currentRegister = 0x04
const calibrationRegister = 0x05

const currentLSBFactor = 0x7FFF // See datasheet for details on calibration
const maxCalibrationVal = 0xFFFE

const Range16V uint16 = 0 // Voltage range values
const Range32V uint16 = 1

const Gain40MV uint16 = 0 // Gain values
const Gain80MV uint16 = 1
const Gain160MV uint16 = 2
const Gain320MV uint16 = 3

const Adc9Bit uint16 = 0     // 9-Bit — 84us.
const Adc10Bit uint16 = 1    // 10-Bit— 148us.
const Adc11Bit uint16 = 2    // 11-Bit— 2766us.
const Adc12Bit uint16 = 3    // 12-Bit— 532us.
const Adc2Samp uint16 = 9    // 2 samples at 12-Bit,— 1.06ms.
const Adc4Samp uint16 = 10   // 4 samples at 12-Bit,— 2.13ms.
const Adc8Samp uint16 = 11   // 8 samples at 12-Bit,— 4.26ms.
const Adc16Samp uint16 = 12  // 16 samples at 12-Bit,— 8.51ms
const Adc32Samp uint16 = 13  // 32 samples at 12-Bit,— 17.02ms.
const Adc64Samp uint16 = 14  // 64 samples at 12-Bit,— 34.05ms.
const Adc128Samp uint16 = 15 // 128 samples at 12-Bit,— 68.10ms.

const ModePowerDown uint16 = 0
const ModeShuntVoltageTrig uint16 = 1
const ModeBusVoltageTrig uint16 = 2
const ModeShuntandBusVoltageTrig uint16 = 3
const ModeDisableADC uint16 = 4
const ModeShuntVoltageCont uint16 = 5
const ModeBusVoltageCont uint16 = 6
const ModeContinuous uint16 = 7

type INA219 struct {
	I2C         i2c.I2CBus
	Bus         float64
	Address     uint8
	Config      uint16
	Calibration uint16
	Current     float64
	Shunt       float64
	Power       float64
	Load        float64
}

// Config word (uint16) with voltage range, gain level, bus & shunt ADC settings and mode selection.
func Config(voltage uint16, gain uint16, busADC uint16, shuntADC uint16, mode uint16) uint16 {
	return voltage<<13 | gain<<11 | busADC<<7 | shuntADC<<3 | mode
}

// CalibrationValue Calculate the config word (uint16) for the calibration register.
func CalibrationValue(maxVolt float64, shuntOhms float64) uint16 {
	currentLSB := (maxVolt / shuntOhms) / currentLSBFactor

	return uint16(math.Trunc(0.04096 / currentLSB * shuntOhms))
}

// New Initialize and return a new ina219 device.
func New(address uint8, i2cbus byte, shuntOhms float64, config uint16) (*INA219, error) {
	deviceBus := i2c.NewI2CBus(i2cbus)
	ina := &INA219{
		I2C:         deviceBus,
		Config:      config,
		Calibration: CalibrationValue(32, shuntOhms),
		Address:     address,
	}

	if err := ina.I2C.WriteWordToReg(ina.Address, configRegister, ina.Config); err != nil {
		return nil, err
	}
	if err := ina.I2C.WriteWordToReg(ina.Address, calibrationRegister, ina.Calibration); err != nil {
		return nil, err
	}

	return ina, nil
}

// Read all values from INA219, convert them to floats and store them in the INA219 struct.
func Read(ina *INA219) error {
	c, err := ina.I2C.ReadWordFromReg(ina.Address, currentRegister)
	if err != nil {
		return err
	}
	ina.Current = current(c)

	p, err := ina.I2C.ReadWordFromReg(ina.Address, powerRegister)
	if err != nil {
		return err
	}
	ina.Power = power(p)

	b, err := ina.I2C.ReadWordFromReg(ina.Address, busVoltRegister)
	if err != nil {
		return err
	}
	ina.Bus = bus(b)

	s, err := ina.I2C.ReadWordFromReg(ina.Address, shuntVoltRegister)
	if err != nil {
		return err
	}
	ina.Shunt = shunt(s)

	return nil
}

// These four functions just convert the data read from the ina219 registers to usable floats.
func current(c uint16) float64 {
	return (float64(int16(c)) * 0.0004)
}

func power(p uint16) float64 {
	return (float64(int16(p)) * 0.001 * 4 * 5 * 0.4)
}

func bus(b uint16) float64 {
	return (float64(int16(b>>3)*4) * 0.001)
}

func shunt(s uint16) float64 {
	return (float64(int16(s)) * 0.00001)
}
