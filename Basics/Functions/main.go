package main

import "fmt"

func main() {
		x:=adddemo(2,3);
		fmt.Print("X = ",x)  // here X=0 - as it return nothing
		
		
		y:=add(2,3);
		fmt.Print("Y = ",y)  // here X=0 - as it return nothing
}

func adddemo(a int,b int)(result int) {
	fmt.Println("First line")
	return 
}

func add(a int, b int)(result int){
	return a+b;
}
