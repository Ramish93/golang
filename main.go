package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Hello to the %v booking app \n", conferenceName)
	fmt.Printf("we have %v tickets and remaining tickets: %v \n", conferenceTickets, remainingTickets)
	fmt.Println("you can get your tickets here")

	
}