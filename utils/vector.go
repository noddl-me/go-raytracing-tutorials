package utils

import "raytracing-tutorials/math32"

type Vector4 struct {
	e [4]float32
}

func (v *Vector4) X() float32 {
	return v.e[0]
}

func (v *Vector4) Y() float32 {
	return v.e[1]
}

func (v *Vector4) Z() float32 {
	return v.e[2]
}

func (v *Vector4) W() float32 {
	return v.e[3]
}

func (v1 *Vector4) Add(v2 *Vector4) {
	v1.e[0] += v2.e[0]
	v1.e[1] += v2.e[1]
	v1.e[2] += v2.e[2]
	v1.e[3] += v2.e[3]
}

func (v1 *Vector4) AddVec(v2 *Vector4) *Vector4 {
	return NewVector4(
		v1.e[0]+v2.e[0],
		v1.e[1]+v2.e[1],
		v1.e[2]+v2.e[2],
		v1.e[3]+v2.e[3],
	)
}

func (v1 *Vector4) AddF(x float32) {
	v1.e[0] += x
	v1.e[1] += x
	v1.e[2] += x
	v1.e[3] += x
}

func (v1 *Vector4) AddFVec(x float32) *Vector4 {
	return NewVector4(
		v1.e[0]+x,
		v1.e[1]+x,
		v1.e[2]+x,
		v1.e[3]+x,
	)
}

func (v1 *Vector4) Sub(v2 *Vector4) {
	v1.e[0] -= v2.e[0]
	v1.e[1] -= v2.e[1]
	v1.e[2] -= v2.e[2]
	v1.e[3] -= v2.e[3]
}

func (v1 *Vector4) SubVec(v2 *Vector4) *Vector4 {
	return NewVector4(
		v1.e[0]-v2.e[0],
		v1.e[1]-v2.e[1],
		v1.e[2]-v2.e[2],
		v1.e[3]-v2.e[3],
	)
}

func (v1 *Vector4) SubF(x float32) {
	v1.e[0] -= x
	v1.e[1] -= x
	v1.e[2] -= x
	v1.e[3] -= x
}

func (v1 *Vector4) SubFVec(x float32) *Vector4 {
	return NewVector4(
		v1.e[0]-x,
		v1.e[1]-x,
		v1.e[2]-x,
		v1.e[3]-x,
	)
}

func (v1 *Vector4) Mul(v2 *Vector4) {
	v1.e[0] *= v2.e[0]
	v1.e[1] *= v2.e[1]
	v1.e[2] *= v2.e[2]
	v1.e[3] *= v2.e[3]
}

func (v1 *Vector4) MulVec(v2 *Vector4) *Vector4 {
	return NewVector4(
		v1.e[0]*v2.e[0],
		v1.e[1]*v2.e[1],
		v1.e[2]*v2.e[2],
		v1.e[3]*v2.e[3],
	)
}

func (v1 *Vector4) MulF(x float32) {
	v1.e[0] *= x
	v1.e[1] *= x
	v1.e[2] *= x
	v1.e[3] *= x
}

func (v1 *Vector4) MulFVec(x float32) *Vector4 {
	return NewVector4(
		v1.e[0]*x,
		v1.e[1]*x,
		v1.e[2]*x,
		v1.e[3]*x,
	)
}

func (v1 *Vector4) Div(v2 *Vector4) {
	v1.e[0] /= v2.e[0]
	v1.e[1] /= v2.e[1]
	v1.e[2] /= v2.e[2]
	v1.e[3] /= v2.e[3]
}

func (v1 *Vector4) DivVec(v2 *Vector4) *Vector4 {
	return NewVector4(
		v1.e[0]/v2.e[0],
		v1.e[1]/v2.e[1],
		v1.e[2]/v2.e[2],
		v1.e[3]/v2.e[3],
	)
}

func (v1 *Vector4) DivF(x float32) {
	v1.e[0] /= x
	v1.e[1] /= x
	v1.e[2] /= x
	v1.e[3] /= x
}

func (v1 *Vector4) DivFVec(x float32) *Vector4 {
	return NewVector4(
		v1.e[0]/x,
		v1.e[1]/x,
		v1.e[2]/x,
		v1.e[3]/x,
	)
}

