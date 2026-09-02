// =============================================================================
// УРОК 1 — IsNegative
// НОВОЕ: package, import, функция, if/else, сравнение int
// =============================================================================

package piscine // имя пакета; все файлы piscine — одна библиотека

import "github.com/01-edu/z01" // подключаем z01 (см. import/z01.txt.go)

// nb int — параметр функции, целое число
func IsNegative(nb int) {
	if nb < 0 { // if / else — ветвление по условию
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
