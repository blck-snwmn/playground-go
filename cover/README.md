# Test Cover
## Test w/ coverage
```
go test ./... -coverprofile cover.out
```

## Use cover w/ total
```
go tool cover -func cover.out
```

# Cover
## Prepare
- Do `mkdir coverdir`
- Install go1.20rc2

## Build
```
go1.20rc2 build -cover -o cover
```

## Use covdata
```
go1.20rc2 tool covdata percent -i coverdir
```