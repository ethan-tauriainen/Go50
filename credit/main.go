/* Author: etcodr
   Go50

   https://cs50.harvard.edu/x/2021/psets/1/credit/

   credit prompts the user for a credit card number,
   then reports whether it is a valid American Express,
   MasterCard, or Visa card number */
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var num int
	for {
		num = getInt("Number: ")
		if num > 0 {
			break
		}
	}
	digits := int(math.Log10(float64(num)) + 1)

	if validate(num, digits) {
		fmt.Println(getCardType(num, digits))
	} else {
		fmt.Printf("INVALID\n")
	}
}

/* Using Luhn's Algorithm determine whether or not the card number
   entered is valid */
func validate(num, digits int) bool {
	var x, y int
	for i := 1; i <= digits; i++ {
		d := num % 10
		num /= 10
		if i%2 != 0 {
			// Sum up the digits not being multiplied by 2
			y += d
		} else {
			/* Multiply every other digit by 2, starting with the number's
			   second-to-last digit, and then add those product's digits
			   together */
			d *= 2
			productDigits := int(math.Log10(float64(d)) + 1)
			if productDigits > 1 {
				for d > 0 {
					d2 := d % 10
					d /= 10
					x += d2
				}
			} else {
				x += d
			}
		}
	}
	/* If the total's last digit is 0 (or, put more formally, if the
	   total modulo 10 is congruent to 0), the number is valid! */
	if (x+y)%10 == 0 {
		return true
	}
	return false
}

// Determine type of card (VISA, MASTERCARD, AMEX)
func getCardType(num, digits int) string {
	var x, y int
	if digits == 15 {
		x = num / (int(math.Pow(10, 13)))
		if x == 34 || x == 37 {
			return "AMEX"
		} else {
			return "INVALID"
		}
	} else if digits == 16 {
		x = num / (int(math.Pow(10, 14)))
		y = num / (int(math.Pow(10, 15)))
		if x == 51 || x == 52 || x == 53 || x == 54 || x == 55 {
			return "MASTERCARD"
		} else if y == 4 {
			return "VISA"
		} else {
			return "INVALID"
		}
	} else if digits == 13 {
		x = num / (int(math.Pow(10, 12)))
		if x == 4 {
			return "VISA"
		} else {
			return "INVALID"
		}
	}
	return "INVALID"
}

// My own implementation of the CS50 get_int function
func getInt(prompt string) int {
	var str string
	for {
		fmt.Printf("%s", prompt)
		fmt.Scanln(&str)
		i, err := strconv.Atoi(str)
		if err == nil {
			return i
		}
	}
}
