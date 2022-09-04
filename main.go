package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Hello to the %v booking app \n", conferenceName)
	fmt.Printf("we have %v tickets and remaining tickets: %v \n", conferenceTickets, remainingTickets)
	fmt.Println(&remainingTickets)
	fmt.Printf("conN is %T, remT is %T \n", conferenceName, remainingTickets)

	var bookings = []string{}

	for {
		
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//as username
		fmt.Println("enter firstName")
		fmt.Scan(&firstName)
		fmt.Println("enter lastName")
		fmt.Scan(&lastName)
		fmt.Println("enter email")
		fmt.Scan(&email)
		fmt.Println("enter tickets")
		fmt.Scan(&userTickets)
	
		remainingTickets = remainingTickets - userTickets
	
		bookings = append(bookings, firstName +" "+ lastName)
		fmt.Printf("Slice  %v \n", bookings)
		fmt.Printf("Slice  %v \n", bookings[0])
		fmt.Printf("Slice  %v \n", len(bookings))
	
		fmt.Printf("userName %v %v has booked %v tickets email: %v \n", firstName, lastName, userTickets, email)
		
		firstNames := []string{}
		for _, booking := range bookings{
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}
		fmt.Printf("%v booked first ticket \n", firstNames)

		var noticks bool = remainingTickets == 0
		if noticks {
			// end programme
			fmt.Println("no more tickets")
			break
		}
	}
	
}