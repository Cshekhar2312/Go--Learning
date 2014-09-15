package main

import (
"fmt"
"math"
)

func main() {

	//exercise 1
    var a int
	var b string
	var c bool
	fmt.Printf("%T %v , %T  %v, %T  %v\n",a,a,b,b,c,c);
//exercise 2
	a= 10
	b= "string"
	c=true
	fmt.Printf("%T %v , %T  %v, %T  %v\n",a,a,b,b,c,c);
//exercise 3
	var d float32
	d= math.Pi
	fmt.Printf("%T  %v\n", d,d);

}

