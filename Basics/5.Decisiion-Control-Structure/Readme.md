### Go Decision Control Structures

### 1. `if` Statement

The `if` statement executes a block of code if the specified condition is true.

**Syntax:**

```go
goCopy code
if condition {
    // code to execute if condition is true
}

```

**Example:**

```go
goCopy code
package main

import "fmt"

func main() {
    x := 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    }
}

```

### 2. `if-else` Statement

The `if-else` statement executes one block of code if the condition is true and another block if the condition is false.

**Syntax:**

```go
goCopy code
if condition {
    // code to execute if condition is true
} else {
    // code to execute if condition is false
}

```

**Example:**

```go
goCopy code
package main

import "fmt"

func main() {
    x := 4
    if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is not greater than 5")
    }
}

```

### 3. `if-else if-else` Statement

The `if-else if-else` statement allows you to check multiple conditions sequentially.

**Syntax:**

```go
goCopy code
if condition1 {
    // code to execute if condition1 is true
} else if condition2 {
    // code to execute if condition2 is true
} else {
    // code to execute if both conditions are false
}

```

**Example:**

```go
goCopy code
package main

import "fmt"

func main() {
    x := 10
    if x < 5 {
        fmt.Println("x is less than 5")
    } else if x > 5 {
        fmt.Println("x is greater than 5")
    } else {
        fmt.Println("x is equal to 5")
    }
}

```

### 4. `switch` Statement

The `switch` statement is used to execute one block of code among many based on the value of a variable or an expression.

**Syntax:**

```go
goCopy code
switch expression {
case value1:
    // code to execute if expression == value1
case value2:
    // code to execute if expression == value2
default:
    // code to execute if none of the cases match
}

```

**Example:**

```go
goCopy code
package main

import "fmt"

func main() {
    day := "Tuesday"
    switch day {
    case "Monday":
        fmt.Println("Today is Monday")
    case "Tuesday":
        fmt.Println("Today is Tuesday")
    case "Wednesday":
        fmt.Println("Today is Wednesday")
    default:
        fmt.Println("Today is not Monday, Tuesday, or Wednesday")
    }
}

```

### 5. `switch` with `fallthrough`

In Go, `fallthrough` is used to transfer control to the next case within the same `switch`.

**Syntax:**

```go
goCopy code
switch expression {
case value1:
    // code to execute if expression == value1
    fallthrough
case value2:
    // code to execute if fallthrough from above case or expression == value2
default:
    // code to execute if none of the cases match
}

```

**Example:**

**Scenario:**
If you scored more than 60, you get ice cream. If more than 70, then ice cream + Mobile, if more than 80, ice cream + Mobile + Bike, if more than 90, then ice cream + Mobile + Bike + Goa Trip.

**Example:**

```go
goCopy code
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your marks: ")
    marksInput, _ := reader.ReadString('\n')

    // Trim the newline character from the input and convert it to an integer
    marksInput = strings.TrimSpace(marksInput)
    marks, err := strconv.Atoi(marksInput)
    if err != nil {
        fmt.Println("Invalid input. Please enter a valid number.")
        return
    }

    switch {
    case marks > 90:
        fmt.Println("You got Goa Trip")
        fallthrough
    case marks > 80:
        fmt.Println("You got Bike")
        fallthrough
    case marks > 70:
        fmt.Println("You got Mobile")
        fallthrough
    case marks > 60:
        fmt.Println("You got Ice Cream")
        fallthrough
    case marks > 40:
        fmt.Println("\n\nYou Passed")
    default:
        fmt.Println("You have failed")
    }
}

```

**Execution Flow:**

- If the user enters a value like `85`, the program will print:
    
    ```
    Copy code
    You got Bike
    You got Mobile
    You got Ice Cream
    You Passed
    
    ```
    
- If the user enters `95`, the program will print:
    
    ```
    Copy code
    You got Goa Trip
    You got Bike
    You got Mobile
    You got Ice Cream
    You Passed
    
    ```
    

This approach correctly implements the use of `fallthrough` and ensures that the marks are handled as integers for comparison.


### 6. `switch` with no condition

You can use a `switch` statement without an expression, which behaves like multiple `if-else` statements.

**Syntax:**

```go
goCopy code
switch {
case condition1:
    // code to execute if condition1 is true
case condition2:
    // code to execute if condition2 is true
default:
    // code to execute if none of the conditions are true
}

```

**Example:**

```go
goCopy code
package main

import "fmt"

func main() {
    x := 7
    switch {
    case x < 5:
        fmt.Println("x is less than 5")
    case x > 5:
        fmt.Println("x is greater than 5")
    default:
        fmt.Println("x is equal to 5")
    }
}

```

### Summary

- **if statement:** Checks a single condition.
- **if-else statement:** Checks a condition and provides an alternative if the condition is false.
- **if-else if-else statement:** Checks multiple conditions in sequence.
- **switch statement:** Selects a block of code to execute based on the value of an expression.
- **fallthrough in switch:** Allows execution to fall through to the next case.
- **switch with no condition:** Equivalent to multiple if-else statements, used for checking multiple conditions.

