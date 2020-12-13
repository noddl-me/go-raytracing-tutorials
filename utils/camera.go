package utils

import "raytracing-tutorials/math32"

type Camera struct {
	Origin          *Vector3
	LowerLeftCorner *Vector3
	Horizontal      *Vector3
	Vertical        *Vector3
	u, v, w         *Vector3
	lenRadius       float32
}

func (c *Camera) GetRay(u, v float32) *Ray {
	dir := c.LowerLeftCorner.AddVec(c.Horizontal.MulFVec(u))
	dir.Add(c.Vertical.MulFVec(v))
	dir.Sub(c.Origin)
	if c.u == nil {
		return NewRay(c.Origin, dir)
	}
	rd := NewRandomVector3InUnitDisk()
	rd.MulF(c.lenRadius)
	offset := c.u.MulFVec(rd.X()).AddVec(c.v.MulFVec(rd.Y()))
	dir.Sub(offset)
	return NewRay(c.Origin.AddVec(offset), dir)
}

func NewCamera() *Camera {
	return &Camera{
		NewVector3(0.0, 0.0, 0.0),
		NewVector3(-2.0, -1.0, -1.0),
		NewVector3(4.0, 0.0, 0.0),
		NewVector3(0.0, 2.0, 0.0),
		nil, nil, nil, 0.0,
	}
}

func NewCameraWithOptions(center, lookAt, vUp *Vector3, vFov, aspectRatio float32) *Camera {
	theta := ToRadians(vFov)
	h := math32.TanGo(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := center.SubVec(lookAt).Unit()
	u := vUp.Cross(w).Unit()
	v := w.Cross(u)
	u.MulF(viewportWidth)
	v.MulF(viewportHeight)

	ret := &Camera{
		center,
		center.SubVec(w),
		u,
		v,
		nil, nil, nil, 0.0,
	}
	ret.LowerLeftCorner.Add(ret.Horizontal.DivFVec(-2.0))
	ret.LowerLeftCorner.Add(ret.Vertical.DivFVec(-2.0))
	return ret
}

func NewCameraWithFocusOptions(center, lookAt, vUp *Vector3,
	vFov, aspectRatio, aperture, focusDist float32) *Camera {
	theta := ToRadians(vFov)
	h := math32.TanGo(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := center.SubVec(lookAt).Unit()
	u := vUp.Cross(w).Unit()
	v := w.Cross(u)

	ret := &Camera{
		center,
		center.SubVec(w.MulFVec(focusDist)),
		u.MulFVec(focusDist * viewportWidth),
		v.MulFVec(focusDist * viewportHeight),
		u, v, w, aperture / 2.0,
	}
	ret.LowerLeftCorner.Add(ret.Horizontal.DivFVec(-2.0))
	ret.LowerLeftCorner.Add(ret.Vertical.DivFVec(-2.0))
	return ret
}
