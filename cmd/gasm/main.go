package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/mathetake/gasm/wasi"
	"github.com/mathetake/gasm/wasm"
)

func main() {
	buf, err := ioutil.ReadFile("wasm/fibonacci.wasm")
	if err != nil {
		log.Fatalln(err)
	}

	mod, err := wasm.DecodeModule(bytes.NewBuffer(buf))
	if err != nil {
		log.Fatalln(err)
	}

	vm, err := wasm.NewVM(mod, wasi.Modules)
	if err != nil {
		log.Fatalln(err)
	}

	for _, c := range []struct {
		in, exp int32
	}{
		{in: 20, exp: 6765},
		{in: 10, exp: 55},
		{in: 5, exp: 5},
	} {
		vm.ExecExportedFunction("fibonacci", uint64(c.in))
		/*
			ret, retTypes, err := vm.ExecExportedFunction("fibonacci", uint64(c.in))
			require.NoError(t, err)
			require.Len(t, ret, len(retTypes))
			require.Equal(t, wasm.ValueTypeI32, retTypes[0])
			require.Equal(t, c.exp, int32(ret[0]))
		*/
	}
}
