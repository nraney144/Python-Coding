package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	remainingTickets = -1
	fmt.printf("Conferencetickets", "remainingTicketsis", "conferenceName is%T,Conferencetickets,remainingTicketsis,conferenceName is")

	fmt.Println("Welcome to %v booking application.", conferenceName)
	fmt.Println(conferenceName)
	fmt.println("We have total of %v tickets and %v and this many are still available.", conferenceTickets, remainingTickets)
	fmt.println("Get your tickets here to attend.")

	var userName String
	var userTickets int
	//ask user for their name
	fmt.println("Enter your first name.")
	fmt.Scan(&firstname)
	fmt.println("Enter your last name.")
	fmt.Scan(&lastname)
	fmt.println("Enter your e-mail address.")
	fmt.Scan(&emailaddress)
	fmt.println("Enter your tickets.")
	fmt.Scan(&tickets)
	remainingTickets = remainingTickets - userTickets
	fmt.println("Thank you for booking your tickets. You will receive a confirmation e-mail at")
	userName = "Tom"
	userTickets = 2
	fmt.printLn(firstName, userTickets)
	fmt.printf("%v tickets remaining for %v", remainingTickets, conferenceName)
}
