package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	fmt.Printf("conferenceTickets is %T , remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if !isValidName {
		fmt.Println("First name or last name is too short")
		return
	}

	if !isValidEmail {
		fmt.Println("email address doesn't contain @ sign")
		return
	}

	if !isValidTicketNumber {
		fmt.Println("Number of tickets is invalid")
		return
	}

	if !isValidName || !isValidEmail || !isValidTicketNumber {
		fmt.Printf("Your input data is invalid. Try again.")
		return
	}

	bookTickets(userTickets, firstName, lastName, email)

	wg.Add(1)
	go sendTicket(userTickets, firstName, lastName, email)

	firstNames := getFirstNames()
	fmt.Printf("The first names of bookings are: %v\n", firstNames)

	if remainingTickets == 0 {
		fmt.Println("Our conference is booked out. Come back next year.")
	}

	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceName, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	//ask user for their name
	fmt.Printf("Thank you %v %v for booking %v tickets. Your will receive a confirmation email on %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("########################")
	wg.Done()
}
