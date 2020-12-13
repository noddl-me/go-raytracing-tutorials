package main

import (
	"raytracing-tutorials/utils"
)

func main() {
	w, h := 200, 100
	img := utils.NewRGBA(w, h)

	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			r := float32(i) / float32(w)
			g := float32(j) / float32(h)
			b := float32(0.2)
			img.SetRGB(i, j, utils.NewVector3(r, g, b))
		}
	}

	img.SavePNG("result.png")
}
