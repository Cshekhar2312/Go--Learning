package main

import (
"fmt"
)

func main(){

	const a = 100 //untyped 
	const b int =100 //typed

fmt.Printf(" a:= %T { %v }, \n b:= %T { %v }",a,a,b,b)



//exercise 2
	const m1='a'
	const m2='b'
	var prod int = m1*m2
	fmt.Printf("\n %T {%v}", prod,prod)

}