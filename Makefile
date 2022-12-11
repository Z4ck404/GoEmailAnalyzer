build: clean
	go build -o bin/mailp

clean:
	rm -rf bin/mailp*

fmt: 
	$(GOIMPORTS) -w $(shell pwd)
	go fmt ./...

GOIMPORTS = go run golang.org/x/tools/cmd/goimports@v0.1.12