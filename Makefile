pre-build:
	go env -w GOOS=js GOARCH=wasm

build-frontend: 
	go build -o ./website/res/main.wasm ./src/frontend/

build-emulater:
	go build -o ./website/res/emulater.wasm ./src/backend/emulater

clean:

all: pre-build build-frontend build-emulater