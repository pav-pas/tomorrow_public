// =============================================================================
// УРОК 4 — PrintComb_v2 (альтернатива уроку 3)
// НОВОЕ: оператор &&, фильтр через if, ASCII-коды вместо символов
// =============================================================================

package piscine

import "github.com/01-edu/z01"

func PrintComb_v2() {
	for x := '0'; x <= '7'; x++ {
		for y := '1'; y <= '8'; y++ {
			for z := '2'; z <= '9'; z++ {
				// && — И: оба условия должны быть true
				if x < y && y < z {
					z01.PrintRune(x)
					z01.PrintRune(y)
					z01.PrintRune(z)
					if x < '7' {
						z01.PrintRune(44) // ',' в ASCII
						z01.PrintRune(32) // ' ' в ASCII
					}
				}
			}
		}
	}
	z01.PrintRune(10) // '\n' в ASCII
}
