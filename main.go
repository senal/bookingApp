package main

import "fmt"

func main() {
	var confierenceName = "Go Conference"
	var userName string
	var userTickets int
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	/*
		how to concatenate, strings
	*/
	userName = "Ranga"
	userTickets = 5

	fmt.Printf("Hey %v, Welcome to %v booking application\n", userName, confierenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, (remainingTickets - userTickets))
	fmt.Println("Get your tickets here to attend")

	fmt.Printf("User %v, has bought %v tickets", userName, userTickets)

}
