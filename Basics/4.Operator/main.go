//------------------------------------Operators in Go-------------------------------

package main

import "fmt"

func main() {
	arithmetaic(23, 54)
	relational(11, 11)
	logical(true, false)
	bitwise(23, 25)
	// assignment();
	miscellaneous()
}

func arithmetaic(x int, y int) {
	a := x
	b := y

	fmt.Println("Arithmetaic :- ")

	fmt.Printf("Addition of %d + %d  = %d", a, b, a+b)
	fmt.Printf("\nSubtraction of %d - %d  = %d", a, b, a-b)
	fmt.Printf("\nMultiplication of %d * %d  = %d", a, b, a*b)
	fmt.Printf("\nDivision of %d / %d  = %d", a, b, a/b)
	fmt.Printf("\nModulus of %d %% %d  = %d", a, b, a%b)

}

func relational(x int, y int) {
	fmt.Println("\n\nRelational Operator :- ")

	fmt.Printf("Greator than - Is %d > %d  = %t \n", x, y, x > y)
	fmt.Printf("Less than - Is %d < %d  = %t \n", x, y, x > y)
	fmt.Printf("Equals - Is %d == %d  = %t \n", x, y, x == y)
	fmt.Printf("Not Equals - Is %d != %d  = %t \n", x, y, x != y)
	fmt.Printf("Greator than Equals - Is %d >= %d  = %t \n", x, y, x >= y)
	fmt.Printf("Less than Equals - Is %d <= %d  = %t \n", x, y, x <= y)

}

func logical(a bool, b bool) {
	fmt.Println("\nLogical Operators :- ")
	fmt.Printf("%t && %t = %t\n", a, b, a && b)
	fmt.Printf("%t || %t = %t\n", a, b, a || b)
	fmt.Printf("!%t = %t\n", a, !a)
}

func bitwise(a int, b int) {
	fmt.Println("\nBitwise Operator :- ")
	fmt.Printf("%d & %d = %d\n", a, b, a&b)
	fmt.Printf("%d | %d = %d\n", a, b, a|b)
	fmt.Printf("%d ^ %d = %d\n", a, b, a^b)
	fmt.Printf("%d << 1 = %d\n", a, a<<1)
	fmt.Printf("%d >> 1 = %d\n", a, a>>1)
}

func miscellaneous() {
	var a int = 12
	var ptr *int

	ptr = &a

	// without var using svd
	// ptr:=a
	// a:=12
	fmt.Println("\nPointer and destructuring :- ")

	fmt.Printf("Address of a : %p \n", &a)
	fmt.Printf("Value of ptr(a address) : %p \n", ptr)
	fmt.Printf("Value pointed to ptr(a)  : %d \n", *ptr)
}
