package main 

import "fmt"

func main()  {
	age := 20

	if age > 18 {
		fmt.Println("You are eligible to be married")
	} else if age < 18 {
		fmt.Println("You are not eligible to be married");
	} else {
		fmt.Println("You are just a teenager, not eligible to be married!");
	}

	// >=, <=, < , <=, ==
	//  and - &&
	// or - ||
	// not => !
	sex := "male"

	if age == 20 && sex == "male" {
		fmt.Println("You are ready to marry!")
	} else if age == 30 || sex == "male" {
		fmt.Println("You are male and maybe 30 years old")
	}

	isPretty := false

	if (!isPretty) {
		fmt.Println("You are not pretty!")
	}

	i := 3

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Not one or two")
	}
}
