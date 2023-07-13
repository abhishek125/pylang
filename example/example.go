package main

// #include "example_wrap.c"
import "C"
import "unsafe"

// PrintString calls the Python function to print a string.
func PrintString(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C.print_string(cstr)
}

