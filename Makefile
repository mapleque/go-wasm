.PHONY: run
run: bin/serve main.wasm
	./bin/serve -p 8080

.PHONY: build
build: main.wasm bin/serve

main.wasm: main.go
	GOOS=js GOARCH=wasm go build -o main.wasm

bin/serve: serve/main.go
	-mkdir bin
	go build -o bin/serve serve/main.go

clean:
	-rm -rf bin
	-rm -f main.wasm
