SRC         = main.go
ENV         = GOOS=js GOARCH=wasm
BIN_TARGET  = webcv
WASM_TARGET = ${BIN_TARGET}.wasm
TARGET_DIR  = bin

wasm: clear prepare
	@ ${ENV} go build -o ${TARGET_DIR}/${WASM_TARGET} ${SRC}
	@ echo 'target built: ${TARGET_DIR}/${BIN_TARGET}'

bin: clear prepare
	@ go build -o ${TARGET_DIR}/${BIN_TARGET} ${SRC}
	@ echo 'target built: ${TARGET_DIR}/${BIN_TARGET}'

clear:
	@ rm -f ${TARGET_DIR}/*

prepare:
	@ if [[ ! -d ${TARGET_DIR} ]]; then mkdir ${TARGET_DIR}; fi

