// =============================================================================
// УРОК 17 — StrRev
// НОВОЕ: return string, индексация среза s[i], string([]rune)
// =============================================================================

package piscine

func my_StrLen(str string) int {
	astr := []rune(str)
	a := 1
	for index := range astr {
		a = index
	}
	return a + 1
}

func StrRev(s string) string { // функция ВОЗВРАЩАЕТ string (не печатает)
	a := my_StrLen(s)
	a = a - 1

	a_s := []rune(s)
	b_s := []rune(s)

	for index := range a_s {
		b_s[index] = a_s[a-index] // ИНДЕКСАЦИЯ среза: a_s[номер]
	}

	return string(b_s) // срез рун обратно в строку
}

// StrRev("Hi") → "iH"
