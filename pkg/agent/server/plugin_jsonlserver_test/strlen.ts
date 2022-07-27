#!/usr/bin/env -S deno run -A

// ws://127.0.0.1:8000

const META = {
  "os": "web",
  "pwd": window.location?.pathname || "/dev/null",
  "arch": "js",
  "distro": "js",
  "username": "webuser",
  "hostname": window.location?.hostname || "unknown",
};

const VERSION = {
  "GitCommit": "e040989acfeeebcbab79b1825222df08138c749b",
  "GitState": "dirty",
  "GitBranch": "master",
  "GitSummary": "release-75-20220614-5-ge040989ac-dirty",
  "BuildDate": "2022-06-20T05:31:29Z",
  "Version": "v0.0.15",
  "GoVersion": "go1.18.3",
};

const encoder = new TextEncoder();
const decoder = new TextDecoder();

function handle(conn, cmd) {
  console.log("handling", cmd);
  if (cmd.startsWith("{")) {
    try {
      let json = JSON.parse(cmd);
      console.log("got json", json);
      if (json.method) {
        handleMethod(conn, json.method, json.args);
      }
    } catch (e) {
    }
  }
}

function handleMethod(conn, method, args) {
  if (method == "reverse") {
    let payload = JSON.stringify({
      output: args[0].split().reverse().join(),
    }) + `\n`;
    if (args && args.length > 0) {
      console.log("sending", payload);
      conn.send(encoder.encode(payload));
    }
  }
}

class Agent {
  constructor(base, id) {
    const WS_URL = `${base}/api/rpc`;

    const WS_URL_ID = `${base}/api/jsonl?id=${id}`;

    const HEADER = JSON.stringify({
      "id": id,
      "name": `name_of_${id}`,
      "tags": [],
      "meta": META,
      "version": VERSION,
    }) + "\n";

    // new websocket connection
    let ws = new WebSocket(WS_URL);

    ws.binaryType = "arraybuffer";

    // register as agent
    ws.onopen = () => {
      console.log("sending HEADER", HEADER);
      ws.send(encoder.encode(HEADER));
      console.log("sent HEADER");
    };

    ws.onclose = (e) => {
      console.log("closed", e);
    };

    ws.onerror = (e) => {
      console.log("errord", e);
    };

    ws.onmessage = (e) => {
      let cmd = decoder.decode(e.data);
      console.log("recved", cmd);
      if (cmd == "JSONL\n") {
        let conn = new WebSocket(WS_URL_ID);
        conn.binaryType = "text";
        conn.onmessage = async (e) => {
          let data = await e.data.text();
          console.log("recv", data);
          handle(conn, data);
          conn.send(encoder.encode(`${data.length}\n`));
          // let recv = decoder.decode(e.data)
          // console.log("recv", recv);
        };
      }
    };
  }
}

const BASE = `wss://btwiuse-k0s-44wpv764f6qr-8000.githubpreview.dev`;
const ID = `abcd`; // Math.random();

let agent = new Agent(BASE, ID);
