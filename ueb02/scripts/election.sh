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
  eval "printf 'control_message:{command:START_ELECTION}' | protoc.exe --encode=ueb02.Message Message.proto | nc localhost $port"
done <<<"$ELECTION_PORTS"

# Chapter 4 example
#eval "printf 'control_message:{command:START_ELECTION}' | protoc.exe --encode=ueb02.Message Message.proto | nc localhost 5003"
#eval "printf 'control_message:{command:START_ELECTION}' | protoc.exe --encode=ueb02.Message Message.proto | nc localhost 5004"
#printf 'control_message:{command:EXIT}' | protoc.exe --encode=ueb02.Message Message.proto | nc localhost 5000 &&
#printf 'control_message:{command:EXIT}' | protoc.exe --encode=ueb02.Message Message.proto | nc localhost 5001 &&
#printf 'control_message:{command:EXIT}' | protoc.exe --encode=ueb02.Message Message.proto | nc localhost 5002 &&
#printf 'control_message:{command:EXIT}' | protoc.exe --encode=ueb02.Message Message.proto | nc localhost 5003 &&
#printf 'control_message:{command:EXIT}' | protoc.exe --encode=ueb02.Message Message.proto | nc localhost 5004 &&
#printf 'control_message:{command:EXIT}' | protoc.exe --encode=ueb02.Message Message.proto | nc localhost 5005 &&
#printf 'control_message:{command:EXIT}' | protoc.exe --encode=ueb02.Message Message.proto | nc localhost 5006
#[p-01] 21:49:10.612381 Received control message: EXIT
#[p-01] 21:49:10.612381 pred 2
#[p-02] 21:49:10.645381 Received control message: EXIT
#[p-02] 21:49:10.645381 pred 3
#[p-03] 21:49:10.679382 Received control message: EXIT
#[p-03] 21:49:10.679382 pred 5
#[p-04] 21:49:10.712381 Received control message: EXIT
#[p-04] 21:49:10.712381 pred 6
#[p-05] 21:49:10.746381 Received control message: EXIT
#[p-05] 21:49:10.746381 pred 6
#[p-06] 21:49:10.780381 Received control message: EXIT
#[p-06] 21:49:10.780381 pred 5
#[p-07] 21:49:10.813381 Received control message: EXIT
#[p-07] 21:49:10.813381 pred 6
