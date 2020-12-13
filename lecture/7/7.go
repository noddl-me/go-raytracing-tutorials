package main

import (
	"raytracing-tutorials/utils"
)

func RayColor(r *utils.Ray, world utils.HitTable) *utils.Vector3 {
	if rec := world.Hit(r, 0, utils.Inf32(1)); rec != nil {
		rec.Normal.Add(utils.NewVector3(1.0, 1.0, 1.0))
		rec.Normal.MulF(0.5)
		return rec.Normal
	}
	t := 0.5 * (r.Direction.Unit().Y() + 1.0)
	v1 := utils.NewVector3(1.0, 1.0, 1.0)
	v1.MulF(1.0 - t)
	v2 := utils.NewVector3(0.5, 0.7, 1.0)
	v2.MulF(t)
	v1.Add(v2)
	return v1
}

const (
	w               = 200
	h               = 100
	samplesPerPixel = 100
)

func main() {
	img := utils.NewRGBA(w, h)

	cam := utils.NewCamera()

	world := utils.NewHitTableList()
	world.Add(utils.NewSphere(utils.NewVector3(0.0, 0.0, -1), 0.5))
	world.Add(utils.NewSphere(utils.NewVector3(0.0, -100.5, -1), 100))

	for j := 0; j < h; j++ {
		utils.PrintProgress(j, h)
		for i := 0; i < w; i++ {
			color := utils.NewVector3(0.0, 0.0, 0.0)
			for k := 0; k < samplesPerPixel; k++ {
				u := (float32(i) + utils.Randf()) / float32(w)
				v := (float32(j) + utils.Randf()) / float32(h)
				r := cam.GetRay(u, v)
				color.Add(RayColor(r, world))
			}

			color.MulF(1.0 / samplesPerPixel)
			img.SetRGB(i, h-j-1, color)
		}
	}

	img.SavePNG("result.png")
}
