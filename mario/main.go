// Author: etcodr
// Go50

// https://cs50.harvard.edu/x/2021/psets/1/mario/more/

// mario recreates adjacent pyramids of blocks using hashes for bricks
package main

import (
	"fmt"
)

func main() {
	var height int
	for {
		fmt.Printf("Height: ")
		fmt.Scanf("%d", &height)
		if height >= 1 && height <= 8 {
			break
		}
	}

	for i := 0; i < height; i++ {
		// print the spaces
		for j := height - 1; j > i; j-- {
			fmt.Printf(" ")
		}
		// print left side hashes
		for k := 0; k <= i; k++ {
			fmt.Printf("#")
		}
		// print gap
		fmt.Printf("  ")
		// print right side hashes
		for l := 0; l <= i; l++ {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}
