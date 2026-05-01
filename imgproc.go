package cv

/*
#include <stdlib.h>
#include "imgproc.h"
*/
import "C"
import "unsafe"

const (
	COLOR_BGR2BGRA = 0 //!< add alpha channel to RGB or BGR image
	COLOR_RGB2RGBA = COLOR_BGR2BGRA

	COLOR_BGRA2BGR = 1 //!< remove alpha channel from RGB or BGR image
	COLOR_RGBA2RGB = COLOR_BGRA2BGR

	COLOR_BGR2RGBA = 2 //!< convert between RGB and BGR color spaces (with or without alpha channel)
	COLOR_RGB2BGRA = COLOR_BGR2RGBA

	COLOR_RGBA2BGR = 3
	COLOR_BGRA2RGB = COLOR_RGBA2BGR

	COLOR_BGR2RGB = 4
	COLOR_RGB2BGR = COLOR_BGR2RGB

	COLOR_BGRA2RGBA = 5
	COLOR_RGBA2BGRA = COLOR_BGRA2RGBA

	COLOR_BGR2GRAY  = 6 //!< convert between RGB/BGR and grayscale, @ref color_convert_rgb_gray "color conversions"
	COLOR_RGB2GRAY  = 7
	COLOR_GRAY2BGR  = 8
	COLOR_GRAY2RGB  = COLOR_GRAY2BGR
	COLOR_GRAY2BGRA = 9
	COLOR_GRAY2RGBA = COLOR_GRAY2BGRA
	COLOR_BGRA2GRAY = 10
	COLOR_RGBA2GRAY = 11

	COLOR_BGR2BGR565  = 12 //!< convert between RGB/BGR and BGR565 (16-bit images)
	COLOR_RGB2BGR565  = 13
	COLOR_BGR5652BGR  = 14
	COLOR_BGR5652RGB  = 15
	COLOR_BGRA2BGR565 = 16
	COLOR_RGBA2BGR565 = 17
	COLOR_BGR5652BGRA = 18
	COLOR_BGR5652RGBA = 19

	COLOR_GRAY2BGR565 = 20 //!< convert between grayscale to BGR565 (16-bit images)
	COLOR_BGR5652GRAY = 21

	COLOR_BGR2BGR555  = 22 //!< convert between RGB/BGR and BGR555 (16-bit images)
	COLOR_RGB2BGR555  = 23
	COLOR_BGR5552BGR  = 24
	COLOR_BGR5552RGB  = 25
	COLOR_BGRA2BGR555 = 26
	COLOR_RGBA2BGR555 = 27
	COLOR_BGR5552BGRA = 28
	COLOR_BGR5552RGBA = 29

	COLOR_GRAY2BGR555 = 30 //!< convert between grayscale and BGR555 (16-bit images)
	COLOR_BGR5552GRAY = 31

	COLOR_BGR2XYZ = 32 //!< convert RGB/BGR to CIE XYZ, @ref color_convert_rgb_xyz "color conversions"
	COLOR_RGB2XYZ = 33
	COLOR_XYZ2BGR = 34
	COLOR_XYZ2RGB = 35

	COLOR_BGR2YCrCb = 36 //!< convert RGB/BGR to luma-chroma (aka YCC), @ref color_convert_rgb_ycrcb "color conversions"
	COLOR_RGB2YCrCb = 37
	COLOR_YCrCb2BGR = 38
	COLOR_YCrCb2RGB = 39

	COLOR_BGR2HSV = 40 //!< convert RGB/BGR to HSV (hue saturation value) with H range 0..180 if 8 bit image, @ref color_convert_rgb_hsv "color conversions"
	COLOR_RGB2HSV = 41

	COLOR_BGR2Lab = 44 //!< convert RGB/BGR to CIE Lab, @ref color_convert_rgb_lab "color conversions"
	COLOR_RGB2Lab = 45

	COLOR_BGR2Luv = 50 //!< convert RGB/BGR to CIE Luv, @ref color_convert_rgb_luv "color conversions"
	COLOR_RGB2Luv = 51
	COLOR_BGR2HLS = 52 //!< convert RGB/BGR to HLS (hue lightness saturation) with H range 0..180 if 8 bit image, @ref color_convert_rgb_hls "color conversions"
	COLOR_RGB2HLS = 53

	COLOR_HSV2BGR = 54 //!< backward conversions HSV to RGB/BGR with H range 0..180 if 8 bit image
	COLOR_HSV2RGB = 55

	COLOR_Lab2BGR = 56
	COLOR_Lab2RGB = 57
	COLOR_Luv2BGR = 58
	COLOR_Luv2RGB = 59
	COLOR_HLS2BGR = 60 //!< backward conversions HLS to RGB/BGR with H range 0..180 if 8 bit image
	COLOR_HLS2RGB = 61

	COLOR_BGR2HSV_FULL = 66 //!< convert RGB/BGR to HSV (hue saturation value) with H range 0..255 if 8 bit image, @ref color_convert_rgb_hsv "color conversions"
	COLOR_RGB2HSV_FULL = 67
	COLOR_BGR2HLS_FULL = 68 //!< convert RGB/BGR to HLS (hue lightness saturation) with H range 0..255 if 8 bit image, @ref color_convert_rgb_hls "color conversions"
	COLOR_RGB2HLS_FULL = 69

	COLOR_HSV2BGR_FULL = 70 //!< backward conversions HSV to RGB/BGR with H range 0..255 if 8 bit image
	COLOR_HSV2RGB_FULL = 71
	COLOR_HLS2BGR_FULL = 72 //!< backward conversions HLS to RGB/BGR with H range 0..255 if 8 bit image
	COLOR_HLS2RGB_FULL = 73

	COLOR_LBGR2Lab = 74
	COLOR_LRGB2Lab = 75
	COLOR_LBGR2Luv = 76
	COLOR_LRGB2Luv = 77

	COLOR_Lab2LBGR = 78
	COLOR_Lab2LRGB = 79
	COLOR_Luv2LBGR = 80
	COLOR_Luv2LRGB = 81

	COLOR_BGR2YUV = 82 //!< convert between RGB/BGR and YUV
	COLOR_RGB2YUV = 83
	COLOR_YUV2BGR = 84
	COLOR_YUV2RGB = 85

	//! YUV 4:2:0 family to RGB
	COLOR_YUV2RGB_NV12 = 90
	COLOR_YUV2BGR_NV12 = 91
	COLOR_YUV2RGB_NV21 = 92
	COLOR_YUV2BGR_NV21 = 93
	COLOR_YUV420sp2RGB = COLOR_YUV2RGB_NV21
	COLOR_YUV420sp2BGR = COLOR_YUV2BGR_NV21

	COLOR_YUV2RGBA_NV12 = 94
	COLOR_YUV2BGRA_NV12 = 95
	COLOR_YUV2RGBA_NV21 = 96
	COLOR_YUV2BGRA_NV21 = 97
	COLOR_YUV420sp2RGBA = COLOR_YUV2RGBA_NV21
	COLOR_YUV420sp2BGRA = COLOR_YUV2BGRA_NV21

	COLOR_YUV2RGB_YV12 = 98
	COLOR_YUV2BGR_YV12 = 99
	COLOR_YUV2RGB_IYUV = 100
	COLOR_YUV2BGR_IYUV = 101
	COLOR_YUV2RGB_I420 = COLOR_YUV2RGB_IYUV
	COLOR_YUV2BGR_I420 = COLOR_YUV2BGR_IYUV
	COLOR_YUV420p2RGB  = COLOR_YUV2RGB_YV12
	COLOR_YUV420p2BGR  = COLOR_YUV2BGR_YV12

	COLOR_YUV2RGBA_YV12 = 102
	COLOR_YUV2BGRA_YV12 = 103
	COLOR_YUV2RGBA_IYUV = 104
	COLOR_YUV2BGRA_IYUV = 105
	COLOR_YUV2RGBA_I420 = COLOR_YUV2RGBA_IYUV
	COLOR_YUV2BGRA_I420 = COLOR_YUV2BGRA_IYUV
	COLOR_YUV420p2RGBA  = COLOR_YUV2RGBA_YV12
	COLOR_YUV420p2BGRA  = COLOR_YUV2BGRA_YV12

	COLOR_YUV2GRAY_420  = 106
	COLOR_YUV2GRAY_NV21 = COLOR_YUV2GRAY_420
	COLOR_YUV2GRAY_NV12 = COLOR_YUV2GRAY_420
	COLOR_YUV2GRAY_YV12 = COLOR_YUV2GRAY_420
	COLOR_YUV2GRAY_IYUV = COLOR_YUV2GRAY_420
	COLOR_YUV2GRAY_I420 = COLOR_YUV2GRAY_420
	COLOR_YUV420sp2GRAY = COLOR_YUV2GRAY_420
	COLOR_YUV420p2GRAY  = COLOR_YUV2GRAY_420

	//! YUV 4:2:2 family to RGB
	COLOR_YUV2RGB_UYVY = 107
	COLOR_YUV2BGR_UYVY = 108
	//COLOR_YUV2RGB_VYUY = 109
	//COLOR_YUV2BGR_VYUY = 110
	COLOR_YUV2RGB_Y422 = COLOR_YUV2RGB_UYVY
	COLOR_YUV2BGR_Y422 = COLOR_YUV2BGR_UYVY
	COLOR_YUV2RGB_UYNV = COLOR_YUV2RGB_UYVY
	COLOR_YUV2BGR_UYNV = COLOR_YUV2BGR_UYVY

	COLOR_YUV2RGBA_UYVY = 111
	COLOR_YUV2BGRA_UYVY = 112
	//COLOR_YUV2RGBA_VYUY = 113
	//COLOR_YUV2BGRA_VYUY = 114
	COLOR_YUV2RGBA_Y422 = COLOR_YUV2RGBA_UYVY
	COLOR_YUV2BGRA_Y422 = COLOR_YUV2BGRA_UYVY
	COLOR_YUV2RGBA_UYNV = COLOR_YUV2RGBA_UYVY
	COLOR_YUV2BGRA_UYNV = COLOR_YUV2BGRA_UYVY

	COLOR_YUV2RGB_YUY2 = 115
	COLOR_YUV2BGR_YUY2 = 116
	COLOR_YUV2RGB_YVYU = 117
	COLOR_YUV2BGR_YVYU = 118
	COLOR_YUV2RGB_YUYV = COLOR_YUV2RGB_YUY2
	COLOR_YUV2BGR_YUYV = COLOR_YUV2BGR_YUY2
	COLOR_YUV2RGB_YUNV = COLOR_YUV2RGB_YUY2
	COLOR_YUV2BGR_YUNV = COLOR_YUV2BGR_YUY2

	COLOR_YUV2RGBA_YUY2 = 119
	COLOR_YUV2BGRA_YUY2 = 120
	COLOR_YUV2RGBA_YVYU = 121
	COLOR_YUV2BGRA_YVYU = 122
	COLOR_YUV2RGBA_YUYV = COLOR_YUV2RGBA_YUY2
	COLOR_YUV2BGRA_YUYV = COLOR_YUV2BGRA_YUY2
	COLOR_YUV2RGBA_YUNV = COLOR_YUV2RGBA_YUY2
	COLOR_YUV2BGRA_YUNV = COLOR_YUV2BGRA_YUY2

	COLOR_YUV2GRAY_UYVY = 123
	COLOR_YUV2GRAY_YUY2 = 124
	//CV_YUV2GRAY_VYUY    = CV_YUV2GRAY_UYVY
	COLOR_YUV2GRAY_Y422 = COLOR_YUV2GRAY_UYVY
	COLOR_YUV2GRAY_UYNV = COLOR_YUV2GRAY_UYVY
	COLOR_YUV2GRAY_YVYU = COLOR_YUV2GRAY_YUY2
	COLOR_YUV2GRAY_YUYV = COLOR_YUV2GRAY_YUY2
	COLOR_YUV2GRAY_YUNV = COLOR_YUV2GRAY_YUY2

	//! alpha premultiplication
	COLOR_RGBA2mRGBA = 125
	COLOR_mRGBA2RGBA = 126

	//! RGB to YUV 4:2:0 family
	COLOR_RGB2YUV_I420 = 127
	COLOR_BGR2YUV_I420 = 128
	COLOR_RGB2YUV_IYUV = COLOR_RGB2YUV_I420
	COLOR_BGR2YUV_IYUV = COLOR_BGR2YUV_I420

	COLOR_RGBA2YUV_I420 = 129
	COLOR_BGRA2YUV_I420 = 130
	COLOR_RGBA2YUV_IYUV = COLOR_RGBA2YUV_I420
	COLOR_BGRA2YUV_IYUV = COLOR_BGRA2YUV_I420
	COLOR_RGB2YUV_YV12  = 131
	COLOR_BGR2YUV_YV12  = 132
	COLOR_RGBA2YUV_YV12 = 133
	COLOR_BGRA2YUV_YV12 = 134

	//! Demosaicing, see @ref color_convert_bayer "color conversions" for additional information
	COLOR_BayerBG2BGR = 46 //!< equivalent to RGGB Bayer pattern
	COLOR_BayerGB2BGR = 47 //!< equivalent to GRBG Bayer pattern
	COLOR_BayerRG2BGR = 48 //!< equivalent to BGGR Bayer pattern
	COLOR_BayerGR2BGR = 49 //!< equivalent to GBRG Bayer pattern

	COLOR_BayerRGGB2BGR = COLOR_BayerBG2BGR
	COLOR_BayerGRBG2BGR = COLOR_BayerGB2BGR
	COLOR_BayerBGGR2BGR = COLOR_BayerRG2BGR
	COLOR_BayerGBRG2BGR = COLOR_BayerGR2BGR

	COLOR_BayerRGGB2RGB = COLOR_BayerBGGR2BGR
	COLOR_BayerGRBG2RGB = COLOR_BayerGBRG2BGR
	COLOR_BayerBGGR2RGB = COLOR_BayerRGGB2BGR
	COLOR_BayerGBRG2RGB = COLOR_BayerGRBG2BGR

	COLOR_BayerBG2RGB = COLOR_BayerRG2BGR //!< equivalent to RGGB Bayer pattern
	COLOR_BayerGB2RGB = COLOR_BayerGR2BGR //!< equivalent to GRBG Bayer pattern
	COLOR_BayerRG2RGB = COLOR_BayerBG2BGR //!< equivalent to BGGR Bayer pattern
	COLOR_BayerGR2RGB = COLOR_BayerGB2BGR //!< equivalent to GBRG Bayer pattern

	COLOR_BayerBG2GRAY = 86 //!< equivalent to RGGB Bayer pattern
	COLOR_BayerGB2GRAY = 87 //!< equivalent to GRBG Bayer pattern
	COLOR_BayerRG2GRAY = 88 //!< equivalent to BGGR Bayer pattern
	COLOR_BayerGR2GRAY = 89 //!< equivalent to GBRG Bayer pattern

	COLOR_BayerRGGB2GRAY = COLOR_BayerBG2GRAY
	COLOR_BayerGRBG2GRAY = COLOR_BayerGB2GRAY
	COLOR_BayerBGGR2GRAY = COLOR_BayerRG2GRAY
	COLOR_BayerGBRG2GRAY = COLOR_BayerGR2GRAY

	//! Demosaicing using Variable Number of Gradients
	COLOR_BayerBG2BGR_VNG = 62 //!< equivalent to RGGB Bayer pattern
	COLOR_BayerGB2BGR_VNG = 63 //!< equivalent to GRBG Bayer pattern
	COLOR_BayerRG2BGR_VNG = 64 //!< equivalent to BGGR Bayer pattern
	COLOR_BayerGR2BGR_VNG = 65 //!< equivalent to GBRG Bayer pattern

	COLOR_BayerRGGB2BGR_VNG = COLOR_BayerBG2BGR_VNG
	COLOR_BayerGRBG2BGR_VNG = COLOR_BayerGB2BGR_VNG
	COLOR_BayerBGGR2BGR_VNG = COLOR_BayerRG2BGR_VNG
	COLOR_BayerGBRG2BGR_VNG = COLOR_BayerGR2BGR_VNG

	COLOR_BayerRGGB2RGB_VNG = COLOR_BayerBGGR2BGR_VNG
	COLOR_BayerGRBG2RGB_VNG = COLOR_BayerGBRG2BGR_VNG
	COLOR_BayerBGGR2RGB_VNG = COLOR_BayerRGGB2BGR_VNG
	COLOR_BayerGBRG2RGB_VNG = COLOR_BayerGRBG2BGR_VNG

	COLOR_BayerBG2RGB_VNG = COLOR_BayerRG2BGR_VNG //!< equivalent to RGGB Bayer pattern
	COLOR_BayerGB2RGB_VNG = COLOR_BayerGR2BGR_VNG //!< equivalent to GRBG Bayer pattern
	COLOR_BayerRG2RGB_VNG = COLOR_BayerBG2BGR_VNG //!< equivalent to BGGR Bayer pattern
	COLOR_BayerGR2RGB_VNG = COLOR_BayerGB2BGR_VNG //!< equivalent to GBRG Bayer pattern

	//! Edge-Aware Demosaicing
	COLOR_BayerBG2BGR_EA = 135 //!< equivalent to RGGB Bayer pattern
	COLOR_BayerGB2BGR_EA = 136 //!< equivalent to GRBG Bayer pattern
	COLOR_BayerRG2BGR_EA = 137 //!< equivalent to BGGR Bayer pattern
	COLOR_BayerGR2BGR_EA = 138 //!< equivalent to GBRG Bayer pattern

	COLOR_BayerRGGB2BGR_EA = COLOR_BayerBG2BGR_EA
	COLOR_BayerGRBG2BGR_EA = COLOR_BayerGB2BGR_EA
	COLOR_BayerBGGR2BGR_EA = COLOR_BayerRG2BGR_EA
	COLOR_BayerGBRG2BGR_EA = COLOR_BayerGR2BGR_EA

	COLOR_BayerRGGB2RGB_EA = COLOR_BayerBGGR2BGR_EA
	COLOR_BayerGRBG2RGB_EA = COLOR_BayerGBRG2BGR_EA
	COLOR_BayerBGGR2RGB_EA = COLOR_BayerRGGB2BGR_EA
	COLOR_BayerGBRG2RGB_EA = COLOR_BayerGRBG2BGR_EA

	COLOR_BayerBG2RGB_EA = COLOR_BayerRG2BGR_EA //!< equivalent to RGGB Bayer pattern
	COLOR_BayerGB2RGB_EA = COLOR_BayerGR2BGR_EA //!< equivalent to GRBG Bayer pattern
	COLOR_BayerRG2RGB_EA = COLOR_BayerBG2BGR_EA //!< equivalent to BGGR Bayer pattern
	COLOR_BayerGR2RGB_EA = COLOR_BayerGB2BGR_EA //!< equivalent to GBRG Bayer pattern

	//! Demosaicing with alpha channel
	COLOR_BayerBG2BGRA = 139 //!< equivalent to RGGB Bayer pattern
	COLOR_BayerGB2BGRA = 140 //!< equivalent to GRBG Bayer pattern
	COLOR_BayerRG2BGRA = 141 //!< equivalent to BGGR Bayer pattern
	COLOR_BayerGR2BGRA = 142 //!< equivalent to GBRG Bayer pattern

	COLOR_BayerRGGB2BGRA = COLOR_BayerBG2BGRA
	COLOR_BayerGRBG2BGRA = COLOR_BayerGB2BGRA
	COLOR_BayerBGGR2BGRA = COLOR_BayerRG2BGRA
	COLOR_BayerGBRG2BGRA = COLOR_BayerGR2BGRA

	COLOR_BayerRGGB2RGBA = COLOR_BayerBGGR2BGRA
	COLOR_BayerGRBG2RGBA = COLOR_BayerGBRG2BGRA
	COLOR_BayerBGGR2RGBA = COLOR_BayerRGGB2BGRA
	COLOR_BayerGBRG2RGBA = COLOR_BayerGRBG2BGRA

	COLOR_BayerBG2RGBA = COLOR_BayerRG2BGRA //!< equivalent to RGGB Bayer pattern
	COLOR_BayerGB2RGBA = COLOR_BayerGR2BGRA //!< equivalent to GRBG Bayer pattern
	COLOR_BayerRG2RGBA = COLOR_BayerBG2BGRA //!< equivalent to BGGR Bayer pattern
	COLOR_BayerGR2RGBA = COLOR_BayerGB2BGRA //!< equivalent to GBRG Bayer pattern

	//! RGB to YUV 4:2:2 family

	COLOR_RGB2YUV_UYVY = 143
	COLOR_BGR2YUV_UYVY = 144
	COLOR_RGB2YUV_Y422 = COLOR_RGB2YUV_UYVY
	COLOR_BGR2YUV_Y422 = COLOR_BGR2YUV_UYVY
	COLOR_RGB2YUV_UYNV = COLOR_RGB2YUV_UYVY
	COLOR_BGR2YUV_UYNV = COLOR_BGR2YUV_UYVY

	COLOR_RGBA2YUV_UYVY = 145
	COLOR_BGRA2YUV_UYVY = 146
	COLOR_RGBA2YUV_Y422 = COLOR_RGBA2YUV_UYVY
	COLOR_BGRA2YUV_Y422 = COLOR_BGRA2YUV_UYVY
	COLOR_RGBA2YUV_UYNV = COLOR_RGBA2YUV_UYVY
	COLOR_BGRA2YUV_UYNV = COLOR_BGRA2YUV_UYVY

	COLOR_RGB2YUV_YUY2 = 147
	COLOR_BGR2YUV_YUY2 = 148
	COLOR_RGB2YUV_YVYU = 149
	COLOR_BGR2YUV_YVYU = 150
	COLOR_RGB2YUV_YUYV = COLOR_RGB2YUV_YUY2
	COLOR_BGR2YUV_YUYV = COLOR_BGR2YUV_YUY2
	COLOR_RGB2YUV_YUNV = COLOR_RGB2YUV_YUY2
	COLOR_BGR2YUV_YUNV = COLOR_BGR2YUV_YUY2

	COLOR_RGBA2YUV_YUY2 = 151
	COLOR_BGRA2YUV_YUY2 = 152
	COLOR_RGBA2YUV_YVYU = 153
	COLOR_BGRA2YUV_YVYU = 154
	COLOR_RGBA2YUV_YUYV = COLOR_RGBA2YUV_YUY2
	COLOR_BGRA2YUV_YUYV = COLOR_BGRA2YUV_YUY2
	COLOR_RGBA2YUV_YUNV = COLOR_RGBA2YUV_YUY2
	COLOR_BGRA2YUV_YUNV = COLOR_BGRA2YUV_YUY2

	COLOR_COLORCVT_MAX = 155
)

