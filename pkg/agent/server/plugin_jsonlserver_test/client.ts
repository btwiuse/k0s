#!/usr/bin/env -S deno run -A --import-map=https://subqns.github.io/subshell/import_map.json

// ws://127.0.0.1:8000

import { ApiPromise, WsProvider } from "@polkadot/api";
import { encodeAddress, signatureVerify } from "@polkadot/util-crypto";

// websocat --binary  'wss://btwiuse-k0s-44wpv764f6qr-8000.githubpreview.dev/api/agent/abcd/jsonl'

export class Client {
  private ws;
  private isReady;
  private encoder = new TextEncoder();
  private decoder = new TextDecoder();
  private id = 0;

  constructor(base, id) {
    const WS_URL_ID = `${base}/api/agent/${id}/jsonl`;

    // new websocket connection
    this.ws = new WebSocket(WS_URL_ID);

    this.ws.binaryType = "arraybuffer";

    // register as agent
    this.ws.onopen = () => {
      this.isReady = true;
      // console.log("sending HEADER", HEADER);
      // ws.send(encoder.encode(HEADER));
      // this.ws.send(encoder.encode('{"method": "ls"}\n'));
      // console.log("sent HEADER");
    };

    this.ws.onclose = (e) => {
      console.log("closed", e);
    };

    this.ws.onerror = (e) => {
      console.log("errord", e);
    };
  }
  async web3Accounts() {
    let ws = this.ws;
    return new Promise((resolve, reject) => {
      ws.onmessage = (e) => {
        let cmd = this.decoder.decode(e.data);
        console.log("new recved", cmd);
        if (cmd.startsWith("{")) {
          resolve(JSON.parse(cmd).output);
        }
      };
      ws.send(this.encoder.encode('{"method": "ls"}\n'));
    });
  }
  async signRaw({ address, data }) {
    let ws = this.ws;
    const result = await new Promise((resolve, reject) => {
      ws.onmessage = (e) => {
        let cmd = this.decoder.decode(e.data);
        console.log("new recved", cmd);
        if (cmd.startsWith("{")) {
          resolve(JSON.parse(cmd));
        }
      };
      ws.send(
        this.encoder.encode(
          JSON.stringify({ method: "signRaw", args: [{address, data}] }) + `\n`,
        ),
      );
    });
    console.log(result)
    return result
  }
  async signPayload(payload) {
    let ws = this.ws;
    const result = await new Promise((resolve, reject) => {
      ws.onmessage = (e) => {
        let cmd = this.decoder.decode(e.data);
        console.log("new recved", cmd);
        if (cmd.startsWith("{")) {
          resolve(JSON.parse(cmd));
        }
      };
      ws.send(
        this.encoder.encode(
          JSON.stringify({ method: "signPayload", args: [payload] }) + `\n`,
        ),
      );
    });
    console.log(result)
    return result
  }
}

const BASE = `wss://btwiuse-k0s-44wpv764f6qr-8000.githubpreview.dev`;
const ID = `abcd`; // Math.random();

let client = new Client(BASE, ID);

while (!client.isReady) {
  await new Promise((resolve) => {
    setTimeout(resolve, 300);
  });
}

let accounts = (await client.web3Accounts()).map(x=>{x.address = encodeAddress(x.address, 0); return x});
console.log(accounts);
// let signature = await client.sign(accounts[0].address, "Hello");
// console.log(signature);

const provider = new WsProvider("wss://rpc.polkadot.io");
const api = await ApiPromise.create({ provider });
// api.setSigner({signRaw: (x)=>{console.log(x); return {signature: 'sig'}}})
api.setSigner(client);
for (let i of [1,2,3]) {
	let addr = accounts[0].address;
	let msg = i + "Hello" + new Date()
  const sig = await api.sign(addr, { data: msg });
  console.log(msg, sig, addr)
  const verify = signatureVerify(msg, sig, addr)
  console.log(verify)
  console.log(encodeAddress(verify.publicKey, 0))
}
await api.tx.system.remark('asdf').signAndSend(accounts[0].address)
