package main

import "fmt"

func main() {
	var first, second int
	var op string
	_, err := fmt.Scan(&first)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err1 := fmt.Scan(&second)
	if err1 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err2 := fmt.Scan(&op)
	if (err2 != nil) || (op != "+" && op != "-" && op != "*" && op != "/") {
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