const (
	INTER_NEAREST = iota
	INTER_LINEAR
	INTER_CUBIC
	INTER_AREA
	INTER_LANCZOS4
	INTER_LINEAR_EXACT
	INTER_NEAREST_EXACT
	INTER_MAX
)

const (
	BORDER_CONSTANT = iota
	BORDER_REPLICATE
	BORDER_REFLECT
	BORDER_WRAP
	BORDER_REFLECT_101
	BORDER_TRANSPARENT

	BORDER_REFLECT101 = BORDER_REFLECT_101
	BORDER_DEFAULT    = BORDER_REFLECT_101
	BORDER_ISOLATED   = 16
)

const (
	ROTATE_90_CLOCKWISE = iota
	ROTATE_180
	ROTATE_90_COUNTERCLOCKWISE
)

const (
	FONT_HERSHEY_SIMPLEX = iota
	FONT_HERSHEY_PLAIN
	FONT_HERSHEY_DUPLEX
	FONT_HERSHEY_COMPLEX
	FONT_HERSHEY_TRIPLEX
	FONT_HERSHEY_COMPLEX_SMALL
	FONT_HERSHEY_SCRIPT_SIMPLEX
	FONT_HERSHEY_SCRIPT_COMPLEX
	FONT_ITALIC = 16
)

