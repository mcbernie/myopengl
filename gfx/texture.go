package gfx

import (
	"errors"
	"image"
	"image/draw"
	_ "image/jpeg" // Import JPEG Decoding
	_ "image/png"  // Import PNG Decoding
	"log"

	//"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/gl/v2.1/gl"
)

//Texture struct for a Texture
type Texture struct {
	handle  uint32
	target  uint32 // same target as gl.BindTexture(<this param>, ...)
	texUnit uint32 // Texture unit that is currently bound to ex: gl.TEXTURE0
	unit    uint32
	width   int32
	height  int32
}

var errUnsupportedStride = errors.New("unsupported stride, only 32-bit colors supported")
var errTextureNotBound = errors.New("texture not found")

//NewTexture initialize a new Texture from Image
func NewTexture(wrapR, wrapS int32) *Texture {

	var handle uint32
	log.Println("generate texture...")
	gl.GenTextures(1, &handle)

	log.Println("texture generated...")

	target := uint32(gl.TEXTURE_2D)

	texture := Texture{
		handle: handle,
		target: target,
	}
	return &texture
}

//SetImage is for setting or replacing a image
func (tex *Texture) SetImage(img image.Image, wrapR, wrapS int32) error {
	if img == nil {
		return errUnsupportedStride
	}

	rgba := image.NewRGBA(img.Bounds())

	draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return errUnsupportedStride
	}

	internalFmt := int32(gl.SRGB_ALPHA)
	format := uint32(gl.RGBA)
	width := int32(rgba.Rect.Size().X)
	height := int32(rgba.Rect.Size().Y)
	pixType := uint32(gl.UNSIGNED_BYTE)
	dataPtr := gl.Ptr(rgba.Pix)

	tex.width = width
	tex.height = height

	tex.Bind(gl.TEXTURE0)
	defer tex.UnBind()

	gl.TexParameteri(tex.target, gl.TEXTURE_WRAP_R, wrapR)
	gl.TexParameteri(tex.target, gl.TEXTURE_WRAP_S, wrapS)
	gl.TexParameteri(tex.target, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(tex.target, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(tex.target, 0, internalFmt, width, height, 0, format, pixType, dataPtr)

	gl.GenerateMipmap(tex.handle)
	return nil
}

//Bind binds a Texture to OpenGl
func (tex *Texture) Bind(unit uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + unit)
	gl.BindTexture(tex.target, tex.handle)
	tex.unit = gl.TEXTURE0 + unit
}

//UnBind remove a Texture from OpenGL
func (tex *Texture) UnBind() {
	tex.unit = 0
	gl.BindTexture(tex.target, 0)
}

//Delete remove a texture from Memory
func (tex *Texture) Delete() {
	gl.DeleteTextures(1, &tex.target)
}

//SetUniform sets the uniform Variable in OpenGL
func (tex *Texture) SetUniform(uniformLoc int32) error {
	if tex.texUnit == 0 {
		return errTextureNotBound
	}

	gl.Uniform1i(uniformLoc, int32(tex.texUnit-gl.TEXTURE0))
	return nil
}

//GetHandle returns own texture handle
func (tex *Texture) GetHandle() uint32 {
	return tex.handle
}
