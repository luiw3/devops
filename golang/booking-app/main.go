package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const conferenceName = "X Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

	// naviganteAndCreateDirectory()
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Println("ERROR! WE DONT HAVE THIS MANY TICKETS")
			break
		}

		fmt.Printf("User %v %v booked %v tickets\n", firstName, lastName, userTickets)
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("There are now %v tickets remaining\n", remainingTickets)

		firstNames := []string{}

		for _, booking := range bookings {
			var name = strings.Fields(booking)
			firstNames = append(firstNames, name[0])
		}

		fmt.Printf("These are our bookings: %v \n", firstNames)
		dir, _ := os.UserHomeDir()
		os.WriteFile(dir+"/Desktop/devops-study/go-created-directory/bookings.txt", []byte(firstNames[0]), 0777)
		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("SOLD OUT")
			break
		}
	}

}

// func naviganteAndCreateDirectory() {
// 	err := os.Mkdir("/home/luisdev/Desktop/devops-study/go-created-directori/test", 0777)

// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
// }
