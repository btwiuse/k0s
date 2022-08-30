package main

import (
	"context"
	_ "embed"
	"log"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/wasi_snapshot_preview1"
)

// fibWasm was compiled from TinyGo testdata/fibonacci.go
// go:embed testdata/fibonacci.wasm
var fibWasm []byte

func main() {
	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfig().
		// WebAssembly 2.0 allows use of any version of TinyGo, including 0.24+.
		WithWasmCore2())
	defer r.Close(ctx) // This closes everything this Runtime created.

	// Instantiate WASI, which implements system I/O such as console output.
	if _, err = wasi_snapshot_preview1.Instantiate(ctx, r); err != nil {
		log.Panicln(err)
	}

	module, err := r.InstantiateModuleFromBinary(ctx, fibWasm)
	if err != nil {
		log.Fatalln(err)
	}

	fibonacci := module.ExportedFunction("fibonacci")

	for _, c := range []struct {
		input, expected uint64 // i32_i32 sig, but wasm.ExportedFunction params and results are uint64
	}{
		{input: 20, expected: 6765},
		{input: 10, expected: 55},
		{input: 5, expected: 5},
	} {
		results, err := fibonacci.Call(ctx, c.input)
		if err != nil {
			log.Fatalln(err)
		}
		if c.expected != results[0] {
			log.Fatalf("expected %d, but have %d", c.expected, results[0])
		}
	}
}
