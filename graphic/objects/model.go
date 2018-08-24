package objects

import (
	"github.com/mcbernie/myopengl/glHelper"
)

type Model struct {
	vao             uint32
	textures        []uint32
	indiciesBuffer  modelBufferType
	vertexCount     int32
	positionsBuffer modelBufferType
	texcoordsBuffer modelBufferType
}

func CreateTestModel(vao uint32, vertexCount int32) *Model {
	return &Model{
		vao:         vao,
		vertexCount: vertexCount,
	}
}

func CreateModel() *Model {
	m := Model{}
	m.createVAO()
	return &m
}

func CreateModelWithData(indicies []int32, positions []float32) *Model {
	m := CreateModel()
	m.bindVAO()
	m.indiciesBuffer = CreateBuffer(glHelper.GlElementArrayBuffer)
	m.writeIndicies(indicies)

	m.positionsBuffer = CreateBuffer(glHelper.GlArrayBuffer)
	m.writePositions(positions)

	m.unbindVAO()
	return m
}

func CreateModelWithDataTexture(indicies []int32, positions, texCoords []float32) *Model {
	m := CreateModel()
	m.bindVAO()
	m.indiciesBuffer = CreateBuffer(glHelper.GlElementArrayBuffer)
	m.writeIndicies(indicies)

	m.positionsBuffer = CreateBuffer(glHelper.GlArrayBuffer)
	m.writePositions(positions)

	m.texcoordsBuffer = CreateBuffer(glHelper.GlArrayBuffer)
	m.writeTexCoords(texCoords)

	m.unbindVAO()
	return m
}

func (m *Model) Delete() {
	m.indiciesBuffer.Delete()
	m.positionsBuffer.Delete()
	m.texcoordsBuffer.Delete()
	glHelper.DeleteVertexArrary(1, &m.vao)
}

func (m *Model) SetIndicies(data []int32) {
	m.bindVAO()
	m.writeIndicies(data)
}
func (m *Model) writeIndicies(data []int32) {
	m.vertexCount = int32(len(data))

	m.indiciesBuffer.Bind()
	m.indiciesBuffer.WriteDataInt(data)
	//m.indiciesBuffer.UnBind() <-- Eingebaut und fehler passiert!!
}

func (m *Model) SetPositions(data []float32) {
	m.bindVAO()
	m.writePositions(data)
}
func (m *Model) writePositions(data []float32) {
	m.positionsBuffer.Bind()
	m.positionsBuffer.WriteData(data)
	m.positionsBuffer.PointDataToAttributeList(0, 3)
	m.positionsBuffer.UnBind()
}

func (m *Model) SetTexCoords(data []float32) {
	m.bindVAO()
	m.writeTexCoords(data)
}
func (m *Model) writeTexCoords(data []float32) {
	m.texcoordsBuffer.Bind()
	m.texcoordsBuffer.WriteData(data)
	m.texcoordsBuffer.PointDataToAttributeList(1, 2)
	m.texcoordsBuffer.UnBind()
}

func (m *Model) createVAO() {
	m.vao = glHelper.GenerateVertexArray(1)
}
func (m *Model) bindVAO() {
	glHelper.BindVertexArray(m.vao)
}
func (m *Model) unbindVAO() {
	glHelper.BindVertexArray(0)
}

func (m *Model) Draw() {
	glHelper.DrawElements(glHelper.GlTriangles, m.vertexCount, glHelper.GlUnsignedInt, glHelper.PtrOffset(0))
}

func (m *Model) Bind() {
	m.bindVAO()
}

func (m *Model) UnBind() {
	m.unbindVAO()
}
