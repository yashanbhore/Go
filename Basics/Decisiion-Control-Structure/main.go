package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("X is less than 5")
	}

	y := 7

	if y > 10 {
		fmt.Println("Y is greater than 10")
	} else if y > 5 {
		fmt.Println("Y is greater than 5 and less than 10")
	}

	//switch-case
	var day string = "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	default:
		fmt.Println("Today is Sunday")
	}

	// switch-case with fallthrough and conditions- (it is used so that even if the case is correct it will continue its execution , it will not get terminated)
	// if you scored more than 60 you get icream. if>70 then icream + Mobile , if > 80 icream + Mobile + Bike , if>90 then icream + Mobile +Bike + Goa Trip

	getmarks := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your Marks")
	strmarks, _ := getmarks.ReadString('\n')

	// Trim the newline character from the input and convert it to an integer
	strmarks = strings.TrimSpace(strmarks)
	marks, err := strconv.Atoi(strmarks) //Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.

	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	switch {
	case marks > 40:
		fmt.Println("\n\nYou  have Passed")
		fallthrough
	case marks > 60:
		fmt.Println("You got Icream")
		fallthrough
	case marks > 70:
		fmt.Println("You got Mobile")
		fallthrough
	case marks > 80:
		fmt.Println("You got Bike")
		fallthrough
	case marks > 90:
		fmt.Println("You got Goa Trip")

	default:
		fmt.Println("You have failed")
	}

}
