package main

import "fmt"

func main() {
	const (
		b, c = iota, iota
		a    = iota
		e

		_
		f
		g = "a"
		h
		i = iota
		j
	)
	const k = iota
	// const l
	const m = iota
	// const n

	fmt.Print(b, c, a, e, f, g, h, i, j, k, m)
}
