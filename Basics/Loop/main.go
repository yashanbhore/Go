package main

import "fmt"

func main() {

	// go has only for loop which is similar to c
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	// Conditional for loop {similar to while}
	i := 6

	for i >= 1 {
		fmt.Println(i)
		i--
	}

	//for loop with range

	nums := []int{1, 2, 3, 4, 5}

	for index, val := range nums {
		fmt.Printf("Index : %d  Value : %d", index, val)
	}

	//break and continue

	for i := 0; i < 10; i++ {
		if i > 7 {
			break
		}
		if i > 3 {
			continue
		}
		fmt.Print("-")
	}
}
