package main

// extern void testSig(int i);
import "C"

import "fmt"

func main() {
	C.testSig(0)
	fmt.Println("this never happens")
}
