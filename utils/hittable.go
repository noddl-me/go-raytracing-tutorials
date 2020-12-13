package utils

type HitRecord struct {
	P, Normal *Vector3
	T         float32
	FrontFace bool
	Mat       Material
}

func (hr *HitRecord) SetFaceNormal(r *Ray) {
	hr.FrontFace = r.Direction.Dot(hr.Normal) < 0.0
	if !hr.FrontFace {
		hr.Normal.MulF(-1.0)
	}
}

func NewHitRecord(p, normal *Vector3, t float32, frontFace bool) *HitRecord {
	return &HitRecord{p, normal, t, frontFace, nil}
}

func NewHitRecordWithRay(p, normal *Vector3, t float32, r *Ray) *HitRecord {
	ret := &HitRecord{p, normal, t, false, nil}
	ret.SetFaceNormal(r)
	return ret
}

func NewHitRecordWithRayMat(p, normal *Vector3, t float32, r *Ray, mat Material) *HitRecord {
	ret := &HitRecord{p, normal, t, false, mat}
	ret.SetFaceNormal(r)
	return ret
}

type HitTable interface {
	Hit(r *Ray, tMin, tMax float32) *HitRecord
}

type HitTableList struct {
	Objects []HitTable
}

func (htl *HitTableList) Hit(r *Ray, tMin, tMax float32) *HitRecord {
	var ret *HitRecord = nil
	closestSoFar := tMax
	for i := range htl.Objects {
		hr := htl.Objects[i].Hit(r, tMin, closestSoFar)
		if hr != nil {
			ret = hr
			closestSoFar = hr.T
		}
	}
	return ret
}

func (htl *HitTableList) Add(obj HitTable) {
	htl.Objects = append(htl.Objects, obj)
}

func NewHitTableList() *HitTableList {
	return &HitTableList{make([]HitTable, 0)}
}
