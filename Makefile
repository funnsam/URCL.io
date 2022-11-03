pre-build:
	go env -w GOOS=js GOARCH=wasm

build-frontend: 
	cp ./src/vendor/v_*.go ./src/frontend
	go build -o ./website/res/main.wasm ./src/frontend/

build-emulater:
	cp ./src/vendor/v_*.go ./src/backend/emulater
	go build -o ./website/res/emulater.wasm ./src/backend/emulater

clean:
	rm ./src/frontend/v_*.go
	rm ./src/backend/*/v_*.go

all: pre-build build-frontend build-emulater