func (v1 *Vector4) Dot(v2 *Vector4) float32 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2] + v1.e[3]*v2.e[3]
}

func (v *Vector4) LengthSquare() float32 {
	return v.Dot(v)
}

func (v *Vector4) Length() float32 {
	return Sqrtf(v.LengthSquare())
}

func (v *Vector4) Unit() *Vector4 {
	return v.DivFVec(v.Length())
}

func NewVector4(x, y, z, w float32) *Vector4 {
	return &Vector4{[4]float32{x, y, z, w}}
}

func CopyVector4(v *Vector4) *Vector4 {
	return &Vector4{[4]float32{v.e[0], v.e[1], v.e[2], v.e[3]}}
}

type Vector3 Vector4

func (v *Vector3) X() float32 {
	return v.e[0]
}

func (v *Vector3) Y() float32 {
	return v.e[1]
}

func (v *Vector3) Z() float32 {
	return v.e[2]
}

func (v1 *Vector3) Add(v2 *Vector3) {
	v1.e[0] += v2.e[0]
	v1.e[1] += v2.e[1]
	v1.e[2] += v2.e[2]
}

func (v1 *Vector3) AddVec(v2 *Vector3) *Vector3 {
	return NewVector3(
		v1.e[0]+v2.e[0],
		v1.e[1]+v2.e[1],
		v1.e[2]+v2.e[2],
	)
}

func (v1 *Vector3) AddF(x float32) {
	v1.e[0] += x
	v1.e[1] += x
	v1.e[2] += x
}

func (v1 *Vector3) AddFVec(x float32) *Vector3 {
	return NewVector3(
		v1.e[0]+x,
		v1.e[1]+x,
		v1.e[2]+x,
	)
}

func (v1 *Vector3) Sub(v2 *Vector3) {
	v1.e[0] -= v2.e[0]
	v1.e[1] -= v2.e[1]
	v1.e[2] -= v2.e[2]
}

func (v1 *Vector3) SubVec(v2 *Vector3) *Vector3 {
	return NewVector3(
		v1.e[0]-v2.e[0],
		v1.e[1]-v2.e[1],
		v1.e[2]-v2.e[2],
	)
}

func (v1 *Vector3) SubF(x float32) {
	v1.e[0] -= x
	v1.e[1] -= x
	v1.e[2] -= x
}

func (v1 *Vector3) SubFVec(x float32) *Vector3 {
	return NewVector3(
		v1.e[0]-x,
		v1.e[1]-x,
		v1.e[2]-x,
	)
}

func (v1 *Vector3) Mul(v2 *Vector3) {
	v1.e[0] *= v2.e[0]
	v1.e[1] *= v2.e[1]
	v1.e[2] *= v2.e[2]
}

func (v1 *Vector3) MulVec(v2 *Vector3) *Vector3 {
	return NewVector3(
		v1.e[0]*v2.e[0],
		v1.e[1]*v2.e[1],
		v1.e[2]*v2.e[2],
	)
}

func (v1 *Vector3) MulF(x float32) {
	v1.e[0] *= x
	v1.e[1] *= x
	v1.e[2] *= x
}

func (v1 *Vector3) MulFVec(x float32) *Vector3 {
	return NewVector3(
		v1.e[0]*x,
		v1.e[1]*x,
		v1.e[2]*x,
	)
}

func (v1 *Vector3) Div(v2 *Vector3) {
	v1.e[0] /= v2.e[0]
	v1.e[1] /= v2.e[1]
	v1.e[2] /= v2.e[2]
}

func (v1 *Vector3) DivVec(v2 *Vector3) *Vector3 {
	return NewVector3(
		v1.e[0]/v2.e[0],
		v1.e[1]/v2.e[1],
		v1.e[2]/v2.e[2],
	)
}

func (v1 *Vector3) DivF(x float32) {
	v1.e[0] /= x
	v1.e[1] /= x
	v1.e[2] /= x
}

func (v1 *Vector3) DivFVec(x float32) *Vector3 {
	return NewVector3(
		v1.e[0]/x,
		v1.e[1]/x,
		v1.e[2]/x,
	)
}

