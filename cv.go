package cv

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo amd64 LDFLAGS: -L${SRCDIR}/amd64
#cgo arm64 LDFLAGS: -L${SRCDIR}/arm64
#cgo LDFLAGS: -lcv
*/
import "C"

type ResizeBorder struct {
	Ratio  float32
	Width  int
	Height int
	Top    int
	Bottom int
	Left   int
	Right  int
}

type Resize32Ratio struct {
	Width  int
	Height int
	ScaleX float32
	ScaleY float32
}

func (m *Mat) Resize32(maxSize, interp int) (*Mat, *Resize32Ratio) {
	originWidth := m.Cols()
	originHeight := m.Rows()
	width := originWidth
	height := originHeight
	if originWidth > originHeight {
		if originWidth > maxSize {
			width = maxSize
			height = maxSize * originHeight / originWidth
		}
	} else if originHeight > maxSize {
		width = maxSize * originWidth / originHeight
		height = maxSize
	}
	rr := &Resize32Ratio{
		Width:  width >> 5 << 5,
		Height: height >> 5 << 5,
	}
	rr.ScaleX = float32(originWidth) / float32(rr.Width)
	rr.ScaleY = float32(originHeight) / float32(rr.Height)
	return m.Resize(NewSize(rr.Width, rr.Height), 0, 0, interp), rr
}

func (m *Mat) ResizeAndCrop(width, height, interp int) *Mat {
	cols := m.Cols()
	rows := m.Rows()
	var x, y, w, h int
	if width*rows > height*cols {
		w = width
		h = rows * width / cols
		y = (h - height) >> 1
		if y < 0 {
			y = -y
		}
	} else {
		w = cols * height / rows
		h = height
		x = (w - width) >> 1
		if x < 0 {
			x = -x
		}
	}
	outer := m.Resize(NewSize(w, h), 0, 0, interp)
	resized := outer.Region(NewRect(x, y, width, height))
	outer.Free()
	return resized
}

func (m *Mat) ResizeWithBorder(width, height, interp, borderType int, color Scalar) (*Mat, *ResizeBorder) {
	var border ResizeBorder
	cols := float32(m.Cols())
	rows := float32(m.Rows())
	border.Ratio = Minf(float32(width)/cols, float32(height)/rows)
	border.Width = int(cols * border.Ratio)
	border.Height = int(rows * border.Ratio)
	dx := width - border.Width
	dy := height - border.Height
	border.Top = dy >> 1
	border.Bottom = dy - border.Top
	border.Left = dx >> 1
	border.Right = dx - border.Left
	inner := m.Resize(NewSize(border.Width, border.Height), 0, 0, interp)
	resized := inner.CopyMakeBorder(border.Top, border.Bottom, border.Left, border.Right, borderType, color)
	inner.Free()
	return resized, &border
}

func (m *Mat) AlignCropRect(landmarks [8]float32) (*Mat, float32, float32) {
	width := Maxf(
		Distancef(landmarks[0], landmarks[1], landmarks[2], landmarks[3]),
		Distancef(landmarks[4], landmarks[5], landmarks[6], landmarks[7]),
	)
	height := Maxf(
		Distancef(landmarks[0], landmarks[1], landmarks[6], landmarks[7]),
		Distancef(landmarks[2], landmarks[3], landmarks[4], landmarks[5]),
	)
	tm := GetPerspectiveTransform(
		[]Point2f{
			NewPoint2f(landmarks[0], landmarks[1]),
			NewPoint2f(landmarks[2], landmarks[3]),
			NewPoint2f(landmarks[4], landmarks[5]),
			NewPoint2f(landmarks[6], landmarks[7]),
		},
		[]Point2f{
			NewPoint2f(0, 0),
			NewPoint2f(width, 0),
			NewPoint2f(width, height),
			NewPoint2f(0, height),
		},
	)
	aligned := m.WarpPerspective(tm, NewSize(int(width), int(height)))
	tm.Free()
	if width*1.5 < height {
		rotated := aligned.Rotate(ROTATE_90_CLOCKWISE)
		aligned.Free()
		return rotated, height, width
	}
	return aligned, width, height
}

func (m *Mat) AlignCropFace(landmarks [10]float32) *Mat {
	tm := GetSimilarityTransformMatrix(landmarks)
	aligned := m.WarpAffine(tm, NewSize(112, 112))
	tm.Free()
	return aligned
}
