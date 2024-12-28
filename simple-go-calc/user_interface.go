package main

import (
	"fmt"
	"strconv"
)

func getUserCommamd() string {
	var operation string
	fmt.Println(`Enter operation to perform(Press Ctrl + C to quit):
	0) Addition
	1) Subtraction
	2) Multiplication
	3) Division
	4) Squareroot
	5) Square`)
	fmt.Print(">")
	fmt.Scanln(&operation)
	return operation

}
func getUserValue() float64 { // for square root
	var userInput float64
	fmt.Println("Enter value to perform squareroot upon: ")
	fmt.Scanln(&userInput)
	return userInput
}

func getUserValues() []float64 {
	var userValues []float64
	var userInput string
	for {
		fmt.Println("Enter value to perform operation upon (enter q to quit)")
		fmt.Scanln(&userInput)
		if userInput == "q" {
			break
		}
		if userInput == "" {
			fmt.Println("No new input provided try again...")
			continue

		}
		finalValue, userError := strconv.ParseFloat(userInput, 64)
		if userError != nil {
			fmt.Println("Error:Invalid input try again")
			continue
		} else {
			userValues = append(userValues, finalValue)
		}

	}
	return userValues
}

func UserOperations() {

	switch getUserCommamd() {
	case "0":
		values := getUserValues()
		additionAnswer := Addition(values)
		fmt.Printf("The addition of %v is: %v\n", values, additionAnswer)

	case "1":
		values := getUserValues()
		subractionAnswer := Subtraction(values)
		fmt.Printf("The subtraction %v is: %v\n", values, subractionAnswer)
	case "2":
		values := getUserValues()
		mutiplicationAnswer := Subtraction(values)
		fmt.Printf("The mutiplication %v is: %v\n", values, mutiplicationAnswer)
	case "3":
		values := getUserValues()
		divisionAnswer := Division(values)
		fmt.Printf("The division of %v is: %v\n", values, divisionAnswer)
	case "4":
		value := getUserValue()
		squareRootAnswer := Squareroot(value)
		fmt.Printf("The squareroot of %v is %v\n", value, squareRootAnswer)
	case "5":
		values := getUserValues()
		squareAnswer := Square(values)
		fmt.Printf("The square of %v is %v\n", values, squareAnswer)
	case "6":
		fmt.Println("Quitting program....")
	default:
		fmt.Println("Invalid command try again")

	}
}
