#!/usr/bin/env bash

if [[ $# -ne 1 ]]; then
  echo "Usage: a4.sh <graphviz file>"
  exit 1
fi

SRC_DIR="./configs/gen/"

CONFIG_PATH="${SRC_DIR}config.csv"
GV_PATH="${SRC_DIR}${1}.gv"

NODE_COUNT=$(sort -n "$GV_PATH" | tail -n 1 | sed "s/[^0-9]*//g")

CMDLINE="tail -q "
for i in $(seq 1 "$NODE_COUNT"); do
  CMDLINE+=">(mynode.exe -f ${CONFIG_PATH} -gv ${GV_PATH} -id ${i} 2>&1) "
done

eval "$CMDLINE"
