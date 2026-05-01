package cv

/*
#include <stdlib.h>
#include "dnn.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type Names struct {
	names C.Names
}

func (n Names) Size() int {
	return int(C.net_names_size(n.names))
}

func (n Names) At(index int) string {
	return C.GoString(C.net_name_at(n.names, C.int(index)))
}

func (n Names) Free() {
	C.net_names_free(n.names)
}

type Net struct {
	net   C.Net
	names Names
}

func (n *Net) Empty() bool {
	return bool(C.net_empty(n.net))
}

func (n *Net) SetInput(blob *Mat) {
	C.net_set_input(n.net, blob.mat)
}

func (n *Net) Forward() []Mat {
	count := n.names.Size()
	outputBlobs := make([]Mat, count)
	C.net_forward(n.net, (*C.Mat)(unsafe.Pointer(&outputBlobs[0])), n.names.names, C.int(count))
	return outputBlobs
}

func (n *Net) SetPreferableBackend(backend int) {
	C.net_set_preferable_backend(n.net, C.int(backend))
}

func (n *Net) SetPreferableTarget(target int) {
	C.net_set_preferable_target(n.net, C.int(target))
}

func (n *Net) GetUnconnectedOutLayersNames() Names {
	return n.names
}

func (n *Net) Free() {
	if n.net != nil {
		C.net_free(n.net)
		n.net = nil
	}
}

func NewNet(model, config string) (*Net, error) {
	c_model := C.CString(model)
	c_config := C.CString(config)
	defer C.free(unsafe.Pointer(c_model))
	defer C.free(unsafe.Pointer(c_config))
	return newNet(C.net_alloc(c_model, c_config))
}

func NewNetWithBuffer(framework string, model, config []byte) (*Net, error) {
	c_framework := C.CString(framework)
	defer C.free(unsafe.Pointer(c_framework))
	return newNet(C.net_alloc_with_buffer(c_framework, (*C.uchar)(&model[0]), C.int(len(model)), (*C.uchar)(&config[0]), C.int(len(config))))
}

func newNet(c_net C.Net) (*Net, error) {
	net := &Net{net: c_net}
	if net.Empty() {
		return net, errors.New("net is empty")
	}
	net.names = Names{C.net_get_unconnected_out_layers_names(c_net)}
	return net, nil
}

func (m *Mat) BlobFromImage(scale float32, size Size, mean Scalar, swapRB, crop bool) *Mat {
	return &Mat{C.net_blob_from_image(m.mat, C.float(scale), C.Size(size), C.Scalar(mean), C.bool(swapRB), C.bool(crop))}
}