const (
	RETR_EXTERNAL = iota
	RETR_LIST
	RETR_CCOMP
	RETR_TREE
	RETR_FLOODFILL
)

const (
	CHAIN_APPROX_NONE = iota + 1
	CHAIN_APPROX_SIMPLE
	CHAIN_APPROX_TC89_L1
	CHAIN_APPROX_TC89_KCOS
)

type AreaRect struct {
	Points [4]C.Point2f
	Size   C.Size2f
	Angle  float32
}

func (ar *AreaRect) GetPoints() (points [][2]float32) {
	points = make([][2]float32, 4)
	for i := 0; i < 4; i++ {
		points[i][0] = float32(ar.Points[i].x)
		points[i][1] = float32(ar.Points[i].y)
	}
	return
}

func (ar *AreaRect) GetWidth() float32 {
	return float32(ar.Size.width)
}

func (ar *AreaRect) GetHeight() float32 {
	return float32(ar.Size.height)
}

func (m *Mat) CvtColor(code int) *Mat {
	dst := NewMat()
	C.cvt_color(m.mat, dst.mat, C.int(code))
	return dst
}

func (m *Mat) Resize(size Size, fx, fy float32, interp int) *Mat {
	dst := NewMat()
	C.resize(m.mat, dst.mat, C.Size(size), C.float(fx), C.float(fy), C.int(interp))
	return dst
}

