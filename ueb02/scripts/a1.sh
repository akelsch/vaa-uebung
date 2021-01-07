#!/usr/bin/env bash

if [[ $# -ne 1 ]] || ! [[ $1 =~ ^[0-9]+$ ]]; then
  echo "Usage: a1.sh <number of nodes>"
  exit 1
fi

NODE_COUNT=$1
NEIGHBOR_COUNT=$((NODE_COUNT - 1))

if ((NEIGHBOR_COUNT == 0)); then
  NEIGHBOR_COUNT=1
fi

CMDLINE="tail -q "
for i in $(seq 1 "$NODE_COUNT"); do
  # m -> rand(1..m) = t
  # s -> No. of starting philosophers
  # p -> No. of rand neighbor philosophers per philosopher
  # amax -> Max connections per philosopher
  CMDLINE+=">(philosopher.exe -f ./configs/config.csv -gv ./configs/topology.gv -id ${i} -m 10 -s 2 -p 2 -amax 2 2>&1) "
done

eval "$CMDLINE"
