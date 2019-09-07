package main

import (
	"fmt"
	"syscall/js"
)

func PassUint8ArrayToGo(this js.Value, args []js.Value) interface{} {
	received := make([]byte, args[0].Get("length").Int())
	_ = js.CopyBytesToGo(received, args[0])
	fmt.Println(received)
	return nil
}

func SetUint8ArrayInGo(this js.Value, args []js.Value) interface{} {
	_ = js.CopyBytesToJS(args[0], []byte{0, 9, 21, 32})
	return nil
}

func registerCallbacks() {
	js.Global().Set("PassUint8ArrayToGo", js.FuncOf(PassUint8ArrayToGo))
	js.Global().Set("SetUint8ArrayInGo", js.FuncOf(SetUint8ArrayInGo))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
