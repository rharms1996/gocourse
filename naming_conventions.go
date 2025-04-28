package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// pascal case
	// ex CalculateArea, UserInfo, NewHTTPRequest
	// structs, interfaces, enums

	// snake_case
	// ex user_ID, first_name, http_request

	// UPPERCASE
	// Use case is constants

	// mixedCase
	// ex javaScript, htmlDocument, isValid

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("employeeID: ", employeeID)

}
