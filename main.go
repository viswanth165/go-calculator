package main

import (
	"fmt"
	"log"
	"main/addition"
	"main/divition"
	"main/multiplication"
	"main/subraction"
)

func operation(left_operand float32, right_operand float32, operator string) float32 {
	switch operator {
	case "+":
		return addition.Addition(left_operand, right_operand)
	case "-":
		return subraction.Subraction(left_operand, right_operand)
	case "*":
		return multiplication.Multiplication(left_operand, right_operand)
	case "/":
		return divition.Divition(left_operand, right_operand)
	default:
		log.Println("Invalid operator")

	}
	return 0
}
func main() {
	var left_operand float32
	var right_operand float32
	var operator string
	var length_of_operator int
	var continuecalcukation string
	clear := true
	get := false
	switc := 1

	log.Println("Enter the number =  ")
	fmt.Scan(&left_operand)

	for clear == true {
		if get == true {
			log.Println("Enter the number =  ")
			fmt.Scan(&right_operand)
			get = false
			switc++
			if switc == 3 {
				output := operation(left_operand, right_operand, operator)
				left_operand = output
				log.Println("Output of the calculation", output)
				switc = 1
			}
		}
		if get == false {
			log.Println("Enter the operater =  ")
			fmt.Scan(&operator)
			length_of_operator = len(operator)
			if length_of_operator > 1 {
				log.Println("Invalid operator")
				get = false
			} else {
				log.Println(operator)
				get = true
				switc++
			}
		}

		if operator == "=" {
			log.Println("Enter c to end the calculation =  ")
			fmt.Scanf("%c", &continuecalcukation)
			if continuecalcukation == "c" {
				log.Println(left_operand)
				clear = false
			} else {
				get = false
				switc--
			}
		}
	}
}