func (m *Mat) Rotate(code int) *Mat {
	dst := NewMat()
	C.rotate(m.mat, dst.mat, C.int(code))
	return dst
}

// borderType: Border type, one of the BorderTypes, except for BORDER_TRANSPARENT and BORDER_ISOLATED.
func (m *Mat) CopyMakeBorder(top, bottom, left, right, borderType int, color Scalar) *Mat {
	dst := NewMat()
	C.copy_make_border(m.mat, dst.mat, C.int(top), C.int(bottom), C.int(left), C.int(right), C.int(borderType), C.Scalar(color))
	return dst
}

func (m *Mat) WarpAffine(n *Mat, size Size) *Mat {
	dst := NewMat()
	C.warp_affine(m.mat, dst.mat, n.mat, C.Size(size))
	return dst
}

func (m *Mat) WarpPerspective(n *Mat, size Size) *Mat {
	dst := NewMat()
	C.warp_perspective(m.mat, dst.mat, n.mat, C.Size(size))
	return dst
}

// m: Source arrays. They all should have the same depth, CV_8U, CV_16U or CV_32F , and the same size. Each of them can have an arbitrary number of channels.
func (m *Mat) CalcHist() *Mat {
	dst := NewMat()
	C.calc_hist(m.mat, dst.mat)
	return dst
}

