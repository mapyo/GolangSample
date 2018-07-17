.PHONY: clean build start

clean:
	rm -rf ./hello-world/hello-world
	
build:
	GOOS=linux GOARCH=amd64 vgo build -o hello-world/bin/hello-world ./hello-world

start:
	sam local start-api
