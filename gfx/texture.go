package gfx

import (
	"errors"
	"image"
	"image/draw"
	_ "image/jpeg" // Import JPEG Decoding
	_ "image/png"  // Import PNG Decoding

	"github.com/mcbernie/myopengl/glHelper"
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
	glHelper.GenTextures(1, &handle)

	target := uint32(glHelper.GlTexture2D)

	texture := Texture{
		handle: handle,
		target: target,
	}
	return &texture
}

//NewTextureFromFile Loads an Texture from Image file
func NewTextureFromFile(path string) *Texture {

	img, err := LoadImageFromFile(path)

	if err != nil {
		panic("error on loading texture from file!")
	}

	var handle uint32
	glHelper.GenTextures(1, &handle)
	target := uint32(glHelper.GlTexture2D)

	texture := Texture{
		handle: handle,
		target: target,
	}

	texture.SetImage(img, glHelper.GlClampToEdge, glHelper.GlClampToEdge)

	return &texture
}

//GetHandle retruns texture Handle
func (tex *Texture) GetHandle() uint32 {
	return tex.handle
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

	internalFmt := int32(glHelper.GlSrgbAlpha)
	format := uint32(glHelper.GlRgbA)
	width := int32(rgba.Rect.Size().X)
	height := int32(rgba.Rect.Size().Y)
	pixType := uint32(glHelper.GlUnsignedByte)
	dataPtr := glHelper.Ptr(rgba.Pix)

	tex.width = width
	tex.height = height

	tex.Bind(tex.unit)
	defer tex.UnBind()

	glHelper.TexParameteri(tex.target, glHelper.GlTextureWrapR, wrapR)
	glHelper.TexParameteri(tex.target, glHelper.GlTextureWrapS, wrapS)
	glHelper.TexParameteri(tex.target, glHelper.GlTextureMinFilter, glHelper.GlLinear)
	glHelper.TexParameteri(tex.target, glHelper.GlTextureMagFilter, glHelper.GlLinear)

	glHelper.TexImage2D(tex.target, 0, internalFmt, width, height, 0, format, pixType, dataPtr)

	glHelper.TexParameteri(tex.target, glHelper.GlTextureMinFilter, glHelper.GlLinearMipmapLinear)
	glHelper.TexParameteri(tex.target, glHelper.GlTextureLodBias, 0)
	glHelper.GenerateMipmap(tex.target)

	return nil
}

//Bind binds a Texture to OpenGl
func (tex *Texture) Bind(unit uint32) {
	glHelper.ActiveTexture(glHelper.GlTexture0 + unit)
	glHelper.BindTexture(tex.target, tex.handle)

	tex.unit = glHelper.GlTexture0 + unit
}

//UnBind remove a Texture from OpenGL
func (tex *Texture) UnBind() {
	tex.unit = 0
	//glHelper.ActiveTexture(0)
	//glHelper.BindTexture(tex.target, 0)
}

//Delete remove a texture from Memory
func (tex *Texture) Delete() {
	glHelper.DeleteTextures(1, &tex.handle)
}

//SetUniform sets the uniform Variable in OpenGL
func (tex *Texture) SetUniform(uniformLoc int32) error {
	if tex.texUnit == 0 {
		return errTextureNotBound
	}

	glHelper.Uniform1i(uniformLoc, int32(tex.texUnit-glHelper.GlTexture0))
	return nil
}
