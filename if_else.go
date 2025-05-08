package main

import "fmt"

func main() {

	// if condition {
	// block of code
	// }

	//age := 25
	//if age >= 18 {
	//	fmt.Println("I align with my highest self")
	//}

	// if condition {
	// block of code
	// } else if {
	// block of code
	// } else {
	// block of code
	// }

	//temperature := 25
	//if temperature >= 30 {
	//	fmt.Println("It's hot outside")
	//	} else {
	//fmt.Println("It's cool outside")
	//}

	//score := 85

	//if score >= 90 {
	//	fmt.Println("Grade A")
	//} else if score >= 80 {
	//	fmt.Println("Grade B")
	//} else if score >= 70 {
	//	fmt.Println("Grade C")
	//} else {
	//	fmt.Println("Grade D")
	//}

	//if score >= 90 {
	//	fmt.Println("Grade A")
	//}
	//if score >= 80 {
	//	fmt.Println("Grade B")
	//}
	//if score >= 70 {
	//	fmt.Println("Grade C")
	//} else {
	//	fmt.Println("Grade D")
	//}
	// This line will be executed after one of the conditions

	// if condition1 {
	// code block
	// if condition2 {
	// code block 2
	//}
	//}

	//num := 18
	//if num%2 == 0 {
	//if num%3 == 0 {
	//	fmt.Println("number is divisible by 2 and 3")
	//} else {
	//	fmt.Println("Number is divisible by 2 but not 3")
	//}
	//} else {
	//	fmt.Println("Not divisible by 2")
	//}

	//num := 3
	//if num%5 == 0 {
	//	if num%3 == 0 {
	//	fmt.Println("This number is divisible by 5 and 3")
	//} else {
	//		fmt.Println("This number is divisible by 5 and not 3")
	//}
	//} else {
	//	fmt.Println("Not divisible by 5")
	//}

	num := 3
	if num%3 == 0 {
		if num%10 == 0 {
			fmt.Println("This number is divisible by 3 and 10")
		} else {
			fmt.Println("This number is divisible by 3 and not 10")
		}
	} else {
		fmt.Println("This number is not divisible by 3")
	}
}
