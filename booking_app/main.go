package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string
	var userTickets int
	//var bookings = [50]string{} // empty with 50 slots
	//var bookings [50]string Array!!
	var bookings []string // Slice!
	const conferenceTickets = 50

	fmt.Printf("conferenceTickets is %T, remainingTickes is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//ask for user input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) // passing the memory location

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, firstName+" "+lastName)

	/*fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice lentgh: %v\n", len(bookings))*/

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("These are all our bookings: %v\n", bookings)
}
