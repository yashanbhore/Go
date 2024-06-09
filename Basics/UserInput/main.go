package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	greet := "Heloo "
	fmt.Print(greet)

	// Buffio

	reader := bufio.NewReader(os.Stdin)

	// error or syntax
	na, _ := reader.ReadString('\n') // this returns 2  values

	fmt.Print(na)
}
