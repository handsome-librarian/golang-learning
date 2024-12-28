package main

import (
	"fmt"
	"math"
)

func Addition(values []float64) float64 {
	var sumOfUserValues float64
	sumOfUserValues = values[0]
	for i := 1; i < len(values); i++ {
		sumOfUserValues += values[i]
	}
	return sumOfUserValues
}

func Subtraction(values []float64) float64 {
	var finalValue float64
	finalValue = values[0]
	for i := 1; i < len(values); i++ {
		finalValue -= values[i]
	}
	return finalValue
}
func Mutiplication(values []float64) float64 {
	var finalValue float64
	finalValue = values[0]
	for i := 1; i < len(values); i++ {
		finalValue *= values[i]
	}
	return finalValue
}

func Division(values []float64) float64 {
	var finalValue float64
	finalValue = values[0]
	for i := 1; i < len(values); i++ {
		if values[i] == 0 {
			fmt.Println("ZeroDivisionError can't divide a number by zero")
			return 0
		}
		finalValue /= values[i]
	}
	return finalValue
}

func Squareroot(userValue float64) float64 { // used the Newton-Rasphon method to approximately get the square root
	if userValue < 0 {
		return math.NaN()
	}
	if userValue == 0 {
		return 0
	}
	const tolerance = 0.001
	var nextGuess float64
	guess := userValue / 2
	for {
		nextGuess = (guess + userValue/guess) / 2
		if math.Abs(nextGuess-guess) < tolerance {
			break
		}
		guess = nextGuess

	}
	const precision = 10 // Number of decimal places
	scale := math.Pow(10, precision)
	return math.Round(nextGuess*scale) / scale

}
func Square(userValues []float64) []float64 {
	var finalAnswer []float64
	for _, value := range userValues {
		squareOfValue := value * value
		finalAnswer = append(finalAnswer, squareOfValue)
	}

	return finalAnswer
}
