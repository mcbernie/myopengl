package fonts

type textMeshData struct {
	vertexPositions []float32
	textureCoords   []float32
}

func createTextMeshData(vertexPositions []float32, textureCoords []float32) *textMeshData {

	return &textMeshData{
		vertexPositions: vertexPositions,
		textureCoords:   textureCoords,
	}
}

func (t *textMeshData) getVertexCount() int32 {
	return int32(len(t.vertexPositions) / 2)
}
