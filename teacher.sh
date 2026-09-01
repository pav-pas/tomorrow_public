#!/usr/bin/env bash

# step 1: номер ключевого интервью (Annabel Church, Buckingham Place, line 179)
# tr -d '\r' убирает Windows-перенос строки, иначе путь к файлу ломается
INTERVIEWNUMBER=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2 | tr -d '\r')

# step 2: печатаем переменную
echo $INTERVIEWNUMBER

# step 3: печатаем содержимое интервью
cat interviews/interview-"$INTERVIEWNUMBER"

# step 4: печатаем MAIN_SUSPECT (задаётся снаружи при запуске теста)
echo $MAIN_SUSPECT