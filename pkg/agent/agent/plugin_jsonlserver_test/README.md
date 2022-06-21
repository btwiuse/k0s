dump example agent header

```
k0s agent -verbose
cat agent.json | jq -c > agent.jsonl
```

mock agent, register only. cannot respond to requests

```
cat agent.jsonl - | websocat --text ws://127.0.0.1:8000/api/rpc
```

deno strlen agent, can respond to requests

```
./strlen.ts
```

browser strlen agent, same functionality

```
browser .
```

mock client, now you can send input to strlen server and get line length as
response!

```
websocat --binary  'ws://127.0.0.1:8000/api/agent/abcd/jsonl'
```

example output:

```
$ websocat --binary  'ws://127.0.0.1:8000/api/agent/abcd/jsonl'
foo
4
hello world
12
```
