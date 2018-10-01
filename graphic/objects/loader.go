package objects

import (

	//"github.com/go-gl/gl/v4.1-core/gl" // OR:

	"log"

	"github.com/mcbernie/myopengl/graphic/helper"
)

//Loader holds all vaos and vbos handler for cleanup and access
// Helper Generating VAO by calling of LoadToVAO
type Loader struct {
	vaos     []uint32
	vbos     []uint32
	textures []*TextureCleaner
}

//MakeLoader Creat
func MakeLoader() *Loader {
	return &Loader{}
}

//LoadVertexAndTextureToVAO Creates a VAO with position and textureCoord Data and returns a uint32 pointer to VAO
func (l *Loader) LoadVertexAndTextureToVAO(positions []float32, texCoords []float32) uint32 {
	vao := l.createVAO()
	l.storeDataInAttributeList(0, 2, positions)
	l.storeDataInAttributeList(1, 2, texCoords)
	l.unbindVAO()

	return vao
}

//LoadToVAO returns an RawModel and creates an VAO and save handler in Loader struct
func (l *Loader) LoadToVAO(positions []float32, indicies []int32) *RawModel {
	vao := l.createVAO()
	l.bindIndiciesBuffer(indicies)
	l.storeDataInAttributeList(0, 3, positions)
	l.unbindVAO()

	return CreateRawModel(vao, len(indicies))
}

//LoadToVAOWithTexCoords returns an RawModel and creates an VAO and save handler in Loader struct
func (l *Loader) LoadToVAOWithTexCoords(positions []float32, indicies []int32, texCoords []float32) *RawModel {
	vao := l.createVAO()
	l.bindIndiciesBuffer(indicies)
	l.storeDataInAttributeList(0, 3, positions)
	l.storeDataInAttributeList(1, 2, texCoords)
	l.unbindVAO()

	return CreateRawModel(vao, len(indicies))
}

func (l *Loader) UpdateVAO(vao uint32, positions []float32, indicies []int32) {
	helper.BindVertexArray(vao)
	defer l.unbindVAO()

	//var vbo uint32
	//gl.GetIntegerv(gl.ARRAY_BUFFER_BINDING, &vbo)
	//gl.BindBuffer

}

func (l *Loader) createVAO() uint32 {
	//log.Println("createVAO??")
	//debug.PrintStack()
	vao := helper.GenerateVertexArray(1)
	l.vaos = append(l.vaos, vao)
	helper.BindVertexArray(vao)

	return vao
}

//LoadTexture Loads an Texture returns the handle
func (l *Loader) LoadTexture(filename string) uint32 {
	log.Println("LoadTexture...:", filename)
	tex := NewTextureFromFile(filename)
	handle := tex.GetHandle()
	l.textures = append(l.textures, CreateTextureCleaner(filename, handle))
	return handle
}

func (l *Loader) storeDataInAttributeList(attributeNumber uint32, coordinateSize int32, data []float32) {

	vbo := helper.GenBuffers(1)
	log.Println("storeDateInAttribute VBO:", vbo)
	l.vbos = append(l.vbos, vbo)
	//gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	helper.BindBuffer(helper.GlArrayBuffer, vbo)

	//gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
	helper.BufferData(helper.GlArrayBuffer, len(data)*4, helper.Ptr(data), helper.GlStaticDraw)
	//gl.EnableVertexAttribArray(0)
	//gl.VertexAttribPointer(attributeNumber, coordinateSize, gl.FLOAT, false, 0, gl.PtrOffset(0))
	helper.VertexAttribPointer(attributeNumber, coordinateSize, helper.GlFloat, false, 0, helper.PtrOffset(0))
	//gl.DisableVertexAttribArray(0)

	helper.BindBuffer(helper.GlArrayBuffer, 0)
}

func (l *Loader) unbindVAO() {
	helper.BindVertexArray(0)
}

func (l *Loader) bindIndiciesBuffer(indicies []int32) {
	//var vbo uint32
	//gl.GenBuffers(1, &vbo)
	vbo := helper.GenBuffers(1)
	l.vbos = append(l.vbos, vbo)
	//gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vbo)
	helper.BindBuffer(helper.GlElementArrayBuffer, vbo)
	//gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicies)*4, gl.Ptr(indicies), gl.STATIC_DRAW)
	helper.BufferData(helper.GlElementArrayBuffer, len(indicies)*4, helper.Ptr(indicies), helper.GlStaticDraw)
}

//func (l *Loader) storeDataInIntBuffer( indicies []int32)

//CleanUP delete all VAOS and Buffers from opengl memory
func (l *Loader) CleanUP() {
	for _, vao := range l.vaos {
		helper.DeleteVertexArrary(1, &vao)
	}

	for _, vbo := range l.vbos {
		helper.DeleteBuffer(vbo)
	}

	for _, texture := range l.textures {
		texture.Remove()
	}

}
