#!/usr/bin/env bash
# Первая строка говорит системе: запускай этот файл через bash # УДАЛИТЬ

# Переходим в папку mystery рядом со скриптом; если папки нет — выходим с ошибкой # УДАЛИТЬ
cd "$(dirname "$0")/mystery" || exit 1

# Из файла crimescene берём улики (строки с CLUE) и вытаскиваем минимальный рост # УДАЛИТЬ
# grep ищет текст, grep -o вырезает из строки только совпадение (например 6'), а не всю строку # УДАЛИТЬ
# min_height=$(grep "CLUE" crimescene | grep "at least" | grep -o "6'") # УДАЛИТЬ
min_height=$(grep "CLUE" crimescene | grep "at least" | grep -o "[0-9]'")

# В улике упомянуты карты клубов; AAA и Delta_SkyMiles знаем по имени # УДАЛИТЬ
# Библиотеку и музей ищем в папке memberships через ls и grep # УДАЛИТЬ
library=$(ls memberships | grep -i library | head -n 1)
museum=$(ls memberships | grep -i museum | grep -i bash | head -n 1)

# Ищем всех Annabel в people; grep "	F	" оставляет только женщин (свидетельница из кафе) # УДАЛИТЬ
# while read — читаем каждую строку: имя, пол, возраст, адрес (разделены табом) # УДАЛИТЬ
grep "Annabel" people | grep "	F	" | while IFS=$'\t' read -r pname gender age address; do

	# Из адреса "Buckingham Place, line 179" достаём улицу и номер строки # УДАЛИТЬ
	street="${address%, line *}"
	line="${address##*, line }"
	# Пробелы в названии улицы меняем на _ (Buckingham Place -> Buckingham_Place) # УДАЛИТЬ
	street_file="streets/${street// /_}"

	# hint5: head берёт первые N строк файла улицы, tail — последнюю из них # УДАЛИТЬ
	# Там будет что-то вроде: SEE INTERVIEW #699607 # УДАЛИТЬ
	interview_ref=$(head -n "$line" "$street_file" | tail -n 1)
	# Из этой строки вытаскиваем только цифры — номер файла интервью # УДАЛИТЬ
	num=$(echo "$interview_ref" | grep -o '[0-9]*')

	# Настоящая свидетельница описала машину: номер L337, Honda, синий # УДАЛИТЬ
	# Если в интервью нет этих слов — это не та Annabel, пропускаем (continue) # УДАЛИТЬ
	cat "interviews/interview-$num" | grep "L337" > /dev/null || continue
	cat "interviews/interview-$num" | grep "Honda" > /dev/null || continue
	cat "interviews/interview-$num" | grep -i "Blue" > /dev/null || continue

	# hint6-7: ищем машины в vehicles цепочкой grep с -A/-B (строки вокруг совпадения) # УДАЛИТЬ
	# Фильтруем: номер L337, Honda, Blue, рост от 6', и строку Owner: # УДАЛИТЬ
	grep -A 5 "L337" vehicles | grep -B 1 -A 4 "Honda" | grep -B 2 -A 3 "Blue" | grep -B 4 -A 1 "$min_height" | grep "Owner:" | while read -r owner_line; do

		# Убираем "Owner: " и лишний символ \r (бывает на Windows) # УДАЛИТЬ
		name="${owner_line#Owner: }"
		name="${name//$'\r'/}"

		# hint8: склеиваем 4 списка членств и считаем, сколько раз встречается имя # УДАЛИТЬ
		# Нужно ровно 4 — человек состоит во всех клубах из кошелька убийцы # УДАЛИТЬ
		count=$(cat memberships/AAA memberships/Delta_SkyMiles "memberships/$library" "memberships/$museum" | grep -c "$name")
		[ "$count" -eq 4 ] || continue

		# Проверяем, снят ли человек с подозрений в интервью # УДАЛИТЬ
		cleared=0
		person=$(grep "$name" people)
		if [ -n "$person" ]; then
			# Если он есть в people — снова head|tail по его улице, читаем его интервью # УДАЛИТЬ
			paddress="${person##*$'\t'}"
			pstreet="${paddress%, line *}"
			pline="${paddress##*, line }"
			pref=$(head -n "$pline" "streets/${pstreet// /_}" | tail -n 1)
			pnum=$(echo "$pref" | grep -o '[0-9]*')
			# Если в интервью написано "not considered" или "not a suspect" — не виновен # УДАЛИТЬ
			grep -i "not considered" "interviews/interview-$pnum" > /dev/null && cleared=1
			grep -i "not a suspect" "interviews/interview-$pnum" > /dev/null && cleared=1
		else
			# Если в people нет — ищем по фамилии во всех интервью (grep -r) # УДАЛИТЬ
			last="${name##* }"
			grep -r "$last" interviews/ | grep -i "not considered" > /dev/null && cleared=1
			grep -r "$last" interviews/ | grep -i "not a suspect" > /dev/null && cleared=1
		fi
		# Снятых с подозрений пропускаем # УДАЛИТЬ
		[ "$cleared" -eq 1 ] && continue

		# Остался один подозреваемый — печатаем его имя (это и есть ответ) # УДАЛИТЬ
		echo "$name"
	done
done
