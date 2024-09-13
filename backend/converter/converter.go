package converter

import (
	"strconv"
)

// consts for length
const (
	mmToMeter = 0.001
	cmToMeter = 0.01
	kmToMeter = 1000
	inToMeter = 0.0254
	ftToMeter = 0.3048
	ydToMeter = 0.9144
	mlToMeter = 1609.344
)

// consts for weight
const (
	mgToKillogram = 0.000001
	gToKillogram  = 0.001
	ozToKillogram = 0.0283495
	lbToKillogram = 0.453592
)

func ConvertTemperature(from, to, amount string) (float64, error) {
	result, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}

	switch from {
	case "Fahrenheit":
		result = (result - 32) / 1.8
	case "Kelvin":
		result = result - 273.15
	}

	switch to {
	case "Fahrenheit":
		result = result*1.8 + 32
	case "Kelvin":
		result = result + 273.15
	}

	return result, nil
}

func ConvertLength(from, to, amount string) (float64, error) {
	buffer, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}

	if from == to {
		return float64(buffer), nil
	}

	return convertFromMeters(to, convertToMeters(from, buffer)), nil
}

func convertToMeters(from string, amount float64) float64 {
	switch from {
	case "mm":
		return amount * mmToMeter
	case "cm":
		return amount * cmToMeter
	case "km":
		return amount * kmToMeter
	case "in":
		return amount * inToMeter
	case "ft":
		return amount * ftToMeter
	case "yd":
		return amount * ydToMeter
	case "ml":
		return amount * mlToMeter
	}

	return amount
}

func convertFromMeters(to string, amount float64) float64 {
	switch to {
	case "mm":
		return amount / mmToMeter
	case "cm":
		return amount / cmToMeter
	case "km":
		return amount / kmToMeter
	case "in":
		return amount / inToMeter
	case "ft":
		return amount / ftToMeter
	case "yd":
		return amount / ydToMeter
	case "ml":
		return amount / mlToMeter
	}

	return amount
}

func ConvertWeight(from, to, amount string) (float64, error) {
	buffer, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0.0, err
	}

	if from == to {
		return float64(buffer), nil
	}

	return convertFromKg(to, convertToKg(from, buffer)), nil
}

func convertToKg(from string, amount float64) float64 {
	switch from {
	case "g":
		return amount * gToKillogram
	case "mg":
		return amount * mgToKillogram
	case "oz":
		return amount * ozToKillogram
	case "lb":
		return amount * lbToKillogram
	}

	return amount
}

func convertFromKg(to string, amount float64) float64 {
	switch to {
	case "g":
		return amount / gToKillogram
	case "mg":
		return amount / mgToKillogram
	case "oz":
		return amount / ozToKillogram
	case "lb":
		return amount / lbToKillogram
	}

	return amount
}
