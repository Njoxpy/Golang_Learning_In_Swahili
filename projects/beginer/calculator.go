package main

import "fmt"

func addition(number1 int, number2 int) int {
	result := number1 + number2
	return result
}

func subtraction(number1 int, number2 int) int {
	result := number1 - number2
	return result
}

func multiplication(number1 int, number2 int) int {
	result := number1 * number2
	return result
}

func division(number1 int, number2 int) int {
	if number2 != 0 {
		result := number1 + number2
		return result
	}
	return 1
}

func main() {
	var choice int = 1

	fmt.Println("Karibu katika Calculator kwa kutumia Lugha ya Go!")

	fmt.Print("Zifuatazo ni operation ambazo zipo katika progam hii, Chagua namba yake!\n")
	fmt.Print(("\t1. Addition\n \t2. Subtraction\n \t3. Multiplication\n \t4. Division\n"))

	fmt.Printf("Jaza chaguo lako la operesheni(1-4): ")
	fmt.Scanf("%d", choice)

	switch choice {
	case 1:
		jibu := addition(2, 3)
		fmt.Printf("Jibu ni %d", jibu)
	case 2:
		jibu := subtraction(2, 3)
		fmt.Printf("Jibu ni %d", jibu)
	case 3:
		jibu := multiplication(2, 3)
		fmt.Printf("Jibu ni %d", jibu)
	case 4:
		jibu := division(2, 3)
		fmt.Printf("Jibu ni %d", jibu)
	default:
		fmt.Println("Chaguo sio sahihi")
	}
}
