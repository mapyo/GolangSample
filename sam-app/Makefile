.PHONY: clean build

clean:
	rm -rf ./hello-world/hello-world
	
build:
	GOOS=linux GOARCH=amd64 vgo build -o hello-world/hello-world ./hello-world
