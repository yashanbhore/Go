package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	greet := "Heloo "
	fmt.Print(greet)

	reader  := bufio.NewReadWriter(os.Stdin)
}