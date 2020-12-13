package main

import (
	"raytracing-tutorials/utils"
)

func main() {
	w, h := 200, 100
	img := utils.NewRGBA(w, h)

	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			r := float32(i) / float32(w) * 255.999
			g := float32(j) / float32(h) * 255.999
			b := float32(0.2) * 255.999
			img.SetRGBA(i, j, utils.RGBAFC{r, g, b, 255.0})
		}
	}

	img.SavePNG("result.png")
}
