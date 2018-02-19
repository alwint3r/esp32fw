.PHONY: all windows macos linux single

all: windows linux macos

windows:
	GOOS=windows GOARCH=amd64 go build -o esp32fw-windows-amd64.exe
	GOOS=windows GOARCH=386 go build -o esp32fw-windows-i386.exe

linux:
	GOOS=linux GOARCH=amd64 go build -o esp32fw-linux-amd64
	GOOS=linux GOARCH=386 go build -o esp32fw-linux-i386

macos:
	GOOS=darwin GOARCH=amd64 go build -o esp32fw-darwin-amd64
	GOOS=darwin GOARCH=386 go build -o esp32fw-darwin-i386

single:
	go build -o esp32fw

clean:
	rm -rf esp32fw*