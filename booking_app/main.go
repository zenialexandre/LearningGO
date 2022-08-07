package main

import (
	"fmt"
	"sync"
	"time"
)

// Declaring variables at package level *global variables*
const conferenceTickets = 50

var conferenceName string = "Go conference"
var bookings = make([]UserData, 0) // List of structs!
var remainingTickets uint = 50

type UserData struct { // storing mixed data types
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var wg = sync.WaitGroup{} // created to wait for a go routine

func main() {
	//var bookings = [50]string{} // empty with 50 slots
	//var bookings [50]string Array!!
	greetUsers()

	for {
		// for loop with statement: for remainingTickets > 0 && len(bookings) < 50 {}
		firstName, lastName, email, userTickets := getUserInput() // values receiving the returns in order
		isValidName, isValidEmail, isValidUserTickets := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)                                              // adding routines to the thread
			go sendTicket(userTickets, firstName, lastName, email) // starts a new go routine
			fmt.Printf("The first name of bookings are: %v\n", getFirstNames())

			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Come back next year.")
				//break // end the loop
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contains a @ sign")
			}

			if !isValidUserTickets {
				fmt.Println("Number of tickets you entered is invalid")
			}
			//continue // go to the next iteration
		}
		wg.Wait() // waiting for completing of the threads
	}
}

func greetUsers() { // params need type
	fmt.Printf("conferenceTickets is %T, remainingTickes is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string { // type of return
	firstNames := []string{}

	for _, booking := range bookings { // _ for unused variable
		//var names = strings.Fields(booking) // to split the string
		//firstNames = append(firstNames, booking["firstName"]) -> for maps
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	//ask for user input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) // passing the memory location

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)

	// map for a user
	/*var userData = make(map[string]string) // key / value
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatInt(int64(userTickets), 10)*/

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	/*fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice lentgh: %v\n", len(bookings))*/

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)                                                       // simulating execution
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) // saving in string var
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("###################")
	wg.Done() // removes the thread that was added
}
