# Run
Use Go1.21rc2

## No GOEXPERIMENT
```bash
go run main.go
```

result:
```
num=10
num=10
num=10
num=10
num=10
num=10
num=10
num=10
num=10
num=10
```

## Use GOEXPERIMENT
```bash
GOEXPERIMENT=loopvar go run main.go 
```

result:
```
num=9
num=4
num=5
num=6
num=7
num=8
num=0
num=1
num=2
num=3
```
