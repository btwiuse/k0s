import { ApiPromise } from 'https://deno.land/x/polkadot@0.0.5/api/mod.ts';

export function log(...x: any[]) {
  console.log(new Date(), ...x);
}

const listener = Deno.listen({
  port: 8000,
  transport: 'tcp'
})

for(;;) {
  const conn = await listener.accept();
  handleConn(conn);
}

async function handleConn(conn: Deno.Conn) {
  const decoder = new TextDecoder();
  const encoder = new TextEncoder();
  for(;;) {
    const buf = new Uint8Array(1024);
    const n = await conn.read(buf);
    if (!n) {
      log("client closed");
      break;
    }
    const data = decoder.decode(buf.subarray(0, n));
    log(data);
    try {
      await conn.write(encoder.encode(data));
    } catch(e) {
      log(e);
      break
    }
  }
}
