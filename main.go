package main 


import "fmt"



func main(){
	conferenceName :="Go Conference ";
	const conferenceTickets int=50;
	var remainingTickets uint=50;

	var bookings=[50]string{"Nana","david","chuks",}
	

	fmt.Printf("conferenceName is %T conferenceTickets is %T  remainingTickets is a %T",conferenceName,conferenceTickets,remainingTickets);


	fmt.Printf("Welcome to %v booking application\n",conferenceName);
	fmt.Printf("we have a total of %v  tickets and  remaining %v",remainingTickets,conferenceTickets)
	fmt.Println("Get your ticket here to attend ");
	
var firstName string
var lastName string;
var email string;
var tickets uint;

fmt.Printf("Enter your firstName");
fmt.Scan(&firstName);

fmt.Printf("Enter your secondname");
fmt.Scan(&lastName);

fmt.Printf("Enter your email");
fmt.Scan(&email);


fmt.Printf("Enter the number of ticket");
fmt.Scan(&tickets);

remainingTickets=remainingTickets-tickets;

fmt.Printf("Thank you %v %v for getting %v you will get a confirmation mail at %v\n",firstName,lastName,tickets,email,);

fmt.Printf("there is only %v remaining",remainingTickets,);




}