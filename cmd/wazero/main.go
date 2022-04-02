package main

import (
	_ "embed"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/wasi"
	"log"
)

// fibWasm was compiled from TinyGo testdata/fibonacci.go
//go:embed testdata/fibonacci.wasm
var fibWasm []byte

func main() {
	runtime := wazero.NewRuntime()

	// Note: fibonacci.go doesn't directly use WASI, but TinyGo needs to be initialized as a WASI Command.
	wm, err := wasi.InstantiateSnapshotPreview1(runtime)
	if err != nil {
		log.Fatalln(err)
	}
	defer wm.Close()

	module, err := runtime.InstantiateModuleFromCode(fibWasm)
	if err != nil {
		log.Fatalln(err)
	}
	defer module.Close()

	fibonacci := module.ExportedFunction("fibonacci")

	for _, c := range []struct {
		input, expected uint64 // i32_i32 sig, but wasm.ExportedFunction params and results are uint64
	}{
		{input: 20, expected: 6765},
		{input: 10, expected: 55},
		{input: 5, expected: 5},
	} {
		results, err := fibonacci.Call(nil, c.input)
		if err != nil {
			log.Fatalln(err)
		}
		if c.expected != results[0] {
			log.Fatalf("expected %d, but have %d", c.expected, results[0])
		}
	}
}
