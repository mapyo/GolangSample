.PHONY: clean build start test

clean:
	rm -rf ./hello-world/hello-world
	
build:
	GOOS=linux GOARCH=amd64 vgo build -o hello-world/bin/hello-world ./hello-world/function/hello

test:
	vgo test ./hello-world/function/hello

start:
	sam local start-api
