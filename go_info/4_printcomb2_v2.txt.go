// =============================================================================
// УРОК 6 — PrintComb2_v2 (альтернатива уроку 5)
// НОВОЕ: else if — вторая ветка условия
// =============================================================================

package piscine

import "github.com/01-edu/z01"

func PrintComb2_v2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {
					if a <= c && b < d {
						// Случай 1: одинаковые десятки, единицы растут (12 15)
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(32)
						z01.PrintRune(c)
						z01.PrintRune(d)
						if (a != '9' || b != '8') || (c != '9' || d != '9') {
							z01.PrintRune(44)
							z01.PrintRune(32)
						}
					} else if a < c && b >= d {
						// else if — иначе, если... Случай 2: десятки меньше (12 34)
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(32)
						z01.PrintRune(c)
						z01.PrintRune(d)
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
	z01.PrintRune(10)
}
