// Author: etcodr
// Go50

// https://cs50.harvard.edu/x/2021/labs/1/hello/

// hello greets the user and converts age to days
package main

import "fmt"

func main() {
	fmt.Printf("Name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Printf("Age: ")
	var age int
	fmt.Scanf("%d", &age)

	fmt.Printf("Hello, %s! You're at least %d days old!\n", name, age*365)
}
