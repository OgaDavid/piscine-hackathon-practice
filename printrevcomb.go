package main

import "github.com/01-edu/z01"

func PrintRevComb() {
	for a := 9; a >= 2; a-- {
		for b := a - 1; b >= 1; b-- {
			for c := b - 1; c >= 0; c-- {
				z01.PrintRune(rune(a + '0'))
				z01.PrintRune(rune(b + '0'))
				z01.PrintRune(rune(c + '0'))
				if a+b+c != 3 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
