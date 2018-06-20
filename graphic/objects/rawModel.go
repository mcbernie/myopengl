package objects

type RawModel struct {
	vao         uint32
	vertexCount int32
}

func CreateRawModel(vao uint32, vertexCount int32) *RawModel {
	return &RawModel{
		vao:         vao,
		vertexCount: vertexCount,
	}
}

func (rm *RawModel) GetVao() uint32 {
	return rm.vao
}

func (rm *RawModel) GetVertexCount() int32 {
	return rm.vertexCount
}
