package main

import "fmt"

func main() {
	var confierenceName = "Go Conference"
	var userName = "Ranga"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	/*
		how to concatenate, strings
	*/
	fmt.Printf("Hey %v, Welcome to %v booking application\n", userName, confierenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}
