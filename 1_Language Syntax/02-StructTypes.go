package main

import (
"fmt"
"time"
)

// exercise 1
type user struct{
	name string
    DOB time.Time
    address string
}

func main(){

	var a user
	fmt.Printf("%T  :  %v \n\n",a ,a )
	a=user{"Chandra Shekhar Mehta",time.Now(),"Army Institute of Technology, Pune"}
	fmt.Printf("%T  :  %v \n\n",a ,a )


	//exercise 2
	//Declare and initialize an anonymous struct type

/*
	b := struct{
		name string
		DOB time.Time
		address string
		anonymous bool
	}{
		name:"Chandra Shekhar Mehta",
		DOB: time.Now(),
		address: "Army Institute of Technology, Pune",
		anonymous: true,
	}


*/
	

	b := struct{
		name string
		DOB time.Time
		address string
		anonymous bool
	}{
		"Chandra Shekhar Mehta",time.Now(),"Army Institute of Technology, Pune",true,
	}


	

fmt.Printf("%T  :  %v \n\n",b ,b )
}