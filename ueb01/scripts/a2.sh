#!/bin/bash

if [[ $# -ne 1 ]] || ! [[ $1 =~ ^[0-9]+$ ]]; then
  echo "Usage: a2.sh <number of nodes>"
  exit 2
fi

NODE_COUNT=$1

CMDLINE="tail -q "
for i in $(seq -w 1 "$NODE_COUNT"); do
  CMDLINE+=">(a2.exe -f ./configs/config.csv -gv ./configs/topology.gv -id ${i} 2>&1 | sed -e 's/^/[${i}]    /;') "
done

eval "${CMDLINE}"
