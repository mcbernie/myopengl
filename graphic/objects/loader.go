package objects

import (

	//"github.com/go-gl/gl/v4.1-core/gl" // OR:

	"log"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/glHelper"
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
	glHelper.BindVertexArray(vao)
	defer l.unbindVAO()

	//var vbo uint32
	//gl.GetIntegerv(gl.ARRAY_BUFFER_BINDING, &vbo)
	//gl.BindBuffer

}

func (l *Loader) createVAO() uint32 {
	//log.Println("createVAO??")
	//debug.PrintStack()
	vao := glHelper.GenerateVertexArray(1)
	l.vaos = append(l.vaos, vao)
	glHelper.BindVertexArray(vao)

	return vao
}

//LoadTexture Loads an Texture returns the handle
func (l *Loader) LoadTexture(filename string) uint32 {
	log.Println("LoadTexture...:", filename)
	tex := gfx.NewTextureFromFile(filename)
	handle := tex.GetHandle()
	l.textures = append(l.textures, CreateTextureCleaner(filename, handle))
	return handle
}

func (l *Loader) storeDataInAttributeList(attributeNumber uint32, coordinateSize int32, data []float32) {

	vbo := glHelper.GenBuffers(1)
	log.Println("storeDateInAttribute VBO:", vbo)
	l.vbos = append(l.vbos, vbo)
	//gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	glHelper.BindBuffer(glHelper.GlArrayBuffer, vbo)

	//gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)
	glHelper.BufferData(glHelper.GlArrayBuffer, len(data)*4, glHelper.Ptr(data), glHelper.GlStaticDraw)
	//gl.EnableVertexAttribArray(0)
	//gl.VertexAttribPointer(attributeNumber, coordinateSize, gl.FLOAT, false, 0, gl.PtrOffset(0))
	glHelper.VertexAttribPointer(attributeNumber, coordinateSize, glHelper.GlFloat, false, 0, glHelper.PtrOffset(0))
	//gl.DisableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (l *Loader) unbindVAO() {
	glHelper.BindVertexArray(0)
}

func (l *Loader) bindIndiciesBuffer(indicies []int32) {
	//var vbo uint32
	//gl.GenBuffers(1, &vbo)
	vbo := glHelper.GenBuffers(1)
	l.vbos = append(l.vbos, vbo)
	//gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vbo)
	glHelper.BindBuffer(glHelper.GlElementArrayBuffer, vbo)
	//gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicies)*4, gl.Ptr(indicies), gl.STATIC_DRAW)
	glHelper.BufferData(glHelper.GlElementArrayBuffer, len(indicies)*4, glHelper.Ptr(indicies), glHelper.GlStaticDraw)
}

//func (l *Loader) storeDataInIntBuffer( indicies []int32)

//CleanUP delete all VAOS and Buffers from opengl memory
func (l *Loader) CleanUP() {
	for _, vao := range l.vaos {
		glHelper.DeleteVertexArrary(1, &vao)
	}

	for _, vbo := range l.vbos {
		gl.DeleteBuffers(1, &vbo)
	}

	for _, texture := range l.textures {
		texture.Remove()
	}

}
