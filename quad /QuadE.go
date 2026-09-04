package piscine

import (
	"github.com/01-edu/z01"
)

func quadconditionE(x int, y int, i int, j int) {
	if (i == 0 && j == 0) || (i == x-1 && j == y-1) && (y != 1 && x != 1) {
		z01.PrintRune('A')
	} else if (i == 0 && j == y-1) || (i == x-1 && j == 0) {
		z01.PrintRune('C')
	} else if (i == 0 || i == x-1) && (j != 0 && j != y-1) {
		z01.PrintRune('B')
	} else if (j == 0 || j == y-1) && (i != 0 && i != x-1) {
		z01.PrintRune('B')
	} else {
		z01.PrintRune(' ')
	}
}

func QuadE(x, y int) {
	if x > 0 && y > 0 {
		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				quadconditionE(x, y, i, j)
			}
			z01.PrintRune('\n')
		}
	}
}
