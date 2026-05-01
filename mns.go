package cv

import "sort"

type Box struct {
	ID    int
	X0    float32
	Y0    float32
	X1    float32
	Y1    float32
	Score float32
	Data  interface{}
}

func NMS(input []*Box, iou float32) []*Box {
	n := len(input)
	if n < 2 {
		return input
	}
	sort.Slice(input, func(i, j int) bool {
		return input[i].Score > input[j].Score
	})
	merged := make(map[int]struct{}, n)
	var output []*Box
	for i := 0; i < n; i++ {
		if _, ok := merged[i]; ok {
			continue
		}
		box0 := input[i]
		output = append(output, box0)
		area0 := (box0.X1 - box0.X0) * (box0.Y1 - box0.Y0)
		for j := i + 1; j < n; j++ {
			if _, ok := merged[j]; ok {
				continue
			}
			box1 := input[j]
			innerX0 := Maxf(box0.X0, box1.X0)
			innerY0 := Maxf(box0.Y0, box1.Y0)
			innerX1 := Minf(box0.X1, box1.X1)
			innerY1 := Minf(box0.Y1, box1.Y1)

			innerW := innerX1 - innerX0
			innerH := innerY1 - innerY0

			if innerW <= 0 || innerH <= 0 {
				continue
			}

			innerArea := innerW * innerH

			area1 := (box1.X1 - box1.X0) * (box1.Y1 - box1.Y0)

			if innerArea/(area0+area1-innerArea) > iou {
				merged[j] = struct{}{}
			}
		}
	}
	return output
}
