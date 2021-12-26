package main

import (
	"fmt"
	"go-tut/helpers"
	"sync"
	"time"
)

// Package level vars - must be initialized with var syntax not :=
const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50

//var bookings = []string{} // slice - list of
//var bookings = make([]map[string]string, 0) // create an empty slice/list of maps
var bookings = make([]UserData, 0) // empty list of userData structs

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// wait group - waits for the launched goroutine to finish
var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helpers.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)                                              // sets the number of goroutines to wait for
		go sendTicket(userTickets, firstName, lastName, email) // go - start a new goroutine

		// firstNames := getFirstNames(bookings)
		var firstNames []string = getFirstNames()

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is to short")
		}
		if !isValidEmail {
			fmt.Println("Email is not valid")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}
	wg.Wait() // blocks until the wait group counter is 0
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		// var names = strings.Fields(booking) // splits a string
		// firstNames = append(firstNames, names[0])
		// firstNames = append(firstNames, booking["firstName"]) map value accessed with [""] syntax
		firstNames = append(firstNames, booking.firstName) // struct value accessed with . syntax
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	// var myslice []string
	// var mymap map[string]string

	// var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket string = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("################")
	wg.Done() // decrements the wait group counter by 1, this is called by the goroutine to indicate that it's finished
}
