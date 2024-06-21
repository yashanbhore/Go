package main

import "fmt"

func main() {
	addy := add(2, 3)
	muly := mul(2, 3)
	sumofnos := sum(1, 2, 3, 4, addy, muly)

	//anonymous function
	sub := func(a int, b int) int {
		return a - b
	}

	greet("Welcome today we are learning about types of function in Go")

	fmt.Println("Add = ", addy)
	fmt.Println("Subtract = ", sub(6, 3))
	fmt.Println("Product = ", muly)
	fmt.Println("Sum of numbers = ", sumofnos)

}

// this will not work we have to specify the return type
// func add(a int, b int){
// 	ans:=a+b
// 	return ans
// }

func greet(mess string) {
	fmt.Println(mess)

}

func add(a int, b int) int {
	ans := a + b
	return ans
}

// Function with named return values
func mul(a int, b int) (result int) {
	return a * b
}

// function takes variable legth arguments
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
