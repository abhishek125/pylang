package main

import (
	"fmt"
	"syscall"
)

// #cgo LDFLAGS: -L. -lexample
// #include "example.h"
import "C"

func main() {
	lib, err := syscall.LoadLibrary("./example.so")
	if err != nil {
		fmt.Println("Error loading library:", err)
		return
	}
	defer syscall.FreeLibrary(lib)

	// Call the Golang function that invokes the Python function
	C.PrintString(C.CString("Hello from Python"))
}