func (m *Mat) Blur(size Size) *Mat {
	dst := NewMat()
	C.blur(m.mat, dst.mat, C.Size(size))
	return dst
}

func (m *Mat) BlurSelf(size Size) {
	C.blur(m.mat, m.mat, C.Size(size))
}

func (m *Mat) GaussianBlur(size Size, sigmaX, sigmaY float32, borderType int) *Mat {
	dst := NewMat()
	C.gaussian_blur(m.mat, dst.mat, C.Size(size), C.float(sigmaX), C.float(sigmaY), C.int(borderType))
	return dst
}

func (m *Mat) GaussianBlurSelf(size Size, sigmaX, sigmaY float32, borderType int) {
	C.gaussian_blur(m.mat, m.mat, C.Size(size), C.float(sigmaX), C.float(sigmaY), C.int(borderType))
}

func (m *Mat) Canny(threshold1, threshold2 float32) *Mat {
	edges := NewMat()
	C.canny(m.mat, edges.mat, C.float(threshold1), C.float(threshold2))
	return edges
}

func (m *Mat) Circle(center Point, radius int, color Scalar, thickness int) {
	C.circle(m.mat, C.Point(center), C.int(radius), C.Scalar(color), C.int(thickness))
}

func (m *Mat) Ellipse(center Point, axes Size, angle, startAngle, endAngle float32, color Scalar, thickness int) {
	C.ellipse(m.mat, C.Point(center), C.Size(axes), C.float(angle), C.float(startAngle), C.float(endAngle), C.Scalar(color), C.int(thickness))
}

