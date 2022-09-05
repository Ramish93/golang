package main

import (
	"fmt"
	"reflect"
	"sync"
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
var wg = sync.WaitGroup{}

func main() {
	greetUsers()
	

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)



		if isValidEmail && isValidName && isValidTicketNumber{
			bookTickets(userTickets , firstName, lastName)
			wg.Add(1)
			go sendTicket(userTickets , firstName, lastName, email)
		
			fmt.Printf("userName %v %v has booked %v tickets email: %v \n", firstName, lastName, userTickets, email)
			
			var firstname = getFirstName()
			fmt.Printf("firstName: %v\n", firstname)

			var noticks bool = remainingTickets == 0
			if noticks {
				// end programme
				fmt.Println("no more tickets")
				// break
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
			// continue
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
	wg.Wait()
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
	time.Sleep(2 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v %v", userTickets, firstName, lastName)
	fmt.Println("----------------")
	fmt.Printf("sending ticket: %v to email %v ", ticket, email)
	fmt.Println("----------------")
	wg.Done()
}

type Animal struct {
	name string ` required max: "100`
	origin string
}
type Bird struct {
	Animal
	spped float32
	canFly bool
}
func animal () {
	b:= Bird{}
	b.name = "emu"
	b.origin= "Aus"
	b.spped = 48
	b.canFly= false
	fmt.Println(b)
	fmt.Println(b.name)

	t:= reflect.TypeOf(Animal{})
	field, _:= t.FieldByName("name")
	fmt.Println(field.Tag)
	

var i interface{} = 1
switch i.(type){
case int:
	fmt.Println("its int")
case string:
	fmt.Println("its string")
case float32:
	fmt.Println("its float32")
}

for i, j := 0,0; i < 5; i,j = i+1,j+2 {
	fmt.Println(i, j)
}
}

