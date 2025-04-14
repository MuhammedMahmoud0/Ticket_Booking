package main

import (
	"fmt"
	"project1/helper"
	"sync"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = conferenceTickets

// booked := []string{}
// var booked []string
// var booked = make([]map[string]string,0)
// var booked []map[string]string
// var booked = make([]UserData,0)
var booked []helper.UserData

var wg = sync.WaitGroup{}

func main() {

	helper.GreatingUser(conferenceName, remainingTickets, conferenceTickets)

	for {
		firstName, lastName, email, bookedTickets := helper.GetUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, bookedTickets, remainingTickets)

		if isValidTicketNumber && isValidName && isValidEmail {
			remainingTickets, booked = helper.BookTickets(remainingTickets, bookedTickets, conferenceName, firstName, lastName, email, booked)
			wg.Add(1)
			go helper.SendTicket(bookedTickets, firstName, lastName, email)
		} else {
			// fmt.Printf("only exist a total number of %v tickets\n", remainingTickets)
			// fmt.Printf("Your input data is invalid, try again\n")
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}

		if remainingTickets == 0 {
			fmt.Println("The conference tickets are sold out")
			break
		}
		wg.Wait()
	}
	helper.PrintFirstNames(booked)
}
