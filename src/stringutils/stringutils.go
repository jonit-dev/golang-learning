package stringutils

import (
	"fmt"
	"regexp"
	"strconv"
)

func GetInput(question string) string {
	var inputVar string

	fmt.Println(question)
	var _, hasError = fmt.Scanln(&inputVar)

	if hasError != nil {
		fmt.Println("An error occurred. Please try again.")
		return GetInput(question)
	}

	return inputVar

}

func GetInputEmail(question string) string {

	email := GetInput(question)

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !emailRegex.MatchString(email) {
		fmt.Println("Invalid email address. Please try again.")
		return GetInputEmail(question)
	}

	return email
}

func ExecStringLearning() {

	conferenceName := "GopherCon" // syntax sugar for var conferenceName = "GopherCon"
	const totalConferenceTickets = 50
	remainingTickets := totalConferenceTickets

	fmt.Printf("*** Welcome to our %v's booking application! ***\n\n", conferenceName)
	fmt.Println("- Another way to interpolate strings can be, for example,'" + conferenceName + "'")
	fmt.Printf("- There are %v/%v tickets available.\n\n", remainingTickets, totalConferenceTickets)

	userName := GetInput("Please enter your name:")
	lastName := GetInput("Please enter your last name:")
	email := GetInputEmail("Please enter your email:")

	purchaseTicketsInput := GetInput("How many tickets would you like to purchase?")

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
