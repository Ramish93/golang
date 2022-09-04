package main

import (
	"fmt"
	"strings"
)
var conferenceName string = "Go Conference"
const conferenceTickets = 50
var remainingTickets int = 50
var bookings = []string{}

func main() {
	greetUsers()
	
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)



		if isValidEmail && isValidName && isValidTicketNumber{
			bookTickets(userTickets , firstName, lastName)
		
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
			var names = strings.Fields(booking)
			var firstName = names[0]
			firstNames = append(firstNames, firstName)
		}
		return firstNames

	}

func validateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
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
		
	bookings = append(bookings, firstName +" "+ lastName)
	fmt.Printf("Slice  %v \n", bookings)
	fmt.Printf("remainingTickets: %v\n", remainingTickets)

}