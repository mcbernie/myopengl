package objects

import (

	//"github.com/go-gl/gl/v4.1-core/gl" // OR:
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/glHelper"
)

//Loader holds all vaos and vbos handler for cleanup and access
// Helper Generating VAO by calling of LoadToVAO
type Loader struct {
	vaos     []uint32
	vbos     []uint32
	textures []uint32
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

func (l *Loader) createVAO() uint32 {
	vao := glHelper.GenerateVertexArray(1)
	l.vaos = append(l.vaos, vao)
	glHelper.BindVertexArray(vao)

	return vao
}

func (l *Loader) LoadTexture(filename string) uint32 {

	tex := gfx.NewTextureFromFile(filename)
	handle := tex.GetHandle()
	l.textures = append(l.textures, handle)
	return handle
}

func (l *Loader) storeDataInAttributeList(attributeNumber uint32, coordinateSize int32, data []float32) {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	l.vbos = append(l.vbos, vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(attributeNumber, coordinateSize, gl.FLOAT, false, 0, gl.PtrOffset(0))
	gl.DisableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (l *Loader) unbindVAO() {
	glHelper.BindVertexArray(0)
}

func (l *Loader) bindIndiciesBuffer(indicies []int32) {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	l.vbos = append(l.vbos, vbo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indicies)*4, gl.Ptr(indicies), gl.STATIC_DRAW)
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
		gl.DeleteTextures(1, &texture)
	}

}
