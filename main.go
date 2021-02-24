package main

import "fmt"

//go run main.go
func main() {
	var frontier int = 15
	result := whileFunction(frontier)
	ifValidation := ifElseTest(frontier)
	switchValidation := switchFunction(frontier)
	switchWithoutValue := switchWithoutExpression(frontier)
	frontierDefer := deferTest(frontier)
	frontierDefer2 := deferTest2(frontier)
	fmt.Printf("value type is: %T, and value is: %v \n", result, result)
	fmt.Printf("value type is: %T, and value is: %v \n", ifValidation, ifValidation)
	fmt.Printf("value type is: %T, and value is: %v \n", switchValidation, switchValidation)
	fmt.Printf("value type is: %T, and value is: %v \n", switchWithoutValue, switchWithoutValue)
	fmt.Printf("value type is: %T, and value is: %v \n", frontierDefer, frontierDefer)
	fmt.Printf("value type is: %T, and value is: %v \n", frontierDefer2, frontierDefer2)
}

func deferTest(frontier int) int {
	defer switchFunction(frontier)
	frontier++
	fmt.Printf("[deferTest]: value frontier its: %v", frontier)
	return frontier
}

func deferTest2(frontier int) int {
	fmt.Print("[deferTest2] Starting")

	for i := 0; i < frontier; i++ {
		defer fmt.Printf("[deferTest2] value : %v \n", i)
	}

	fmt.Print("[finish]")
	frontier++
	return frontier
}

func switchWithoutExpression(frontier int) bool {
	result := false

	switch {
	case frontier == 15:
		fmt.Println("frontier its 15")
		result = true
	case frontier > 15:
		fmt.Println("frontier its 15")
		result = false
	}

	return result
}

func switchFunction(frontier int) int {
	fmt.Printf("[switchFunction]: frontier es %v", frontier)
	result := 0
	switch frontier {
	case 15:
		fmt.Println("El frontier es 15")
		result = frontier + 1
	case 16:
		fmt.Println("El frontier es 16")
		result = frontier - 1
	default:
		fmt.Println("El frontier es una cosa rara")
	}

	return result
}

func forFunction1(frontier int) int {
	result := 0
	for i := 0; i < frontier; i++ {
		result += i
	}

	return result
}

func whileFunction(frontier int) int {
	result := 1
	for result < frontier {
		result += result
	}

	return result
}

func ifTest(frontier int) bool {
	result := false

	if frontier != 16 {
		result = true
	}

	return result
}

func ifTest2(frontier int) bool {
	result := false
	if frontier++; frontier != 16 {
		result = true
	}

	return result
}

func ifElseTest(frontier int) bool {
	result := false
	if frontier++; frontier != 16 {
		result = true
	} else {
		fmt.Println("Estamos en else bitches")
		result = false
	}

	return result
}
