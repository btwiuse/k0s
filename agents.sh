#!/usr/bin/env bash
curl -s https://hub.k0s.io/api/agents/list \
  | jq '.[]|{"targets": ["hub.k0s.io"], "labels": {"job": "\(.name)", "__metrics_path__": "/api/agent/\(.id)/metrics"}}' \
  | jq -s . > /tmp/agents.json;
cp -v /tmp/agents.json .
cat agents.json
