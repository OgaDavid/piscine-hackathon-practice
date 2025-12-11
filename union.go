package main

import (
	"os"
	"github.com/01-edu/z01"
)

func Union() {
	args := os.Args
	if len(args) != 3 { // program name + 2 strings
		z01.PrintRune('\n')
		return
	}

	s1 := args[1]
	s2 := args[2]
	seen := make(map[rune]bool)

	// helper function to process a string
	printUnion := func(s string) {
		for _, r := range s {
			if !seen[r] {
				seen[r] = true
				z01.PrintRune(r)
			}
		}
	}

	printUnion(s1)
	printUnion(s2)
	z01.PrintRune('\n')
}
