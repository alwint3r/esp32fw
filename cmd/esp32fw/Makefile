
OUTPUT_DIR := bin

.PHONY: all windows macos linux single prepare

all: prepare windows linux macos

windows:
	GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/esp32fw-windows-amd64.exe
	GOOS=windows GOARCH=386 go build -o $(OUTPUT_DIR)/esp32fw-windows-i386.exe

linux:
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/esp32fw-linux-amd64
	GOOS=linux GOARCH=386 go build -o $(OUTPUT_DIR)/esp32fw-linux-i386

macos:
	GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)/esp32fw-darwin-amd64
	GOOS=darwin GOARCH=386 go build -o $(OUTPUT_DIR)/esp32fw-darwin-i386

single:
	go build -o $(OUTPUT_DIR)/esp32fw

clean:
	rm -rf $(OUTPUT_DIR)

prepare:
	mkdir -p $(OUTPUT_DIR)