func (m *Mat) Line(p1, p2 Point, color Scalar, thickness int) {
	C.line(m.mat, C.Point(p1), C.Point(p2), C.Scalar(color), C.int(thickness))
}

func (m *Mat) Rectangle(rect Rect, color Scalar, thickness int) {
	C.rectangle(m.mat, C.Rect(rect), C.Scalar(color), C.int(thickness))
}

func (m *Mat) FillPoly(points []Point, color Scalar) {
	C.fill_poly(m.mat, (*C.Point)(&points[0]), C.int(len(points)), C.Scalar(color))
}

func (m *Mat) Polylines(points []Point, isClosed bool, color Scalar, thickness int) {
	C.polylines(m.mat, (*C.Point)(&points[0]), C.int(len(points)), C.bool(isClosed), C.Scalar(color), C.int(thickness))
}

func (m *Mat) PutText(text string, org Point, fontFace int, fontScale float32, color Scalar, thickness int) {
	c_text := C.CString(text)
	defer C.free(unsafe.Pointer(c_text))
	C.put_text(m.mat, c_text, C.Point(org), C.int(fontFace), C.float(fontScale), C.Scalar(color), C.int(thickness))
}

func (m *Mat) Watershed(markers *Mat) {
	C.watershed(m.mat, markers.mat)
}

