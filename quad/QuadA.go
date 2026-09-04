package piscine

import (
	"github.com/01-edu/z01"
)

func quadconditionA(x int, y int, i int, j int) {
	if (i == 0 || i == x-1) && (j == 0 || j == y-1) {
		z01.PrintRune('o')
	} else if (i == 0 || i == x-1) && (j != 0 || j != y-1) {
		z01.PrintRune('|')
	} else if (j == 0 || j == y-1) && (i != 0 || i != x-1) {
		z01.PrintRune('-')
	} else {
		z01.PrintRune(' ')
	}
}

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				quadconditionA(x, y, i, j)
			}
			z01.PrintRune('\n')
		}
	}
}
