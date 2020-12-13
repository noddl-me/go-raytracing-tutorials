package utils

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

// RGBAFC the float32 format of RGBA
type RGBAFC struct {
	R, G, B, A float32
}

// RGBA Impleted the color interface
func (c RGBAFC) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

// RGBAf the float format of RGBA
type RGBAf struct {
	// Pix holds the image's pixels, in R, G, B, A order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []float32
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

// ColorModel get the ColorModel of image
func (p *RGBAf) ColorModel() color.Model { return color.RGBAModel }

// Bounds get the Bound of image
func (p *RGBAf) Bounds() image.Rectangle { return p.Rect }

// At get the (x, y) of image
func (p *RGBAf) At(x, y int) color.Color {
	return p.RGBAAt(x, y)
}

// RGBAAt get the (x, y) of image
func (p *RGBAf) RGBAAt(x, y int) RGBAFC {
	if !(image.Point{x, y}.In(p.Rect)) {
		return RGBAFC{}
	}
	i := p.PixOffset(x, y)
	s := p.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	return RGBAFC{s[0], s[1], s[2], s[3]}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGBAf) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*4
}

// Set set the RGBA to Image
func (p *RGBAf) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := color.RGBAModel.Convert(c).(color.RGBA)
	s := p.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	s[0] = float32(c1.R)
	s[1] = float32(c1.G)
	s[2] = float32(c1.B)
	s[3] = float32(c1.A)
}

// SetRGBA set the RGBA to Image
func (p *RGBAf) SetRGBA(x, y int, c RGBAFC) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	s := p.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	s[0] = float32(c.R)
	s[1] = float32(c.G)
	s[2] = float32(c.B)
	s[3] = float32(c.A)
}

// SetRGB set the RGB from Vector3 to Image
func (p *RGBAf) SetRGB(x, y int, c *Vector3) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	s := p.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	s[0] = Clamp(c.e[0]*255.999, 0.0, 255.0)
	s[1] = Clamp(c.e[1]*255.999, 0.0, 255.0)
	s[2] = Clamp(c.e[2]*255.999, 0.0, 255.0)
	s[3] = 255.0
}

// SetRGBAf set the RGB from Vector4 to Image
func (p *RGBAf) SetRGBAf(x, y int, c *Vector4) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	s := p.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	s[0] = Clamp(c.e[0]*255.999, 0.0, 255.0)
	s[1] = Clamp(c.e[1]*255.999, 0.0, 255.0)
	s[2] = Clamp(c.e[2]*255.999, 0.0, 255.0)
	s[3] = Clamp(c.e[3]*255.999, 0.0, 255.0)
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGBAf) SubImage(r image.Rectangle) *RGBAf {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGBAf{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGBAf{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *RGBAf) Opaque() bool {
	if p.Rect.Empty() {
		return true
	}
	i0, i1 := 3, p.Rect.Dx()*4
	for y := p.Rect.Min.Y; y < p.Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 4 {
			if p.Pix[i] != 0xff {
				return false
			}
		}
		i0 += p.Stride
		i1 += p.Stride
	}
	return true
}

// ClearWith fill the image with color c
func (p *RGBAf) ClearWith(c RGBAFC) {
	w, h := p.Rect.Dx(), p.Rect.Dy()
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			p.SetRGBA(i, j, c)
		}
	}
}

func (p *RGBAf) SavePNG(name string) {
	pngFile, _ := os.Create(name)
	png.Encode(pngFile, p)
	pngFile.Close()
}

func (p *RGBAf) SaveJPEG(name string) {
	pngFile, _ := os.Create(name)
	jpeg.Encode(pngFile, p, nil)
	pngFile.Close()
}

// NewRGBAf returns a new RGBAf image with the given bounds.
func NewRGBAf(r image.Rectangle) *RGBAf {
	w, h := r.Dx(), r.Dy()
	buf := make([]float32, 4*w*h)
	return &RGBAf{buf, 4 * w, r}
}

func NewRGBA(w, h int) *RGBAf {
	buf := make([]float32, 4*w*h)
	return &RGBAf{buf, 4 * w, image.Rect(0, 0, w, h)}
}
