build: 
	go env -w GOOS=js GOARCH=wasm

	cp ./src/vendor/v_*.go ./src/frontend/

	go build -o ./website/res/main.wasm ./src/frontend/

clean:
	rm ./src/frontend/v_*.go

all: build clean