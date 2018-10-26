package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s){
	printf("%s\n", s);
}
*/
import "C"

import (
	"unsafe"
	"fmt"
)

func main() {
	msg := "Hello from stdio\n"
	cmsg := C.CString(msg)
	defer C.free(unsafe.Pointer(cmsg))

	C.myprint(cmsg)
	fmt.Println("done...")
}
