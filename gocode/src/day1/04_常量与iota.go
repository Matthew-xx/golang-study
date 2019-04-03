package main

import "fmt"

func main() {
	const (
		b, c = iota, iota
		a    = iota
		e

		f
	)
	const g = iota
	h := 10
	h++
	fmt.Print(b, c, a, e, f, g, h)
}
