
all:
	go generate x11/defs.go
	sed "s/\*_Ctype_struct__XDisplay/unsafe.Pointer/g" x11/x11_types.go > x11/x11_types.tmp
	sed "s/package x11/package x11\n\nimport \"unsafe\"/g" x11/x11_types.tmp > x11/x11_types.go
	rm -f x11/x11_types.tmp
	go build ./...
