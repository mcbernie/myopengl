package objects

import ( //"math"
	//"math/rand"
	//"log"
	//"github.com/go-gl/gl/v4.1-core/gl" // OR:
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/mcbernie/myopengl/glHelper"
	//"github.com/mcbernie/myopengl/gfx"
)

//Loader holds all vaos and vbos handler for cleanup and access
// Helper Generating VAO by calling of LoadToVAO
type Loader struct {
	vaos []uint32
	vbos []uint32
}

//MakeLoader Creat
func MakeLoader() *Loader {
	return &Loader{}
}

//LoadToVAO returns an RawModel and creates an VAO and save handler in Loader struct
func (l *Loader) LoadToVAO(positions []float32) *RawModel {
	vao := l.createVAO()
	l.storeDataInAttributeList(0, positions)
	l.unbindVAO()

	return CreateRawModel(vao, int32(len(positions)/3))
}

func (l *Loader) createVAO() uint32 {

	vao := glHelper.GenerateVertexArray(1)

	l.vaos = append(l.vaos, vao)

	glHelper.BindVertexArray(vao)

	return vao
}

func (l *Loader) storeDataInAttributeList(attributeNumber uint32, data []float32) {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	l.vbos = append(l.vbos, vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	gl.BufferData(gl.ARRAY_BUFFER, len(data)*4, gl.Ptr(data), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(attributeNumber, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))
	gl.DisableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (l *Loader) unbindVAO() {
	glHelper.BindVertexArray(0)
}

//CleanUP delete all VAOS and Buffers from opengl memory
func (l *Loader) CleanUP() {
	for _, vao := range l.vaos {
		glHelper.DeleteVertexArrary(1, &vao)
	}

	for _, vbo := range l.vbos {
		gl.DeleteBuffers(1, &vbo)
	}

}
