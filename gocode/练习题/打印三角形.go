package main

import "fmt"

// *
// **
// ***
// ****
// *****
// ****
// ***
// **
// *
func main() {
	for i := 0; i < 10; i++ {
		if i < 5 {
			for j := 0; j < i; j++ {
				fmt.Print("*")
			}
		} else {
			for j := 0; j < 10-i; j++ {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}
}
