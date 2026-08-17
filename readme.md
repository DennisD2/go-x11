## go-x11 - Go Wrapper for X11 libraries

Work in progress.

## Internals

Uses CGO to access C code.

Analysis of created CGO wrapper code:
```shell
go tool cgo main.go
```
This creates a directory _obj. And adds all generated glue code from CGO. 
Very interesting is e.g. _obj/_cgo_gotypes.go for all in-betrween types/structs/casts.

## More information
* Good spicker, german - https://opensource.archium.org/index.php/Der_Golang-Spicker
* unsafe package - https://leapcell.medium.com/gos-unsafe-unlocking-performance-hacks-with-a-risk-16d1d8dd9afb