package main

import (
	"fmt"
	"time"
)
var conferenceName string = "Go Conference"
const conferenceTickets = 50
var remainingTickets int = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets int
}

func main() {
	greetUsers()
	
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)



		if isValidEmail && isValidName && isValidTicketNumber{
			bookTickets(userTickets , firstName, lastName)
			go sendTicket(userTickets , firstName, lastName, email)
		
			fmt.Printf("userName %v %v has booked %v tickets email: %v \n", firstName, lastName, userTickets, email)
			
			var firstname = getFirstName()
			fmt.Printf("firstName: %v\n", firstname)

			var noticks bool = remainingTickets == 0
			if noticks {
				// end programme
				fmt.Println("no more tickets")
				break
			}
		} else{
			if !isValidName {
				fmt.Println("firstname or last name is short")
			}
			if !isValidEmail {
				fmt.Println("email is short")
			}
			if !isValidTicketNumber {
				fmt.Println("invlaid ticket number is short")
			}
			continue
		}
	}
	// switch statement example
	// city := "London"

	// switch city{
	// 	case "new your":
	// 		//execute this
	// 	case "singapore", "london":
	// 		//execute this
	// 	case "hong kong":
	// 		//execute this
	// 	default:
	// 		//execute this
	// 		fmt.Println("no valid city")
	// }
}

func greetUsers() {
	fmt.Printf("welcome to %v\n booking app", conferenceName)
	fmt.Printf("we have %v tickets and remaining tickets: %v \n", conferenceTickets, remainingTickets)
}

func getFirstName()[]string {
	firstNames := []string{}
		for _, booking := range bookings{
			firstName := booking.firstName
			firstNames = append(firstNames, firstName)
		}
		return firstNames

	}

func getUserInput () (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int
	//as username
	fmt.Println("enter firstName")
	fmt.Scan(&firstName)
	fmt.Println("enter lastName")
	fmt.Scan(&lastName)
	fmt.Println("enter email")
	fmt.Scan(&email)
	fmt.Println("enter tickets")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets( userTickets int, firstName string, lastName string){
	remainingTickets = remainingTickets - userTickets

	//create a map
	var userData = userData{
		firstName: firstName,
		lastName: lastName,
		numberOfTickets: userTickets,
	}
	
	bookings = append(bookings, userData)
	fmt.Printf("list of bookings map %v", bookings)
	fmt.Printf("Slice  %v \n", bookings)
	fmt.Printf("remainingTickets: %v\n", remainingTickets)

}

func sendTicket (userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v %v", userTickets, firstName, lastName)
	fmt.Println("----------------")
	fmt.Printf("sending ticket: %v to email %v ", ticket, email)
	fmt.Println("----------------")
}