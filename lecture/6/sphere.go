package utils

type Sphere struct {
	v   *Vector4
	Mat Material
}

func (s *Sphere) Hit(r *Ray, tMin, tMax float32) *HitRecord {
	center := ToVector3(s.v)
	oc := r.Origin.SubVec(center)
	a := r.Direction.LengthSquare()
	b := oc.Dot(r.Direction)
	c := oc.LengthSquare() - s.v.W()*s.v.W()
	discriminant := b*b - a*c
	if discriminant < 0.0 {
		return nil
	}
	root := Sqrtf(discriminant)
	t := (-b - root) / a
	if t < tMin || tMax < t {
		t = (-b + root) / a
		if t < tMin || tMax < t {
			return nil
		}
	}

	p := r.At(t)
	return NewHitRecordWithRayMat(
		p,
		p.SubVec(center).DivFVec(s.v.W()),
		t,
		r,
		s.Mat,
	)
}

func NewSphere(center *Vector3, radius float32) *Sphere {
	ret := &Sphere{ToVector4(center), nil}
	ret.v.e[3] = radius
	return ret
}

func NewSphereWithMat(center *Vector3, radius float32, mat Material) *Sphere {
	ret := &Sphere{ToVector4(center), mat}
	ret.v.e[3] = radius
	return ret
}