// Author: etcodr
// Go50

// https://cs50.harvard.edu/x/2021/labs/1/population/

// population calculates the number of years required for the
// population to grow from the start size to the end size
package main

import "fmt"

func main() {
	// prompt for start size
	var startSize = -1
	for startSize < 9 {
		fmt.Printf("Start size: ")
		fmt.Scanf("%d", &startSize)
	}

	// prompt for end size
	var endSize = -1
	for endSize < startSize {
		fmt.Printf("End size: ")
		fmt.Scanf("%d", &endSize)
	}

	// calculate number of years until we reach threshold
	var years, counter int = 0, startSize
	for counter < endSize {
		counter += (counter / 3) - (counter / 4)
		years++
	}

	// print number of years
	fmt.Printf("Years: %d\n", years)
}
