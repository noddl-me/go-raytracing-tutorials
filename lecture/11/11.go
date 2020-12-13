package main

import (
	"fmt"
	"raytracing-tutorials/utils"
	"time"
)

func RayColor(r *utils.Ray, world utils.HitTable, depth int) *utils.Vector3 {
	if depth <= 0 {
		return utils.NewVector3(0.0, 0.0, 0.0)
	}

	if rec := world.Hit(r, 0.001, utils.Inf32(1)); rec != nil {
		if rec.Mat != nil {
			if attenuation, scattered := rec.Mat.Scatter(r, rec); attenuation != nil {
				ret := RayColor(scattered, world, depth-1)
				ret.Mul(attenuation)
				return ret
			}
		}
		return utils.NewVector3(0.0, 0.0, 0.0)
	}

	t := 0.5 * (r.Direction.Unit().Y() + 1.0)
	v := utils.NewVector3(0.5, 0.7, 1.0)
	v.MulF(t)
	v.AddF(1.0 - t)
	return v
}

const (
	w               = 400
	h               = 225
	aspectRatio     = float32(w) / float32(h)
	samplesPerPixel = 100
	maxDepth        = 50
)

func main() {
	t := time.Now()
	img := utils.NewRGBA(w, h)

	cam := utils.NewCameraWithOptions(
		utils.NewVector3(-2.0, 2.0, 1.0),
		utils.NewVector3(0.0, 0.0, -1.0),
		utils.NewVector3(0.0, 1.0, 0.0), 90.0, aspectRatio)

	materialGround := utils.NewLambertian(utils.NewVector3(0.8, 0.8, 0.0))
	materialCenter := utils.NewLambertian(utils.NewVector3(0.1, 0.2, 0.5))
	materialLeft := utils.NewDielectric(1.5)
	materialRight := utils.NewMetal(utils.NewVector3(0.8, 0.6, 0.2), 0.0)

	world := utils.NewHitTableList()
	world.Add(utils.NewSphereWithMat(utils.NewVector3(0.0, -100.5, -1.0), 100.0,
		materialGround))
	world.Add(utils.NewSphereWithMat(utils.NewVector3(0.0, 0.0, -1.0), 0.5,
		materialCenter))
	world.Add(utils.NewSphereWithMat(utils.NewVector3(-1.0, 0.0, -1.0), 0.5,
		materialLeft))
	world.Add(utils.NewSphereWithMat(utils.NewVector3(-1.0, 0.0, -1.0), -0.45,
		materialLeft))
	world.Add(utils.NewSphereWithMat(utils.NewVector3(1.0, 0.0, -1.0), 0.5,
		materialRight))

	for j := 0; j < h; j++ {
		utils.PrintProgress(j, h)
		for i := 0; i < w; i++ {
			color := utils.NewVector3(0.0, 0.0, 0.0)
			for k := 0; k < samplesPerPixel; k++ {
				u := (float32(i) + utils.Randf()) / float32(w)
				v := (float32(j) + utils.Randf()) / float32(h)
				r := cam.GetRay(u, v)
				color.Add(RayColor(r, world, maxDepth))
			}

			color.MulF(1.0 / samplesPerPixel)
			color.Sqrt()
			img.SetRGB(i, h-j-1, color)
		}
	}

	img.SavePNG("result.png")
	fmt.Println(time.Since(t))
}
