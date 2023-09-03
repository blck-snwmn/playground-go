## Use slice (map[int][128]byte)

```bash
$ go run mapslice/bench.go
Alloc = 0 MiB
        TotalAlloc = 0 MiB
Alloc = 293 MiB
        TotalAlloc = 736 MiB
Alloc = 293 MiB
        TotalAlloc = 736 MiB
```

## Use pointer of slice (map[int]*[128]byte)
```bash
$ go run mapslice/bench.go -p=true
Alloc = 0 MiB
        TotalAlloc = 0 MiB
Alloc = 160 MiB
        TotalAlloc = 324 MiB
Alloc = 38 MiB
        TotalAlloc = 324 MiB
```
