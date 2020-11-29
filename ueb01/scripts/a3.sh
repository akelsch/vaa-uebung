#!/bin/bash

if [[ $# -ne 1 ]] || ! [[ $1 =~ ^[0-9]+$ ]]; then
  echo "Usage: a3.sh <number of files>"
  exit 1
fi

NUM_FILES=$1

echo "n,m"
for i in $(seq -w 1 "$NUM_FILES"); do
  NODE_COUNT=$(shuf -i 4-10 -n 1)

  EDGE_MIN=$((NODE_COUNT + 1))
  EDGE_MAX=$(((NODE_COUNT * (NODE_COUNT - 1)) / 2))
  EDGE_LIMIT=$((EDGE_MIN * 2))

  if ((EDGE_MAX > EDGE_LIMIT)); then
    EDGE_MAX=$EDGE_LIMIT
  fi

  EDGE_COUNT=$(shuf -i ${EDGE_MIN}-${EDGE_MAX} -n 1)

  echo "${NODE_COUNT},${EDGE_COUNT}"

  eval "graphgen.exe -n ${NODE_COUNT} -m ${EDGE_COUNT} > ./configs/gen/${i}.gv"
done
