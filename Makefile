SRC         = cmd/main.go
ENV         = GOOS=js GOARCH=wasm
BIN_TARGET  = webcv
WASM_TARGET = ${BIN_TARGET}.wasm

wasm: clear
	${ENV} go build -o bin/${WASM_TARGET} ${SRC}

bin: clear
	go build -o bin/${BIN_TARGET} ${SRC}

clear:
	rm -f bin/*