func (v1 *Vector3) Sqrt() {
	v1.e[0] = Sqrtf(v1.e[0])
	v1.e[1] = Sqrtf(v1.e[1])
	v1.e[2] = Sqrtf(v1.e[2])
}

func (v1 *Vector3) SqrtVec() *Vector3 {
	return NewVector3(
		Sqrtf(v1.e[0]),
		Sqrtf(v1.e[1]),
		Sqrtf(v1.e[2]),
	)
}

func (v1 *Vector3) Dot(v2 *Vector3) float32 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2]
}

func (v *Vector3) LengthSquare() float32 {
	return v.Dot(v)
}

func (v *Vector3) Length() float32 {
	return Sqrtf(v.LengthSquare())
}

func (v1 *Vector3) Cross(v2 *Vector3) *Vector3 {
	return NewVector3(
		v1.e[1]*v2.e[2]-v1.e[2]*v2.e[1],
		v1.e[2]*v2.e[0]-v1.e[0]*v2.e[2],
		v1.e[0]*v2.e[1]-v1.e[1]*v2.e[0],
	)
}

func (v *Vector3) Unit() *Vector3 {
	return v.DivFVec(v.Length())
}

func (v *Vector3) Reflect(n *Vector3) *Vector3 {
	n1 := n.MulFVec(2.0 * v.Dot(n))
	return v.SubVec(n1)
}

func (uv *Vector3) Refract(n *Vector3, etaiOverEtat float32) *Vector3 {
	cosTheta := Minf(uv.MulFVec(-1.0).Dot(n), 1.0)
	rOutParallel := uv.AddVec(n.MulFVec(cosTheta))
	rOutParallel.MulF(etaiOverEtat)
	rOutPerp := CopyVector3(n)
	rOutPerp.MulF(-Sqrtf(1.0 - rOutParallel.LengthSquare()))
	rOutParallel.Add(rOutPerp)
	return rOutParallel
}

func NewVector3(x, y, z float32) *Vector3 {
	return &Vector3{[4]float32{x, y, z}}
}

func CopyVector3(v *Vector3) *Vector3 {
	return &Vector3{[4]float32{v.e[0], v.e[1], v.e[2]}}
}

func NewRandomVector3() *Vector3 {
	return NewVector3(Randf(), Randf(), Randf())
}

func NewRandomVector4() *Vector4 {
	return NewVector4(Randf(), Randf(), Randf(), Randf())
}

func NewRandomVector3MinMax(min, max float32) *Vector3 {
	return NewVector3(RandfMinMax(min, max), RandfMinMax(min, max), RandfMinMax(min, max))
}

func NewRandomVector4MinMax(min, max float32) *Vector4 {
	return NewVector4(RandfMinMax(min, max), RandfMinMax(min, max), RandfMinMax(min, max), RandfMinMax(min, max))
}

func NewRandomVector3InUnitSphere() *Vector3 {
	for true {
		p := NewRandomVector3MinMax(-1, 1)
		if p.LengthSquare() >= 1 {
			continue
		}
		return p
	}
	return nil
}

func NewRandomUnitVector3() *Vector3 {
	a := RandfMinMax(0, 2*PI)
	z := RandfMinMax(-1, 1)
	r := Sqrtf(1 - z*z)
	return NewVector3(r*math32.Cos(a), r*math32.Sin(a), z)
}

func NewRandomVector3InHemisphere(normal *Vector3) *Vector3 {
	inUnitSphere := NewRandomVector3InUnitSphere()
	if inUnitSphere.Dot(normal) <= 0.0 {
		inUnitSphere.MulF(-1.0)
	}
	return inUnitSphere
}

func NewRandomVector3InUnitDisk() *Vector3 {
	for true {
		p := NewVector3(RandfMinMax(-1.0, 1.0), RandfMinMax(-1.0, 1.0), 0)
		if p.LengthSquare() >= 1.0 {
			continue
		}
		return p
	}
	return nil
}

func ToVector3(v *Vector4) *Vector3 {
	return (*Vector3)(v)
}

func ToVector4(v *Vector3) *Vector4 {
	return (*Vector4)(v)
}
