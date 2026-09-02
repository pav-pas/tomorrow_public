// =============================================================================
// УРОК 15 — PrintStr
// НОВОЕ: тип string, range по строке, for _, значение
// =============================================================================

package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) { // string — строка символов
	// range по строке: _ пропускает индекс, word — текущий символ (rune)
	for _, word := range str {
		z01.PrintRune(word)
	}
}
