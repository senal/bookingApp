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

	/*
		what is  a pointer: variable that points to the memory address of a another value
	*/
	fmt.Println("Enter your first name?")
	fmt.Scan(&userName)
	fmt.Println("How many tickets you want to buy?")
	fmt.Scan(&userTickets)

	var bookings [10]string

	bookings[0] = "Ranga"
	bookings[1] = "Senal"
	bookings[2] = "Semira"

	fmt.Printf("Hey %v, Welcome to %v booking application\n", userName, confierenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, (remainingTickets - userTickets))
	fmt.Println("Get your tickets here to attend")

	fmt.Printf("User %v, has bought %v tickets\n", userName, userTickets)

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length %v", len(bookings))

	// define slice (which is a dynamic array which increase number of elements)
	subscribers := []string{}

	subscribers = append(subscribers, "Jude", "Asanka")

	fmt.Printf("Count of subscribers so far %v", len(subscribers))

}
