package stringlearning

import (
	"basics-app/src/stringutils"
	"fmt"
	"strconv"
)

func ExecStringLearning() {

	conferenceName := "GopherCon" // syntax sugar for var conferenceName = "GopherCon"
	const totalConferenceTickets = 50
	remainingTickets := totalConferenceTickets

	fmt.Printf("*** Welcome to our %v's booking application! ***\n\n", conferenceName)
	fmt.Println("- Another way to interpolate strings can be, for example,'" + conferenceName + "'")
	fmt.Printf("- There are %v/%v tickets available.\n\n", remainingTickets, totalConferenceTickets)

	userName := stringutils.GetInput("Please enter your name:")
	lastName := stringutils.GetInput("Please enter your last name:")
	email := stringutils.GetInputEmail("Please enter your email:")

	purchaseTicketsInput := stringutils.GetInput("How many tickets would you like to purchase?")

	fmt.Printf("Purchasing %v tickets...\n", purchaseTicketsInput)

	purchaseTickets, err := strconv.Atoi(purchaseTicketsInput)

	if err != nil {
		fmt.Println("Invalid input. Please try again.")
		return
	}

	if purchaseTickets > remainingTickets {
		fmt.Printf("Sorry, there are only %v tickets left. Try again!\n", remainingTickets)
		return
	}

	remainingTickets -= purchaseTickets

	fmt.Printf("Thank you %v %v (%v)! You have successfully purchased %v ticket(s) for %v.\n", userName, lastName, email, purchaseTickets, conferenceName)

	fmt.Printf("- There are %v/%v tickets available.\n\n", remainingTickets, totalConferenceTickets)

}
