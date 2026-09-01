#!/usr/bin/env bash

# Тестовый запуск teacher.sh (как на проверке: изнутри mystery + MAIN_SUSPECT)
cd "$(dirname "$0")/mystery" || exit 1

export MAIN_SUSPECT="Harvey Dent"
bash ../teacher.sh