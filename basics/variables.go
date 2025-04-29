package basics

import "fmt"

var middleName = "Cane"

func main() {
	//var age int
	//var name string = "Jon"
	//var name1 = "Jace"

	//count := 10
	//lastName := "Smith"

	middleName = "Mayor"
	fmt.Println(middleName)
	// Default values
	//Numeric Types: 0
	//Boolean Types: false
	//string type: ""
	//Pointers, slices, maps, functions, and structs: nilconst

	// ---- SCOPE

	//fmt.Println(firstName)

}

func printname() {
	firstName := "Michael"
	fmt.Println(firstName)
}
