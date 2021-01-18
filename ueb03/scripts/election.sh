#!/usr/bin/env bash

if [[ $# -ne 1 ]]; then
  echo "Usage: election.sh <config file>"
  exit 1
fi

CONFIG_FILE=$1
CONFIG_LINES=$(wc -l <"$CONFIG_FILE")

RAND_NUMBER=$(shuf -i 1-"$CONFIG_LINES")
SHUFFLED_CONFIG=$(shuf -n "$RAND_NUMBER" "$CONFIG_FILE")

ELECTION_PORTS=$(echo "$SHUFFLED_CONFIG" | cut -d "," -f3 | tr -d '\r') # tr: cleanup Windows carriage return

# Change working directory for protoc
cd "./api/pb" || exit

while IFS= read -r port; do
  echo "$port"
  eval "printf 'control_message:{command:START_ELECTION}' | protoc.exe --encode=ueb03.Message Message.proto | nc localhost $port"
done <<<"$ELECTION_PORTS"
