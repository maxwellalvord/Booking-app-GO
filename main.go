package main

import (
		"fmt"
		"strings"
)		
func main() {
	// var conferenceName = "Go Conference" //dry go version below
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	// var bookings []string //dry go version below
	bookings := []string{}

	fmt.Printf("conference tickets is %T, and remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)

		// fmt.Printf("The whole slice: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("slice type: %T\n", bookings)
		// fmt.Printf("slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings { // _ are used to ignore a variable you don't want to use
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first name of bookings are: %v\n", firstNames)
	}
	
	// var bookings = [50]string{}// dry go version below
	// bookings := [50]string{}


}