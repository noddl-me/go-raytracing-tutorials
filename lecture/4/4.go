package main

import (
	"raytracing-tutorials/utils"
)

func HitSphere(ball *utils.Vector4, r *utils.Ray) bool {
	center := utils.ToVector3(ball)
	oc := r.Origin.SubVec(center)
	a := r.Direction.LengthSquare()
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.LengthSquare() - ball.W()*ball.W()
	discriminant := b*b - 4*a*c
	return (discriminant > 0)
}

func RayColor(r *utils.Ray) *utils.Vector3 {
	if HitSphere(utils.NewVector4(0, 0, -1, 0.5), r) {
		return utils.NewVector3(1, 0, 0)
	}
	t := (r.Direction.Unit().Y() + 1.0) * 0.5
	v := utils.NewVector3(0.5, 0.7, 1.0)
	v.MulF(t)
	v.AddF(1.0 - t)
	return v
}

func main() {
	w, h := 200, 100
	img := utils.NewRGBA(w, h)

	lower_left_corner := utils.NewVector3(-2.0, -1.0, -1.0)
	horizontal := utils.NewVector3(4.0, 0.0, 0.0)
	vertical := utils.NewVector3(0.0, 2.0, 0.0)
	origin := utils.NewVector3(0.0, 0.0, 0.0)

	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			u := float32(i) / float32(w)
			v := float32(j) / float32(h)
			dir := lower_left_corner.AddVec(horizontal.MulFVec(u))
			dir.Add(vertical.MulFVec(v))
			r := utils.NewRay(origin, dir)
			color := RayColor(r)
			img.SetRGB(i, h-j-1, color)
		}
	}

	img.SavePNG("result.png")
}
