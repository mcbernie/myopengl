package objects

import (
	"errors"
	"image"
	"image/draw"
	_ "image/jpeg" // Import JPEG Decoding
	_ "image/png"  // Import PNG Decoding

	"github.com/mcbernie/myopengl/graphic/helper"
)

//Texture struct for a Texture
type Texture struct {
	handle  uint32
	target  uint32 // same target as gl.BindTexture(<this param>, ...)
	texUnit uint32 // Texture unit that is currently bound to ex: gl.TEXTURE0
	unit    uint32
	Width   int32
	Height  int32
}

var errUnsupportedStride = errors.New("unsupported stride, only 32-bit colors supported")
var errTextureNotBound = errors.New("texture not found")

//NewTextureDefault initialize a new Texture with default WrapS and WrapR from Image
func NewTextureDefault() *Texture {
	return NewTexture(helper.GlLinear, helper.GlLinear)
}

//NewTexture initialize a new Texture from Image
func NewTexture(wrapR, wrapS int32) *Texture {

	var handle uint32
	helper.GenTextures(1, &handle)

	target := uint32(helper.GlTexture2D)

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
	helper.GenTextures(1, &handle)
	target := uint32(helper.GlTexture2D)

	texture := Texture{
		handle: handle,
		target: target,
	}

	texture.SetImage(img, helper.GlClampToEdge, helper.GlClampToEdge)

	return &texture
}

//GetHandle retruns texture Handle
func (tex *Texture) GetHandle() uint32 {
	return tex.handle
}

//SetDefaultImage create wrapR and wrapS by default
func (tex *Texture) SetDefaultImage(img image.Image) error {
	return tex.SetImage(img, helper.GlClampToEdge, helper.GlClampToEdge)
}

func (tex *Texture) ReplaceImage(img image.Image) error {
	if img == nil {
		return errUnsupportedStride
	}
	rgba := image.NewRGBA(img.Bounds())

	draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return errUnsupportedStride
	}
	internalFmt := int32(helper.GlSrgbAlpha)
	format := uint32(helper.GlRgbA)
	width := int32(rgba.Rect.Size().X)
	height := int32(rgba.Rect.Size().Y)
	pixType := uint32(helper.GlUnsignedByte)
	dataPtr := helper.Ptr(rgba.Pix)

	tex.Width = width
	tex.Height = height
	tex.Bind(tex.unit)
	helper.TexImage2D(tex.target, 0, internalFmt, width, height, 0, format, pixType, dataPtr)

	helper.TexParameteri(tex.target, helper.GlTextureMinFilter, helper.GlLinear)
	helper.TexParameteri(tex.target, helper.GlTextureMagFilter, helper.GlLinear)
	helper.TexParameteri(tex.target, helper.GlTextureMinFilter, helper.GlLinearMipmapLinear)
	//helper.TexParameteri(tex.target, helper.GlTextureLodBias, 0)
	tex.UnBind()
	//log.Println(helper.ErrorCheck())
	return nil

}

//SetImage is for setting or replacing a image
func (tex *Texture) SetImage(img image.Image, wrapR, wrapS int32) error {

	if err := tex.ReplaceImage(img); err != nil {
		return err
	}

	tex.Bind(tex.unit)

	helper.TexParameteri(tex.target, helper.GlTextureWrapR, wrapR)
	helper.TexParameteri(tex.target, helper.GlTextureWrapS, wrapS)
	helper.GenerateMipmap(tex.target)
	tex.UnBind()
	return nil
}

//Bind binds a Texture to OpenGl
func (tex *Texture) Bind(unit uint32) {
	helper.ActiveTexture(helper.GlTexture0 + unit)
	helper.BindTexture(tex.target, tex.handle)

	tex.unit = helper.GlTexture0 + unit
}

//UnBind remove a Texture from OpenGL
func (tex *Texture) UnBind() {
	tex.unit = 0
}

//Delete remove a texture from Memory
func (tex *Texture) Delete() {
	helper.DeleteTextures(1, &tex.handle)
}

//SetUniform sets the uniform Variable in OpenGL
func (tex *Texture) SetUniform(uniformLoc int32) error {
	if tex.texUnit == 0 {
		return errTextureNotBound
	}

	helper.Uniform1i(uniformLoc, int32(tex.texUnit-helper.GlTexture0))
	return nil
}
