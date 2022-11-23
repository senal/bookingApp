package main

import (
	"fmt"
	"strings"
)

func main() {
	const totalTickets int = 50
	var ticketsRemains int = totalTickets
	bookings := []string{}

	//city := "London"

	const conferenceName = "Beauty of GoLang"

	// greet users
	greetUsers(conferenceName)

	for {

		noTicketRemaining := ticketsRemains <= 0
		if noTicketRemaining {
			fmt.Printf("Our conference is booked up, come back next year \n")
			break
		}

		firstName, lastName, userTickets, city := getUserInputs()

		ticketsRemains = ticketsRemains - userTickets

		fmt.Printf("Remaning tickets %v \n", ticketsRemains)

		isValidName, isValidTicketNumber := validateUserInput(firstName, lastName, userTickets, ticketsRemains)

		if isValidName && isValidTicketNumber {
			bookings = append(bookings, firstName+" "+lastName)

			printStats(bookings, ticketsRemains)
			firstNames := getFirstNames(bookings)
			fmt.Printf("First names of the bookings %v \n", firstNames)

			// I need to know how to use string literals in GoLang
			// instead repeating the same printf(..) to print the message, I need to construct the message in advance, then print

			switch city {
			case "London":
				fmt.Printf("You selected %v city\n", city)
			case "Colombo", "Negombo":
				fmt.Printf("You selected Colombo and Negombo\n")
			case "Singapore":
				fmt.Printf("You selected %v city\n", city)
			default:
				fmt.Printf("You selected invalid city\n")
			}

		} else {
			if !isValidName {
				fmt.Printf("First or Last name you entered is too short \n")
				continue
			}

			if !isValidTicketNumber {
				fmt.Printf(" Please enter a valid ticket number \n")
				continue
			}
			fmt.Printf("Please make sure either user name or valid ticket number exists\n")
		}

	}
}

func getUserInputs() (string, string, int, string) {
	var firstName string
	var lastName string
	var city string
	var userTickets int

	fmt.Printf("Please enter you First name:\n")
	fmt.Scan(&firstName)
	fmt.Printf("And your Last name:\n")
	fmt.Scan(&lastName)
	fmt.Printf("How many tickets you want to buy?\n")
	fmt.Scan(&userTickets)
	fmt.Println("Please provide which city you interested")
	fmt.Scan(&city)

	return firstName, lastName, userTickets, city
}

func validateUserInput(firstName string, lastName string, userTickets int, ticketsRemains int) (bool, bool) {
	isValidUserName := isValidUserInfo(firstName, lastName)
	isValidTicketRequest := isValidTicketRequst(userTickets, ticketsRemains)

	return isValidUserName, isValidTicketRequest
}

func isValidUserInfo(firstName string, lastName string) bool {
	isValid := len(firstName) >= 2 && len(lastName) >= 2
	return isValid
}

func isValidTicketRequst(userTickets int, ticketsRemains int) bool {
	isValid := userTickets > 0 && userTickets <= ticketsRemains
	return isValid
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames
}

func greetUsers(confName string) {
	fmt.Printf("Welcome to our %v conference in 2022\n", confName)
}

func printStats(bookings []string, ticketsRemains int) {
	fmt.Printf("Total number of bookings %v \n", len(bookings))
	fmt.Printf("contains of bookings %v \n", bookings)
	fmt.Printf("Total remains tickets %v \n", ticketsRemains)
}
