import * as time from './time.ts'

async function agent(){
	while (true) {
		console.log('running agent in background')
		await time.sleep(1000)
		console.log('done')
	}
}

agent()
