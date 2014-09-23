package main

import (
		"fmt"
		)


type car struct{
	brand string
	licns_no string
	noOfDoors int
	modelNo string
}

func  change(y *car){
	y.brand= "Ford"
	//return *y
}

func main(){

	//exercise 1
	var x int =20
	fmt.Printf("Value at address %v is %v \n",&x,x)


	//exercise 1
	x= 42
	var ptrTo_x *int =&x;
	fmt.Printf("Value at address %v is %v pointing to type \" %T \" initialized to value %v \n",&ptrTo_x,ptrTo_x,*ptrTo_x,*ptrTo_x)

	//exercise 3
	y := car{
		brand: "Audi",
		licns_no: "MH 785 6899",
		noOfDoors: 2,
		modelNo: "Audi R8",
	}

	fmt.Printf("%v",y)
	change(&y)

	fmt.Printf("%v",y)
}
