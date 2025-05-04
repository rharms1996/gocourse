package main

import "fmt"

func main() {

	//sum := 0
	//for {
	//sum += 10
	//fmt.Println("Sum:", sum)
	//if sum >= 10
	//}

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Print("Odd Number:", num)
		num++ // ++ incement operatior increases value by 1 and -- decrement operater decreases value by 1
	}
}
