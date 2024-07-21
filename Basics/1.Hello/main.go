package main

import "fmt"

func main(){
	// fmt.Print("Hello World")

	
	sum(1,2,3,4,5)
}

func sum (nums ...int){
	total:=0

	for i :=0 ; i<len(nums); i++{
		total+=nums[i]
		fmt.Println(nums[i]," \n")
	}

	fmt.Println(total)

}