func MinAreaRect(points []Point) *AreaRect {
	c_areaRect := C.min_area_rect((*C.Point)(unsafe.Pointer(&points[0])), C.int(len(points)))
	return &AreaRect{
		c_areaRect.points,
		c_areaRect.size,
		float32(c_areaRect.angle),
	}
}

func (m *Mat) BoxScore(box *AreaRect) float32 {
	return float32(C.box_score(m.mat, &box.Points[0]))
}

// m: an 8-bit single-channel image.
func (m *Mat) FindContours(mode, method int) [][]Point {
	c_contours := C.find_contours(m.mat, C.int(mode), C.int(method))
	num := int(c_contours.num)
	counts := unsafe.Slice(c_contours.counts, num)
	c_pointsList := unsafe.Slice(c_contours.points, num)
	contours := make([][]Point, num)
	for i := 0; i < num; i++ {
		count := int(counts[i])
		points := make([]Point, count)
		c_points := unsafe.Slice(c_pointsList[i], count)
		for j := 0; j < count; j++ {
			points[j] = Point(c_points[j])
		}
		contours[i] = points
	}
	C.free_contours(c_contours)
	return contours
}

func (m *Mat) FindAreaRects(mode, method int) []*AreaRect {
	c_areaRects := C.find_area_rects(m.mat, C.int(mode), C.int(method))
	num := int(c_areaRects.num)
	c_rects := unsafe.Slice(c_areaRects.rects, num)
	rects := make([]*AreaRect, num)
	for i := 0; i < num; i++ {
		rects[i] = &AreaRect{
			c_rects[i].points,
			c_rects[i].size,
			float32(c_rects[i].angle),
		}
	}
	C.free_area_rects(c_areaRects)
	return rects
}

func CompareHist(hist1, hist2 *Mat) float32 {
	return float32(C.compare_hist(hist1.mat, hist2.mat))
}

func GetAffineTransform(src, dst []Point2f) *Mat {
	return &Mat{C.get_affine_transform((*C.Point2f)(&src[0]), (*C.Point2f)(&dst[0]), C.int(len(src)))}
}

func GetPerspectiveTransform(src, dst []Point2f) *Mat {
	return &Mat{C.get_perspective_transform((*C.Point2f)(&src[0]), (*C.Point2f)(&dst[0]), C.int(len(src)))}
}

func GetRotationMatrix2D(center Point2f, angle, scale float32) *Mat {
	return &Mat{C.get_rotation_matrix_2d(C.Point2f(center), C.float(angle), C.float(scale))}
}

func GetSimilarityTransformMatrix(landmarks [10]float32) *Mat {
	return &Mat{C.get_similarity_transform_matrix((*C.float)(&landmarks[0]))}
}
