# Goina219

_Warning: goina219 is **not** ready for primetime!_

**Goina219** is a simple golang package for configuring and reading the INA219 Bidirectional Current/Power Monitor over I<sup>2</sup>C.

It is currently only available for *nix systems as it leverages the `github.com/NeuralSpaz/i2c` package for I<sup>2</sup>C.

[![License](https://img.shields.io/github/license/JeffAlyanak/goina219.svg)](https://github.com/jeffalyanak/goina219/blob/master/LICENSE.txt)
[![Donate](https://img.shields.io/badge/donate--green.svg)](https://jeff.alyanak.ca/donate)
[![PGP](https://img.shields.io/keybase/pgp/jeffalyanak.svg?label=pgp)](https://jeff.alyanak.ca/pgp)

## Requirements
 * Go (tested on 1.12.1, may work on older)
 * NeuralSpaz I<sup>2</sup>C: [https://github.com/NeuralSpaz/i2c](https://github.com/NeuralSpaz/i2c)

## Installation

With the dependancies installed, simply use:

```bash
go get github.com/jeffalyanak/goina219

```

## Usage

Usage is very simple and an example is included in the `example/` directory.

A [config](#configuration-settings) should be generated:


```golang
config := ina219.Config(
		ina219.Range32V,
		ina219.Gain320MV,
		ina219.Adc12Bit,
		ina219.Adc12Bit,
		ina219.ModeContinuous,
	)
```

And an INA219 struct initialized:

```golang
myINA219, err := ina219.New(
		0x40, // ina219 address
		0x00, // i2c bus
		0.01, // Shunt resistance in ohms
		config,
	)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
```

The read function can be called:

```golang
err := ina219.Read(myINA219)
	if err!= nil {
		// error
	}
```

Power and current can now be accessed from the struct:

```golang
fmt.Printf(
		"Power: %fw, Current: %fa, Voltage: %fv, Shunt: %fv",
		myINA219.Power,
		myINA219.Current,
		myINA219.Bus,
		myINA219.Shunt,
	)
```

## Configuration Settings

### Voltage range options:

|Range|Parameter|
|---|---|
|16V|goina219.Range16V|
|32V|goina219.Range32V|

### Gain options:

|Gain|Parameter|
|---|---|
|40mV|goina219.Gain40MV|
|80mV|goina219.Gain80MV|
|160mV|goina219.Gain160MV|
|320mV|goina219.Gain320MV|

### ADC sampling modes:

|Samples|Bit-depth|Sample Time|Parameter|
|---|---|---|---|
|1|9-bit|84μs|goina219.Adc9Bit|
|1|10-bit|148μs|goina219.Adc10Bit|
|1|11-bit|276μs|goina219.Adc11Bit|
|1|12-bit|532μs|goina219.Adc12Bit|
|2|12-bit|1060μs|goina219.Adc2Samp|
|4|12-bit|2130μs|goina219.Adc4Samp|
|8|12-bit|4260μs|goina219.Adc8Samp|
|16|12-bit|8510μs|goina219.Adc16Samp|
|32|12-bit|17020μs|goina219.Adc32Samp|
|64|12-bit|34050μs|goina219.Adc64Samp|
|128|12-bit|68100μs|goina219.Adc128Samp|

## Contributing
Merge requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
