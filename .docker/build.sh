#!/usr/bin/env bash

for bin in {conntroll,agent,hub,client}; do
  docker build -t conntroll/${bin} -f Dockerfile.${bin} .
done
