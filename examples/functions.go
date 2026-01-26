package main 

import "fmt"

func add(num1 int,num2 int) int {
	sum := num1 + num2

	// fmt.Println(sum)

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2

	mut := num1 * num2 

	return sum, mut
}


func main() {
  a := 10
	b := 20

	sum := add(a, b) // 30

	fmt.Println(sum)

	p, q := getNumbers(a, b) // 30, 200

	fmt.Println(p)

	fmt.Println(q)
}
