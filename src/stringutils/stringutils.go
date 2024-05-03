package stringutils

import (
	"fmt"
	"regexp"
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
