## WASI
### Build
```bash
GOOS=wasip1 GOARCH=wasm go build main.go 
```

### Run
```bash
wasmtime main
```

### Convert to WAT
using https://github.com/WebAssembly/wabt
```
bin/wasm2wat main > gomain.wat
```
