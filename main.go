package main

import "fmt"

func main() {
	var confierenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	/*
		how to concatenate, strings
	*/
	fmt.Println("Welcome to", confierenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

}
