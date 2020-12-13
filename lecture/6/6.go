package main

import (
	"raytracing-tutorials/utils"
)

func RayColor(r *utils.Ray, world utils.HitTable) *utils.Vector3 {
	if rec := world.Hit(r, 0, utils.Inf32(1)); rec != nil {
		rec.Normal.AddF(1.0)
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

func main() {
	w, h := 200, 100
	img := utils.NewRGBA(w, h)

	lower_left_corner := utils.NewVector3(-2.0, -1.0, -1.0)
	horizontal := utils.NewVector3(4.0, 0.0, 0.0)
	vertical := utils.NewVector3(0.0, 2.0, 0.0)
	origin := utils.NewVector3(0.0, 0.0, 0.0)

	world := utils.NewHitTableList()
	world.Add(utils.NewSphere(utils.NewVector3(0.0, 0.0, -1), 0.5))
	world.Add(utils.NewSphere(utils.NewVector3(0.0, -100.5, -1), 100))

	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			u := float32(i) / float32(w)
			v := float32(j) / float32(h)
			dir := lower_left_corner.AddVec(horizontal.MulFVec(u))
			dir.Add(vertical.MulFVec(v))
			r := utils.NewRay(origin, dir)
			color := RayColor(r, world)
			img.SetRGB(i, h-j-1, color)
		}
	}

	img.SavePNG("result.png")
}
