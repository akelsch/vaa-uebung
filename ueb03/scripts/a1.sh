#!/usr/bin/env bash

if [[ $# -ne 1 ]] || ! [[ $1 =~ ^[0-9]+$ ]]; then
  echo "Usage: a1.sh <number of nodes>"
  exit 1
fi

NODE_COUNT=$1

GRAPH_FILE="./configs/tmp.gv"
CONFIG_FILE="./configs/tmp.csv"

function graphgen() {
  EDGE_MIN=$((NODE_COUNT + 1))
  EDGE_MAX=$(((NODE_COUNT * (NODE_COUNT - 1)) / 2))
  EDGE_LIMIT=$((EDGE_MIN * 2))

  if ((EDGE_MAX > EDGE_LIMIT)); then
    EDGE_MAX=$EDGE_LIMIT
  fi

  EDGE_COUNT=$(shuf -i ${EDGE_MIN}-${EDGE_MAX} -n 1)

  echo "n,m"
  echo "${NODE_COUNT},${EDGE_COUNT}"

  eval "graphgen.exe -n ${NODE_COUNT} -m ${EDGE_COUNT} > ${GRAPH_FILE}"
}

function confgen() {
  true >$CONFIG_FILE
  for i in $(seq 1 "$NODE_COUNT"); do
    PORT=$((5000 + i - 1))
    echo "${i},localhost,${PORT}" >>$CONFIG_FILE
  done
}

graphgen && confgen

CMDLINE="tail -q "
for i in $(seq 1 "$NODE_COUNT"); do
  # CMDLINE+=">(baccount.exe -f ./configs/config.csv -gv ./configs/topology.gv -id ${i} 2>&1) "
  CMDLINE+=">(baccount.exe -f ${CONFIG_FILE} -gv ${GRAPH_FILE} -id ${i} 2>&1) "
done

eval "$CMDLINE"
