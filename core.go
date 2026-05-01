package cv

// #include "core.h"
import "C"
import (
	"unsafe"
)

const (
	MatType8U = iota
	MatType8S
	MatType16U
	MatType16S
	MatType32S
	MatType32F
	MatType64F
	MatType16F
	MatType8UC1  = MatType8U | MatChannels1
	MatType8UC2  = MatType8U | MatChannels2
	MatType8UC3  = MatType8U | MatChannels3
	MatType8UC4  = MatType8U | MatChannels4
	MatType8SC1  = MatType8S | MatChannels1
	MatType8SC2  = MatType8S | MatChannels2
	MatType8SC3  = MatType8S | MatChannels3
	MatType8SC4  = MatType8S | MatChannels4
	MatType16UC1 = MatType16U | MatChannels1
	MatType16UC2 = MatType16U | MatChannels2
	MatType16UC3 = MatType16U | MatChannels3
	MatType16UC4 = MatType16U | MatChannels4
	MatType16SC1 = MatType16S | MatChannels1
	MatType16SC2 = MatType16S | MatChannels2
	MatType16SC3 = MatType16S | MatChannels3
	MatType16SC4 = MatType16S | MatChannels4
	MatType32SC1 = MatType32S | MatChannels1
	MatType32SC2 = MatType32S | MatChannels2
	MatType32SC3 = MatType32S | MatChannels3
	MatType32SC4 = MatType32S | MatChannels4
	MatType32FC1 = MatType32F | MatChannels1
	MatType32FC2 = MatType32F | MatChannels2
	MatType32FC3 = MatType32F | MatChannels3
	MatType32FC4 = MatType32F | MatChannels4
	MatType64FC1 = MatType64F | MatChannels1
	MatType64FC2 = MatType64F | MatChannels2
	MatType64FC3 = MatType64F | MatChannels3
	MatType64FC4 = MatType64F | MatChannels4
	MatType16FC1 = MatType16F | MatChannels1
	MatType16FC2 = MatType16F | MatChannels2
	MatType16FC3 = MatType16F | MatChannels3
	MatType16FC4 = MatType16F | MatChannels4
)

const (
	MatChannels1 = iota << 3
	MatChannels2
	MatChannels3
	MatChannels4
)

func GetNumThreads() int {
	return int(C.get_num_threads())
}

func SetNumThreads(n int) {
	C.set_num_threads(C.int(n))
}

type Point C.Point

func NewPoint(x, y int) Point {
	return Point{C.int(x), C.int(y)}
}

type Point2f C.Point2f

func NewPoint2f(x, y float32) Point2f {
	return Point2f{C.float(x), C.float(y)}
}

type Size C.Size

func NewSize(width, height int) Size {
	return Size{C.int(width), C.int(height)}
}

type Rect C.Rect

func NewRect(x, y, width, height int) Rect {
	return Rect{C.int(x), C.int(y), C.int(width), C.int(height)}
}

type Scalar C.Scalar

func NewScalar(v1, v2, v3, v4 float32) Scalar {
	return Scalar{C.float(v1), C.float(v2), C.float(v3), C.float(v4)}
}

type Mat struct {
	mat C.Mat
}

func (m *Mat) Empty() bool {
	return bool(C.mat_empty(m.mat))
}

func (m *Mat) Clone() *Mat {
	return &Mat{C.mat_clone(m.mat)}
}

func (m *Mat) ConvertTo(mt int) *Mat {
	dst := NewMat()
	C.mat_convert_to(m.mat, dst.mat, C.int(mt))
	return dst
}

func (m *Mat) ConvertToAB(mt int, alpha, beta float32) *Mat {
	dst := NewMat()
	C.mat_convert_to_ab(m.mat, dst.mat, C.int(mt), C.float(alpha), C.float(beta))
	return dst
}

func (m *Mat) Region(r Rect) *Mat {
	return &Mat{C.mat_region(m.mat, C.Rect(r))}
}

func (m *Mat) Reshape(cn, rows int) *Mat {
	return &Mat{C.mat_reshape(m.mat, C.int(cn), C.int(rows))}
}

func (m *Mat) Rows() int {
	return int(C.mat_rows(m.mat))
}

func (m *Mat) Cols() int {
	return int(C.mat_cols(m.mat))
}

func (m *Mat) Channels() int {
	return int(C.mat_channels(m.mat))
}

func (m *Mat) Type() int {
	return int(C.mat_type(m.mat))
}

func (m *Mat) Total() int {
	return int(C.mat_total(m.mat))
}

func (m *Mat) ElemSize() int {
	return int(C.mat_elem_size(m.mat))
}

func (m *Mat) Data() unsafe.Pointer {
	return C.mat_data(m.mat)
}

func (m *Mat) Mul(n *Mat, scale float32) *Mat {
	return &Mat{C.mat_mul(m.mat, n.mat, C.float(scale))}
}

func (m *Mat) Divide(s Scalar) {
	C.mat_divide(m.mat, C.Scalar(s))
}

func (m *Mat) AddByte(b byte) {
	C.mat_add_uchar(m.mat, C.uchar(b))
}

func (m *Mat) SubtractByte(b byte) {
	C.mat_subtract_uchar(m.mat, C.uchar(b))
}

func (m *Mat) MultiplyByte(b byte) {
	C.mat_multiply_uchar(m.mat, C.uchar(b))
}

func (m *Mat) DivideByte(b byte) {
	C.mat_divide_uchar(m.mat, C.uchar(b))
}

func (m *Mat) AddFloat(v float32) {
	C.mat_add_float(m.mat, C.float(v))
}

func (m *Mat) SubtractFloat(v float32) {
	C.mat_subtract_float(m.mat, C.float(v))
}

func (m *Mat) MultiplyFloat(v float32) {
	C.mat_multiply_float(m.mat, C.float(v))
}

func (m *Mat) DivideFloat(v float32) {
	C.mat_divide_float(m.mat, C.float(v))
}

func (m *Mat) AddMat(n *Mat) {
	C.mat_add_mat(m.mat, n.mat)
}

func (m *Mat) SubtractMat(n *Mat) {
	C.mat_subtract_mat(m.mat, n.mat)
}

func (m *Mat) MultiplyMatrix(n *Mat) *Mat {
	return &Mat{C.mat_multiply_matrix(m.mat, n.mat)}
}

func (m *Mat) T() *Mat {
	return &Mat{C.mat_t(m.mat)}
}

func (m *Mat) Split() []Mat {
	ms := make([]Mat, m.Channels())
	C.mat_split(m.mat, (*C.Mat)(unsafe.Pointer(&ms[0])))
	return ms
}

func (m *Mat) Free() {
	if m.mat != nil {
		C.mat_free(m.mat)
		m.mat = nil
	}
}

func NewMat() *Mat {
	return &Mat{C.mat_alloc()}
}

func NewMatWithData(rows, cols, mt int, data unsafe.Pointer) *Mat {
	return &Mat{C.mat_alloc_with_data(C.int(rows), C.int(cols), C.int(mt), data)}
}

func Merge(ms []Mat) *Mat {
	dst := NewMat()
	C.mat_merge((*C.Mat)(unsafe.Pointer(&ms[0])), C.int(len(ms)), dst.mat)
	return dst
}
