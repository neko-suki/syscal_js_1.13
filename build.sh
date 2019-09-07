#!/bin/bash
GOOS=js GOARCH=wasm go build -o main.wasm main.go
#curl -sO https://raw.githubusercontent.com/golang/go/go1.13/misc/wasm/wasm_exec.js
