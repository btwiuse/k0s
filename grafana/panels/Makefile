.PHONY: all nodes

all: nodes
	@ ./make

nodes:
	@ echo - __hub__ > nodelist.yaml
	@ curl -sL https://k0s.herokuapp.com/api/agents/list | jq .[].name | jq -s . | yj -jy | sort -h >> nodelist.yaml
