// =============================================================================
// УРОК 19 — BasicAtoi2
// НОВОЕ: флаг-валидатор (checker bool), byte(), проверка ASCII 48..57
// =============================================================================

package piscine

func BasicAtoi2(s string) int {
	o_number := 0
	c := 0
	checker := true // флаг: пока true — строка валидна
	a_s := []rune(s)

	for _, word := range a_s {
		if byte(word) >= 48 && byte(word) <= 57 { // byte() — приведение к байту для сравнения с ASCII
			for i := '0'; i < word; i++ {
				c++
			}
			o_number = o_number*10 + c
			c = 0
		} else {
			checker = false
		}
	}

	if checker {
		return o_number
	}
	return 0 // любой мусор в строке → 0
}
