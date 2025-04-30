package helper

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// var wg = sync.WaitGroup{}  //this can not be used because there is one in the main.go so they can't work together at the same time

func GreatingUser(conferenceName string, remainingTickets uint, conferenceTickets uint) {
	fmt.Printf("Welcome To %v Booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func PrintFirstNames(booked []UserData) {
	firstNames := []string{}
	for _, booking := range booked {
		// names := strings.Fields(booking)
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("These are all our bookings: %v\n", firstNames)

}

func ValidateUserInput(firstName string, lastName string, email string, bookedTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := bookedTickets <= remainingTickets && bookedTickets > 0
	return isValidName, isValidEmail, isValidTicketNumber
}

func GetUserInput() (string, string, string, uint) {
	var (
		firstName     string
		lastName      string
		email         string
		bookedTickets uint
	)
	fmt.Print("Enter your first name:\n")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name:\n")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email:\n")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets you want:\n")
	fmt.Scan(&bookedTickets)

	return firstName, lastName, email, bookedTickets
}

func BookTickets(remainingTickets uint, bookedTickets uint, conferenceName, firstName string, lastName string, email string, booked []UserData) (uint, []UserData) {
	remainingTickets = remainingTickets - bookedTickets
	// name := firstName + " " + lastName
	// var userData map[string]string
	// var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: bookedTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(bookedTickets), 10)

	// booked = append(booked, name)
	booked = append(booked, userData)
	fmt.Printf("list of booked tickets %v\n", booked)
	fmt.Printf("Thank you %v %v for booking %v tickets. you will recive a confirmation email at %v\n", firstName, lastName, bookedTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	return remainingTickets, booked
}

func SendTicket(bookedTickets uint, firstName string, lastName string, email string, wg *sync.WaitGroup) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", bookedTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("####################")
	wg.Done()
}
