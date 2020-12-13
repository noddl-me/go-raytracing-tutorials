package utils

import "raytracing-tutorials/math32"

type Material interface {
	Scatter(r_in *Ray, rec *HitRecord) (*Vector3, *Ray)
}

type Lambertian struct {
	Albedo *Vector3
}

func (l *Lambertian) Scatter(r_in *Ray, rec *HitRecord) (*Vector3, *Ray) {
	dir := rec.Normal.AddVec(NewRandomUnitVector3())
	scattered := NewRay(rec.P, dir)
	attenuation := CopyVector3(l.Albedo)
	return attenuation, scattered
}

func NewLambertian(a *Vector3) *Lambertian {
	return &Lambertian{a}
}

type Metal struct {
	Albedo *Vector3
	Fuzz   float32
}

func (m *Metal) Scatter(r_in *Ray, rec *HitRecord) (*Vector3, *Ray) {
	reflected := r_in.Direction.Unit().Reflect(rec.Normal)
	randV := NewRandomVector3InUnitSphere()
	randV.MulF(m.Fuzz)
	reflected.Add(randV)
	scattered := NewRay(rec.P, reflected)
	attenuation := CopyVector3(m.Albedo)
	if scattered.Direction.Dot(rec.Normal) <= 0.0 {
		return nil, nil
	}
	return attenuation, scattered
}

func NewMetal(a *Vector3, fuzz float32) *Metal {
	if fuzz > 1.0 {
		fuzz = 1.0
	}
	return &Metal{a, fuzz}
}

type Dielectric struct {
	RefIdx float32
}

func reflectance(cosine, refIdx float32) float32 {
	// Use Schlick's approximation for reflectance.
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math32.Pow((1-cosine), 5.0)
}

func (d *Dielectric) Scatter(r_in *Ray, rec *HitRecord) (*Vector3, *Ray) {
	attenuation := NewVector3(1.0, 1.0, 1.0)
	refractionRatio := d.RefIdx
	if rec.FrontFace {
		refractionRatio = 1.0 / refractionRatio
	}
	unitDirection := r_in.Direction.Unit()
	direction := unitDirection.Refract(rec.Normal, refractionRatio)
	// cosTheta := Minf(unitDirection.MulFVec(-1.0).Dot(rec.Normal), 1.0)
	// sinTheta := Sqrtf(1.0 - cosTheta*cosTheta)

	// cannotRefract := refractionRatio*sinTheta > 1.0
	// direction := unitDirection.Reflect(rec.Normal)
	// if !cannotRefract && reflectance(cosTheta, refractionRatio) < Randf() {
	// 	direction = unitDirection.Refract(rec.Normal, refractionRatio)
	// }
	scattered := NewRay(rec.P, direction)
	return attenuation, scattered
}

func NewDielectric(ri float32) *Dielectric {
	return &Dielectric{ri}
}
