SRC         = cmd/main.go
ENV         = GOOS=js GOARCH=wasm
BIN_TARGET  = webcv
WASM_TARGET = ${BIN_TARGET}.wasm

wasm:
	${ENV} go build -o bin/${WASM_TARGET} ${SRC}

bin: 
	go build -o bin/${BIN_TARGET} ${SRC}

