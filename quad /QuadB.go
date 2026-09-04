package piscine

import (
	"github.com/01-edu/z01"
)

func quadconditionB(x int, y int, i int, j int) {
	if (i == 0 && j == 0) || (i == x-1 && j == y-1) && (y != 1 && x != 1) {
		z01.PrintRune('/')
	} else if (j == y-1 && i == 0) || (j == 0 && i == x-1) {
		z01.PrintRune('\\')
	} else if (i == 0 || i == x-1) && (j != 0 || j != y-1) {
		z01.PrintRune('*')
	} else if (j == 0 || j == y-1) && (i != 0 || i != x-1) {
		z01.PrintRune('*')
	} else {
		z01.PrintRune(' ')
	}
}

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				quadconditionB(x, y, i, j)
			}
			z01.PrintRune('\n')
		}
	}
}
