package main 


import "fmt"



func main(){
	conferenceName :="Go Conference ";
	const conferenceTickets int=50;
	var remainingTickets uint=50;
	

	fmt.Printf("conferenceName is %T conferenceTickets is %T  remainingTickets is a %T",conferenceName,conferenceTickets,remainingTickets);


	fmt.Printf("Welcome to %v booking application\n",conferenceName);
	fmt.Printf("we have a total of %v  tickets and  remaining %v",remainingTickets,conferenceTickets)
	fmt.Println("Get your ticket here to attend ");
	
var userName string

var userTickets int


userName="chuks";
userTickets=2
fmt.Println(userName)
fmt.Printf("usern %v booked %v tickets",userName,userTickets,)


}
