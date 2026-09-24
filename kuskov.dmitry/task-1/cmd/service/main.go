package main

import "fmt"

func main() {
	fmt.Println("Please enter the first operand")
	var first, second int
	var op string
	_, err := fmt.Scan(&first)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	fmt.Println("Please enter the second operand")

	_, err1 := fmt.Scan(&second)
	if err1 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	fmt.Println("Please enter the operator")

	fmt.Scan(&op)
	if op != "+" && op != "-" && op != "*" && op != "/" {
		fmt.Println("Invalid operation")
		return
	}
	
	var result int
	if op == "+" {
		result = first + second
	}
	if op == "-" {
		result = first - second
	}
	if op == "*" {
		result = first * second
	}
	if op == "/" {
		if second == 0{
			fmt.Println("Division by zero")
			return
		}
		result = first / second
	}

	fmt.Println(result)
}
