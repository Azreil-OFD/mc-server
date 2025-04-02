#!/bin/bash

SCRIPTS_DIR="./scripts"
COMMAND="$1"

# Проверяем, передана ли команда
if [[ -z "$COMMAND" ]]; then
  echo "❌ Не указана команда."
  echo
  show_help
  exit 1
fi

SCRIPT_PATH="$SCRIPTS_DIR/$COMMAND.sh"
HELP_PATH="$SCRIPTS_DIR/$COMMAND.txt"

function show_help {
  echo "🛠 Доступные команды:"
  for script in "$SCRIPTS_DIR"/*.sh; do
    cmd_name=$(basename "$script" .sh)
    desc_file="$SCRIPTS_DIR/$cmd_name.txt"
    if [[ -f "$desc_file" ]]; then
      description=$(cat "$desc_file")
    else
      description="(описание отсутствует)"
    fi
    printf "  %-10s — %s\n" "$cmd_name" "$description"
  done
}

# Выполняем команду, если скрипт существует
if [[ -f "$SCRIPT_PATH" ]]; then
  bash "$SCRIPT_PATH"
else
  echo "❌ Команда '$COMMAND' не найдена."
  echo
  show_help
  exit 1
fi
