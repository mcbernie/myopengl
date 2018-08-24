package objects

import (
	"github.com/mcbernie/myopengl/glHelper"
)

type modelBufferType struct {
	bufferType uint32
	vbo        uint32
}

func CreateBuffer(bufferType uint32) modelBufferType {
	vbo := glHelper.GenBuffers(1)
	return modelBufferType{
		vbo:        vbo,
		bufferType: bufferType,
	}
}
func (m modelBufferType) Bind() {
	glHelper.BindBuffer(m.bufferType, m.vbo)
}
func (m modelBufferType) UnBind() {
	glHelper.BindBuffer(m.bufferType, 0)
}
func (m modelBufferType) Delete() {
	glHelper.DeleteBuffer(m.vbo)
}

func (m modelBufferType) WriteData(data []float32) {
	glHelper.BufferData(m.bufferType, len(data)*4, glHelper.Ptr(data), glHelper.GlStaticDraw)
}

func (m modelBufferType) WriteDataInt(data []int32) {
	glHelper.BufferData(m.bufferType, len(data)*4, glHelper.Ptr(data), glHelper.GlStaticDraw)
}

func (m modelBufferType) PointDataToAttributeList(attributeNumber uint32, coordinateSize int32) {
	glHelper.VertexAttribPointer(attributeNumber, coordinateSize, glHelper.GlFloat, false, 0, glHelper.PtrOffset(0))
}
