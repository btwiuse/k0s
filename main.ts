#!/usr/bin/env -S deno run -A

import './pkg/plugin.ts'
import * as time from './pkg/time.ts'

while (true) {
	await time.sleep(1000);
}
