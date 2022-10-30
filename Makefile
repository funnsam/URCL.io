build:
	go env -w GOOS=js GOARCH=wasm

	go build -o ./docs/res/main.wasm ./src/wasm/ 

all:
	build deploy