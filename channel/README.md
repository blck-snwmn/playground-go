## Use nil channel
```bash
$ go run nilch/main.go 
chl 1
chl 2
1
2
chr 3
4
chr 4
chr 5
5
6
chr 6
        chr closed 6
chl 7
chl 8
        chl closed 8
3
```

## USe non nil channel
when other channel is closed, loop is continue.

```bash
$ go run nilch/main.go 
chl 1
chl 2
1
2
chl 3
3
chl 4
        chl closed 4
chr 5
chr 6
4
5
chr 7

~~

chl 5922
chr 5923
        chr closed 5923
chr 5924
        chr closed 5924
chl 5925
        chl closed 5925
3
```
