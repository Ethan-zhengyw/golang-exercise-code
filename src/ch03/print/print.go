// Package print implements C print method
//
// xxxxxxxxxxxx
// xxxxxxxxxxxx
package print

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

// Print create c string and put it to file
func Print(s string)  {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}
