//#!/usr/bin/env -S deno run -A

// ws://127.0.0.1:8000

const BASE = `wss://btwiuse-k0s-44wpv764f6qr-8000.githubpreview.dev`;

const WS_URL = `${BASE}/api/rpc`;

const ID = `abcd`; // Math.random();

const WS_URL_ID = `${BASE}/api/jsonl?id=${ID}`;

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

const HEADER = JSON.stringify({
  "id": `${ID}`,
  "name": `name_of_${ID}`,
  "tags": [],
  "meta": META,
  "version": VERSION,
}) + "\n";

// new websocket connection
let ws = new WebSocket(WS_URL);

ws.binaryType = "arraybuffer";

const encoder = new TextEncoder();
const decoder = new TextDecoder();

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
    console.log(WS_URL_ID);
    let conn = new WebSocket(WS_URL_ID);
    conn.binaryType = "blob";
    conn.onmessage = async (e) => {
      let data = await e.data.text();
      console.log("recv", data);
      conn.send(encoder.encode(`${data.length}\n`));
      // let recv = decoder.decode(e.data)
      // console.log("recv", recv);
    };
  }
};
