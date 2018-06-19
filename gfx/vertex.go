package gfx

//Vertex simple Vertex structure
type Vertex struct {
	V []float32
}

//CreateVertex setup a simple Vertex
func CreateVertex(x, y, z, texCoordX, texCoordY float32) *Vertex {
	return &Vertex{V: []float32{x, y, z, texCoordX, texCoordY}}
}

//Box shows a simple Box
type Box struct {
	x, y, z    float32
	texScaling float32
	texture    Texture
	vertices   []*Vertex
}

//GetVertices returns the vertices
func (box *Box) GetVertices() []float32 {
	var v []float32
	for _, vertex := range box.vertices {
		v = append(v, vertex.V...)
	}

	return v
}

//CreateBox Create a Simple Würfel / Box
func CreateBox(tX float32) *Box {

	// 0.5 meaning 0 position is in center of box /Pivot Point is center
	return &Box{
		x:          0,
		y:          0,
		z:          0,
		texScaling: tX,
		vertices: []*Vertex{
			// Unkown:
			CreateVertex(-0.5, -0.5, -0.5, 0.0, 0.0),
			CreateVertex(0.5, -0.5, -0.5, tX, 0.0),
			CreateVertex(0.5, 0.5, -0.5, tX, tX),
			CreateVertex(0.5, 0.5, -0.5, tX, tX),
			CreateVertex(-0.5, 0.5, -0.5, 0.0, tX),
			CreateVertex(-0.5, -0.5, -0.5, 0.0, 0.0),
			// Front:
			CreateVertex(-0.5, -0.5, 0.5, 0.0, tX),
			CreateVertex(0.5, -0.5, 0.5, tX, tX),
			CreateVertex(0.5, 0.5, 0.5, tX, 0.0),
			CreateVertex(0.5, 0.5, 0.5, tX, 0.0),
			CreateVertex(-0.5, 0.5, 0.5, 0.0, 0.0),
			CreateVertex(-0.5, -0.5, 0.5, 0.0, tX),
			// Left
			CreateVertex(-0.5, 0.5, 0.5, tX, 0.0),
			CreateVertex(-0.5, 0.5, -0.5, 0.0, 0.0),
			CreateVertex(-0.5, -0.5, -0.5, 0.0, tX),
			CreateVertex(-0.5, -0.5, -0.5, 0.0, tX),
			CreateVertex(-0.5, -0.5, 0.5, tX, tX),
			CreateVertex(-0.5, 0.5, 0.5, tX, 0.0),
			// Right
			CreateVertex(0.5, 0.5, 0.5, 0.0, 0.0),
			CreateVertex(0.5, 0.5, -0.5, tX, 0.0),
			CreateVertex(0.5, -0.5, -0.5, tX, tX),
			CreateVertex(0.5, -0.5, -0.5, tX, tX),
			CreateVertex(0.5, -0.5, 0.5, 0.0, tX),
			CreateVertex(0.5, 0.5, 0.5, 0.0, 0.0),
			// Else
			CreateVertex(-0.5, -0.5, -0.5, 0.0, tX),
			CreateVertex(0.5, -0.5, -0.5, tX, tX),
			CreateVertex(0.5, -0.5, 0.5, tX, 0.0),
			CreateVertex(0.5, -0.5, 0.5, tX, 0.0),
			CreateVertex(-0.5, -0.5, 0.5, 0.0, 0.0),
			CreateVertex(-0.5, -0.5, -0.5, 0.0, tX),
			// Else
			CreateVertex(-0.5, 0.5, -0.5, 0.0, tX),
			CreateVertex(0.5, 0.5, -0.5, tX, tX),
			CreateVertex(0.5, 0.5, 0.5, tX, 0.0),
			CreateVertex(0.5, 0.5, 0.5, tX, 0.0),
			CreateVertex(-0.5, 0.5, 0.5, 0.0, 0.0),
			CreateVertex(-0.5, 0.5, -0.5, 0.0, tX),
		},
	}
}
