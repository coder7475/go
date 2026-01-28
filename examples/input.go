package main 

import "fmt"

func sum(var num1 int, var num2 int) int {
	sum := num1 + num2

	return sum
}

func getUserName() string {
	fmt.Println("Enter your name: ")
	
	var name string
	fmt.Scanln(&name)

	return name
}


func main()  {
	name := getUserName()

	var num1 int
	var num2 int

	fmt.Println("Enter first number - ")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number - ")
	fmt.Scanln(&num2)

	sum := sum(num1, num2)

	fmt.Println("Hello, ", name)
	fmt.Println("Summation: ", sum)
}
