package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Hello to the %v booking app \n", conferenceName)
	fmt.Printf("we have %v tickets and remaining tickets: %v \n", conferenceTickets, remainingTickets)
	fmt.Println(&remainingTickets)
	fmt.Printf("conN is %T, remT is %T \n", conferenceName, remainingTickets)

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
	fmt.Printf("userName %v %v has booked %v tickets email: %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining", remainingTickets)
	
}