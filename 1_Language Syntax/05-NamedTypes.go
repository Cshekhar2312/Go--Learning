package main

import (
	"fmt"
)

type Counter int
func main(){
	var x Counter

	fmt.Printf("type: %T \n value: %v \n",x,x)

	// x= 10  won't work because x is a counter type variable and 10 is int type

	x= Counter(10)
	fmt.Printf("type: %T \n value: %v \n",x,x)



	y:=100
	//x=y also won't work b/c of the above reason
	x=Counter(y)
	fmt.Printf("type: %T \n value: %v",x,x)
}