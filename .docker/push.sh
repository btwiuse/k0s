#!/usr/bin/env bash
for bin in {conntroll,agent,hub,client}; do
  docker --config ~/.docker/conntroll-config.json push conntroll/${bin}
done
