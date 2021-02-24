import threading, time

def agent():
    while True:
        print("running agent in thread")
        time.sleep(1)
        print("done")

threading.Thread( target = agent ).start()
