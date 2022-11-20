package main

import (
	"fmt"
	"strings"
)

func main() {
	const totalTickets int = 50
	var ticketsRemains int = totalTickets
	bookings := []string{}

	for {
		var firstName string
		var lastName string
		var userTickets int
		fmt.Printf("Please enter you First name:\n")
		fmt.Scan(&firstName)
		fmt.Printf("And your Last name:\n")
		fmt.Scan(&lastName)
		fmt.Printf("How many tickets you want to buy?\n")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidTicketNumber := userTickets > 0 && userTickets <= ticketsRemains

		if isValidName && isValidTicketNumber {
			bookings = append(bookings, firstName+" "+lastName)

			ticketsRemains = ticketsRemains - userTickets

			firstNames := []string{}

			// _ in Go is to indicate not used variable
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				var firstName = names[0]
				firstNames = append(firstNames, firstName)
			}

			fmt.Printf("Total number of bookings %v \n", len(bookings))
			fmt.Printf("contains of bookings %v \n", bookings)
			fmt.Printf("Total remains tickets %v \n", ticketsRemains)

			fmt.Printf("First names of the bookings %v \n", firstNames)

			noTicketRemaining := ticketsRemains <= 0
			if noTicketRemaining {
				fmt.Printf("Our conference is booked up, come back next year \n")
				break
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
