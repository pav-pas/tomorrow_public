// =============================================================================
// УРОК 6 — PrintCombN
// НОВОЕ: bool, функция возвращает true/false, параметр n, разбор цифр числа
// =============================================================================
//
// n комбинаций из разных цифр: PrintCombN(3) → 012, 013, ... 789
//
// =============================================================================

package piscine

import "github.com/01-edu/z01"

func che(v int) {
	c := '0'
	if v == 0 {
		z01.PrintRune(c)
		return
	}
	for i := 1; i <= v%10; i++ {
		c++
	}
	for i := -1; i >= v%10; i-- {
		c++
	}
	if v/10 != 0 {
		che(v / 10)
	}
	z01.PrintRune(c)
}

// bool — тип «да/нет»; функция ВОЗВРАЩАЕТ значение через return
func my_check(p int) bool {
	for {
		if p == 0 {
			break
		}
		// Извлекаем цифры: p%10 — последняя, (p/10)%10 — предпоследняя
		if p/10 != 0 && p%10 <= ((p/10)%10) {
			return false // цифры не строго растут
		}
		p /= 10 // p /= 10 то же что p = p / 10
	}
	return true
}

func check_ok(p int) bool {
	if p < 9 {
		return true
	}
	if p%10 == 9 {
		for {
			if p == 0 {
				break
			}
			if p/10 != 0 && p%10 != ((p/10)%10)+1 {
				return true
			}
			p /= 10
		}
		return false // последняя комбинация — запятую не ставим
	}
	return true
}

func PrintCombN(n int) {
	mx_ln := 1
	for i := 2; i <= n; i++ {
		mx_ln *= 10 // 10^(n-1)
	}
	for i := mx_ln / 10; i < mx_ln; i++ {
		if my_check(i) == true {
			if mx_ln >= 10 {
				z01.PrintRune('0')
			}
			che(i)
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	for i := mx_ln; i <= mx_ln*9; i++ {
		if my_check(i) == true {
			che(i)
			if check_ok(i) == true {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
