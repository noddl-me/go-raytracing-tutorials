package main

import (
	"fmt"
	"raytracing-tutorials/utils"
	"sync"
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

func RandomScene() *utils.HitTableList {
	world := utils.NewHitTableList()

	groundMaterial := utils.NewLambertian(utils.NewVector3(0.5, 0.5, 0.5))
	world.Add(utils.NewSphereWithMat(utils.NewVector3(0, -1000, 0), 1000, groundMaterial))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := utils.Randf()
			center := utils.NewVector3(float32(a)+0.9*utils.Randf(), 0.2, float32(b)+0.9*utils.Randf())

			if (center.SubVec(utils.NewVector3(4, 0.2, 0))).Length() > 0.9 {
				var sphereMaterial utils.Material

				if chooseMat < 0.8 {
					// diffuse
					albedo := utils.NewRandomVector3()
					albedo.Mul(utils.NewRandomVector3())
					sphereMaterial = utils.NewLambertian(albedo)
				} else if chooseMat < 0.95 {
					// metal
					albedo := utils.NewRandomVector3MinMax(0.5, 1)
					fuzz := utils.RandfMinMax(0, 0.5)
					sphereMaterial = utils.NewMetal(albedo, fuzz)
				} else {
					// glass
					sphereMaterial = utils.NewDielectric(1.5)
				}
				world.Add(utils.NewSphereWithMat(center, 0.2, sphereMaterial))
			}
		}
	}

	material1 := utils.NewDielectric(1.5)
	world.Add(utils.NewSphereWithMat(utils.NewVector3(0, 1, 0), 1.0, material1))

	material2 := utils.NewLambertian(utils.NewVector3(0.4, 0.2, 0.1))
	world.Add(utils.NewSphereWithMat(utils.NewVector3(-4, 1, 0), 1.0, material2))

	material3 := utils.NewMetal(utils.NewVector3(0.7, 0.6, 0.5), 0.0)
	world.Add(utils.NewSphereWithMat(utils.NewVector3(4, 1, 0), 1.0, material3))

	return world
}

const (
	w               = 1200
	h               = 800
	aspectRatio     = float32(w) / float32(h)
	samplesPerPixel = 500
	maxDepth        = 50
)

func main() {
	t := time.Now()
	img := utils.NewRGBA(w, h)

	lookFrom := utils.NewVector3(13.0, 2.0, 3.0)
	lookAt := utils.NewVector3(0.0, 0.0, 0.0)
	cam := utils.NewCameraWithFocusOptions(
		lookFrom, lookAt, utils.NewVector3(0.0, 1.0, 0.0), 20.0, aspectRatio,
		0.1, 10.0)

	world := RandomScene()

	for j := 0; j < h; j++ {
		utils.PrintProgress(j, h)
		wg := sync.WaitGroup{}
		for i := 0; i < w; i++ {
			wg.Add(1)
			go func(i, j int) {
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
				wg.Done()
			}(i, j)
		}
		wg.Wait()
	}

	img.SavePNG("result.png")
	fmt.Println(time.Since(t))
}
