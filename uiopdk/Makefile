pre-build:
	go env -u GOOS GOARCH

build-initiazer: 
	go build -o ./binary/uiopdk-init.exe ./src/uiopdk-init/

clean:

all: pre-build build-initiazer