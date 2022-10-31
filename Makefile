build:
	go env -w GOOS=js GOARCH=wasm

	go build -o ./website/res/main.wasm ./src/wasm/ 

all:
	build deploy