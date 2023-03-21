
# Run benchmark
```
go test -bench . -benchmem
```
Save file
```
go test -bench . -benchmem > crypto
```

# Compare
```
benchstat  math crypto 
```

```
goos: linux
goarch: amd64
pkg: github.com/blck-snwmn/playground-go/benchcmp
cpu: 11th Gen Intel(R) Core(TM) i7-11700 @ 2.50GHz
           │     math     │                 crypto                 │
           │    sec/op    │    sec/op     vs base                  │
Sample-16    126.1n ± ∞ ¹   890.3n ± ∞ ¹         ~ (p=1.000 n=1) ²
Sample2-16   1.150µ ± ∞ ¹   4.097µ ± ∞ ¹         ~ (p=1.000 n=1) ²
geomean      380.8n         1.910µ        +401.53%
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05

           │     math      │              crypto              │
           │     B/op      │     B/op       vs base           │
Sample-16        0.0 ± ∞ ¹     112.0 ± ∞ ¹  ~ (p=1.000 n=1) ²
Sample2-16   0.000Ki ± ∞ ¹   1.000Ki ± ∞ ¹  ~ (p=1.000 n=1) ²
geomean                  ³     338.7        ?
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05
³ summaries must be >0 to compute geomean

           │    math     │             crypto             │
           │  allocs/op  │  allocs/op   vs base           │
Sample-16    0.000 ± ∞ ¹   1.000 ± ∞ ¹  ~ (p=1.000 n=1) ²
Sample2-16   0.000 ± ∞ ¹   1.000 ± ∞ ¹  ~ (p=1.000 n=1) ²
geomean                ³   1.000        ?
¹ need >= 6 samples for confidence interval at level 0.95
² need >= 4 samples to detect a difference at alpha level 0.05
³ summaries must be >0 to compute geomean
```