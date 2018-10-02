package objects

import (
	"log"

	"github.com/mcbernie/myopengl/graphic/helper"
)

type modelBufferType struct {
	bufferType uint32
	vbo        uint32
}

func CreateBuffer(bufferType uint32) modelBufferType {
	vbo := helper.GenBuffers(1)
	return modelBufferType{
		vbo:        vbo,
		bufferType: bufferType,
	}
}
func (m modelBufferType) Bind() {
	helper.BindBuffer(m.bufferType, m.vbo)
}
func (m modelBufferType) UnBind() {
	helper.BindBuffer(m.bufferType, 0)
}
func (m modelBufferType) Delete() {
	log.Println("remove modelBuffer:", m.vbo)
	helper.DeleteBuffer(m.vbo)
}

func (m modelBufferType) WriteData(data []float32) {
	helper.BufferData(m.bufferType, len(data)*4, helper.Ptr(data), helper.GlStaticDraw)
}

func (m modelBufferType) WriteDataInt(data []int32) {
	helper.BufferData(m.bufferType, len(data)*4, helper.Ptr(data), helper.GlStaticDraw)
}

func (m modelBufferType) PointDataToAttributeList(attributeNumber uint32, coordinateSize int32) {
	helper.VertexAttribPointer(attributeNumber, coordinateSize, helper.GlFloat, false, 0, helper.PtrOffset(0))
}
