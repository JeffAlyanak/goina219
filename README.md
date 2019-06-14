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
go get github.com/jeffalyanak/ina219

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

Voltage range options:


* Range16V
* Range32V

Gain options:

* Gain40MV
* Gain80MV
* Gain160MV 
* Gain320MV

ADC sampling modes:

* Adc9Bit, _9-Bit — 84us_.
* Adc10Bit, _10-Bit— 148us_.
* Adc11Bit, _11-Bit— 2766usv_.
* Adc12Bit, _12-Bit— 532us_.
* Adc2Samp, _2 samples at 12-Bit,— 1.06ms_.
* Adc4Samp, _4 samples at 12-Bit,— 2.13ms_.
* Adc8Samp, _8 samples at 12-Bit,— 4.26ms_.
* Adc16Samp, _16 samples at 12-Bit,— 8.51ms_
* Adc32Samp, _32 samples at 12-Bit,— 17.02ms_.
* Adc64Samp, _64 samples at 12-Bit,— 34.05ms_.
* Adc128Samp, _128 samples at 12-Bit,— 68.10ms_.

## Contributing
Merge requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
