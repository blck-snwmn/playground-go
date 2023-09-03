## Use slice (map[int][128]byte)

```bash
$ go run mapslice/bench.go -m=s
Alloc = 0 MiB
        TotalAlloc = 0 MiB
Alloc = 293 MiB
        TotalAlloc = 735 MiB
Alloc = 293 MiB
        TotalAlloc = 735 MiB
```

## Use pointer of slice (map[int]*[128]byte)
```bash
$ go run mapslice/bench.go -m=p
Alloc = 0 MiB
        TotalAlloc = 0 MiB
Alloc = 160 MiB
        TotalAlloc = 202 MiB
Alloc = 38 MiB
        TotalAlloc = 202 MiB
```

## Use pointer of slice (map[int][129]byte)
```bash
$ go run mapslice/bench.go -m=big
Alloc = 0 MiB
        TotalAlloc = 0 MiB
Alloc = 175 MiB
        TotalAlloc = 354 MiB
Alloc = 38 MiB
        TotalAlloc = 354 MiB
```

## Use pointer of slice (map[int]*[129]byte)
```bash
$ go run mapslice/bench.go -m=pbig
Alloc = 0 MiB
        TotalAlloc = 0 MiB
Alloc = 175 MiB
        TotalAlloc = 217 MiB
Alloc = 38 MiB
        TotalAlloc = 217 MiB
```