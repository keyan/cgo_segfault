package main

// rdynamic - includes all symbols to the symbol table when creating dynamically
// linked executables. https://gcc.gnu.org/legacy-ml/gcc-help/1999-12n/msg00076.html

// #cgo LDFLAGS: -rdynamic
// #cgo CFLAGS: -fexceptions -funwind-tables
//
// extern void testSig(int i);
import "C"

import "fmt"

func main() {
	f()
	fmt.Println("Hello, world")
}

func f() {
	C.testSig(0)
}
