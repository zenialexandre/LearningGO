package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	//Validating user input
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidUserTickets bool = userTickets > 0 && userTickets <= int(remainingTickets)
	return isValidName, isValidEmail, isValidUserTickets // returning more than one value
}
