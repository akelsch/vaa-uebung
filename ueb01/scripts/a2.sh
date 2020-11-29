#!/usr/bin/env bash

if [[ $# -ne 1 ]] || ! [[ $1 =~ ^[0-9]+$ ]]; then
  echo "Usage: a2.sh <number of nodes>"
  exit 1
fi

NODE_COUNT=$1

CMDLINE="tail -q "
for i in $(seq 1 "$NODE_COUNT"); do
  CMDLINE+=">(mynode.exe -f ./configs/config.csv -gv ./configs/topology.gv -id ${i} 2>&1) "
done

eval "$CMDLINE"
