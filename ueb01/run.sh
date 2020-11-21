#!/bin/bash

if [[ $# -ne 1 ]] || ! [[ $1 =~ ^[0-9]+$ ]]; then
  echo "Usage: run.sh <number of nodes>"
  exit 2
fi

NODE_COUNT=$1
NEIGHBOR_COUNT=$((NODE_COUNT - 1))

CMDLINE="tail -q "
for i in $(seq -w 1 "$NODE_COUNT"); do
  CMDLINE+=">(ueb01 -f ./testdata/config.csv -id ${i} -n ${NEIGHBOR_COUNT} 2>&1 | sed -e 's/^/[${i}]    /;') "
done

eval "${CMDLINE}"
