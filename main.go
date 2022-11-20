package main

import "fmt"

func main() {
	const totalTickets int = 50
	var ticketsRemains int = totalTickets
	bookings := []string{}

	for ticketsRemains > 0 {
		var userName string
		var userTickets int
		fmt.Printf("Please enter you First name:\n")
		fmt.Scan(&userName)
		fmt.Printf("How many tickets you want to buy?\n")
		fmt.Scan(&userTickets)
		bookings = append(bookings, userName)

		ticketsRemains = ticketsRemains - userTickets

		fmt.Printf("Total number of bookings %v \n", len(bookings))
		fmt.Printf("Total remains tickets %v \n", ticketsRemains)

	}
}
