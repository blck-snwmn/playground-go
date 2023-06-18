use Go1.21rc1
## WASI
### Build
```bash
GOOS=wasip1 GOARCH=wasm go build -o main.wasm main.go 
```

### Run
```bash
wasmtime main.wasm
```

### Convert to WAT
using https://github.com/WebAssembly/wabt
```
bin/wasm2wat main > gomain.wat
```
