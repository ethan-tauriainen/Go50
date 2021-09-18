// Author: etcodr
// Go50

// https://cs50.harvard.edu/x/2021/psets/1/cash/

// cash determines how much change is owed,
// then prints the minimum number of coins
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var dollars float64

	// do-while loop
	for {
		dollars = getFloat("Change owed: ")
		if dollars > 0 {
			break
		}
	}

	cents := int(math.Round(dollars * 100))
	coins := compute_coins(cents)

	fmt.Printf("%d\n", coins)
}

func compute_coins(cents int) int {
	var coins int

	// while loop
	for cents > 0 {
		if cents >= 25 {
			cents -= 25
			coins++
		} else if cents >= 10 {
			cents -= 10
			coins++
		} else if cents >= 5 {
			cents -= 5
			coins++
		} else {
			cents -= 1
			coins++
		}
	}

	return coins
}

// my own implementation of the CS50 get_float function
func getFloat(prompt string) float64 {
	var str string

	for {
		fmt.Printf("%s", prompt)
		fmt.Scanln(&str)

		f, err := strconv.ParseFloat(str, 64)

		if err == nil {
			return f
		}
	}
}
