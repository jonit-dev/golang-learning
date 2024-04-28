package main

import "fmt"

func main() {

	var conferenceName = "GopherCon"
	const totalConferenceTickets = 50
	var remainingTickets = totalConferenceTickets

	fmt.Printf("Welcome to our %s's booking application!\n", conferenceName)
	fmt.Println("Another way to interpolate strings can be, for example,'" + conferenceName + "'")
	fmt.Printf("There are %d tickets available.", totalConferenceTickets)
	fmt.Println("Please select an option:")

}
