package cv

/*
#include <stdlib.h>
#include "imgcodecs.h"
*/
import "C"
import "unsafe"

const (
	FLAG_UNCHANGED           = -1
	FLAG_GRAYSCALE           = 0
	FLAG_COLOR               = 1
	FLAG_ANYDEPTH            = 2
	FLAG_ANYCOLOR            = 4
	FLAG_LOAD_GDAL           = 8
	FLAG_REDUCED_GRAYSCALE_2 = 16
	FLAG_REDUCED_COLOR_2     = 17
	FLAG_REDUCED_GRAYSCALE_4 = 32
	FLAG_REDUCED_COLOR_4     = 33
	FLAG_REDUCED_GRAYSCALE_8 = 64
	FLAG_REDUCED_COLOR_8     = 65
	FLAG_IGNORE_ORIENTATION  = 128
)

func IMRead(filename string, flags int) *Mat {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	return &Mat{C.imread(c_filename, C.int(flags))}
}

func IMDecode(data []byte, flags int) *Mat {
	return &Mat{C.imdecode((*C.uchar)(&data[0]), C.int(len(data)), C.int(flags))}
}

func (m *Mat) IMWrite(filename string) bool {
	c_filename := C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	return bool(C.imwrite(c_filename, m.mat))
}
