// =============================================================================
// УРОК 5 — PrintNbr
// НОВОЕ: рекурсия, switch, return, n/10 и n%10, вспомогательные функции
// =============================================================================

package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n == -n { // крайний случай: минимальный int, -n переполняется
			printMinInt(n)
			return // return — немедленный выход из функции
		}
		n = -n
	}
	printDigits(n)
}

func printMinInt(n int) {
	if n <= -10 {
		printMinInt(n / 10) // РЕКУРСИЯ: функция вызывает сама себя
	}
	// switch — выбор по значению (как цепочка if/else if)
	switch n % 10 { // % — остаток от деления (= последняя цифра)
	case 0:
		z01.PrintRune('0')
	case -1:
		z01.PrintRune('1')
	case -2:
		z01.PrintRune('2')
	case -3:
		z01.PrintRune('3')
	case -4:
		z01.PrintRune('4')
	case -5:
		z01.PrintRune('5')
	case -6:
		z01.PrintRune('6')
	case -7:
		z01.PrintRune('7')
	case -8:
		z01.PrintRune('8')
	case -9:
		z01.PrintRune('9')
	}
}

func printDigits(n int) {
	if n >= 10 {
		printDigits(n / 10) // / — целочисленное деление (убирает последнюю цифру)
	}
	switch n % 10 {
	case 0:
		z01.PrintRune('0')
	case 1:
		z01.PrintRune('1')
	case 2:
		z01.PrintRune('2')
	case 3:
		z01.PrintRune('3')
	case 4:
		z01.PrintRune('4')
	case 5:
		z01.PrintRune('5')
	case 6:
		z01.PrintRune('6')
	case 7:
		z01.PrintRune('7')
	case 8:
		z01.PrintRune('8')
	case 9:
		z01.PrintRune('9')
	}
}

// Рекурсия printDigits(123): сначала 1, потом 2, потом 3 → на экране "123"
