package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Conferencetickets", "remainingTicketsis", "conferenceName is%T,Conferencetickets,remainingTicketsis,conferenceName is")

	fmt.Println("Welcome to %v booking application.", conferenceName)
	fmt.Println(conferenceName)
	fmt.Println("We have total of %v tickets and %v and this many are still available.", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
	for {
		var firstName String
		var lastName String

		bookings := string

		var userName String
		var userTickets int
		//ask user for their name
		fmt.Println("Enter your first name.")
		fmt.Scan(&firstname)
		fmt.Println("Enter your last name.")
		fmt.Scan(&lastname)
		fmt.Println("Enter your e-mail address.")
		fmt.Scan(&emailaddress)
		fmt.Println("Enter your tickets.")
		fmt.Scan(&tickets)
		remainingTickets = remainingTickets - userTickets
		bookings = append(booking, firstName+" "+lastName)

		fmt.Printf("The whole array%v\n", bookings)
		fmt.Printf("The first value%v\n", bookings[0])
		fmt.Printf("Array type%v\n", bookings)
		fmt.Printf("Array length%v\n", len(bookings))

		fmt.Println("Thank you for booking your tickets. You will receive a confirmation e-mail at")
		userName = "Tom"
		userTickets = 2
		fmt.Println(firstName, userTickets)
		fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(bookings)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are all our bookings.", bookings)
	}
}
