#!/usr/bin/env bash
curl -o example.dashboard.py https://raw.githubusercontent.com/weaveworks/grafanalib/master/grafanalib/tests/examples/example.dashboard.py
generate-dashboard -o frontend.json example.dashboard.py
