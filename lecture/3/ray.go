package utils

type Ray struct {
	Origin, Direction *Vector3
}

func (r *Ray) At(t float32) *Vector3 {
	return r.Origin.AddVec(r.Direction.MulFVec(t))
}

func NewRay(ori, dir *Vector3) *Ray {
	return &Ray{ori, dir}
}
