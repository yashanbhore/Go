package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "golang.org/x/text/number"
)

func main() {



	fmt.Print("Enter input")

	reader := bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n')

	// if we want to add rating = input+1 it is not possible

	// numberrating,err := strconv.ParseFloat(input,64) // ther is an issue with this it will also has \n for eg 4\n
	numberrating,err := strconv.ParseFloat(strings.TrimSpace(input),64) 


	if(err !=nil){
		fmt.Println(err)
	}else {
		
		fmt.Print(input, "  + 1 = ", numberrating+1 , "  ")
	}


}
