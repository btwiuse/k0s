all:
	go build -v -trimpath -o conntroll .
	strip conntroll
	upx conntroll
