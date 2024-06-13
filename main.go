package main 


import(
	"fmt"
	"strings"

	

) 



func main(){


for{
	// conferenceName :="Go Conference ";
	// const conferenceTickets int=50;
	var remainingTickets uint=50;




	// var bookings=[50]string{"Nana","david","chuks",}
	

	// fmt.Printf("conferenceName is %T conferenceTickets is %T  remainingTickets is a %T",conferenceName,conferenceTickets,remainingTickets);
 bookings :=[]string{}


	// fmt.Printf("Welcome to %v booking application\n",conferenceName);
	// fmt.Printf("we have a total of %v  tickets and  remaining %v",remainingTickets,conferenceTickets)
	// fmt.Println("Get your ticket here to attend ");
	

// bookings[0]="Nana"
// bookings[1]="chuks"

for  remainingTickets > 0 && len(bookings)<50{


var firstName string;
var lastName string;
var email string;
var tickets uint;

fmt.Printf("Enter your firstName\n");
fmt.Scan(&firstName);

fmt.Printf("Enter your Lastname");


fmt.Scan(&lastName);


fmt.Printf("Enter your email");
fmt.Scan(&email);


fmt.Printf("Enter the number of ticket");
fmt.Scan(&tickets);

  isValidName :=len(firstName)>= 2  && len(lastName)>=2;
  isValidEmail:=strings.Contains(email,"@")
  isValidTicketNumber:=tickets > 0 && tickets <= remainingTickets;


//   isInValidCity:= city!="singpore" && city != "London";
;
if isValidName && isValidEmail && isValidTicketNumber{
	remainingTickets=remainingTickets-tickets;
bookings=append(bookings,firstName +" "+lastName)


// fmt.Printf("The whole slice:%v\n ",bookings);
// fmt.Printf("The first value :%v\n ",bookings[0]);
// fmt.Printf("slice type:%T\n ",bookings);
// fmt.Printf("slice length:%v\n ",len(bookings));




fmt.Printf("Thank you %v %v for getting %v you will get a confirmation mail at %v\n",firstName,lastName,tickets,email,);

fmt.Printf("there is only %v remaining",remainingTickets,);

fmt.Printf("These are all our bookings:%v\n",bookings);


firstNames:=[]string{}

for _,booking:=range bookings{
	var names=strings.Fields(booking)



	

	var firstName=names[0]
	firstNames=append(firstNames,firstName)
}

fmt.Printf("the first name of the booking are %v\n ",firstNames);

noTicketRemaining:=remainingTickets==0

if noTicketRemaining{

	fmt.Print("The conference ticket are booked out already we are so sorry ")
	break;

}






}else{
	fmt.Printf("Your input data is invalid\n");
	break;

}

	


}

}

}