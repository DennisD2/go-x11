## go-x11 - Go Wrapper for X11 libraries

Uses CGO to access C code.

Analysis of created CGO wrapper code:
```shell
go tool cgo main.go
```
This creates a directory _obj. And adds all generated glue code from CGO. 
Very interesting is e.g. _obj/_cgo_gotypes.go for all in-betrween types/structs/casts.

## More information
* Good spicker, german - https://opensource.archium.org/index.php/Der_Golang-Spicker