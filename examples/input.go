package main 

import "fmt"

func main()  {
	fmt.Println("Enter your name: ")
	
	var name string
	fmt.Scanln(&name)

	var num1 int 
	var num2 int 

	fmt.Println("Enter first number - ")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number - ")
	fmt.Scanln(&num2)

	sum := num1 + num2

	fmt.Println("Hello, ", name)
	fmt.Println("Summation: ", sum)

}
