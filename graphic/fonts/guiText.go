package fonts

import "log"

type GUIText struct {
	textString string
	fontSize   float32

	textMeshVao uint32
	vertexCount int32
	colour      [3]float32

	position      [2]float32
	bottomRight   [2]float32
	lineMaxSize   float32
	numberOfLines int32

	font *FontType

	centerText bool
}

func CreateGuiText(text string, fontSize float32, font *FontType, position [2]float32, maxLineLength float32,
	centered bool) *GUIText {

	g := &GUIText{
		textString:  text,
		fontSize:    fontSize,
		font:        font,
		position:    position,
		lineMaxSize: maxLineLength,

		centerText: centered,
		colour:     [3]float32{0.0, 0.0, 0.0},
	}

	TextMaster.LoadText(g)

	return g

}

func (g *GUIText) GetSize() [2]float32 {

	log.Println("bottomRight:", g.bottomRight)
	log.Println("position:", g.position)
	x := g.bottomRight[0] - g.position[0]
	y := g.bottomRight[1] - g.position[1]

	return [2]float32{x, y}
}

func (g *GUIText) SetNumberOfLines(number int32) {
	g.numberOfLines = number
}

func (g *GUIText) Remove() {
	TextMaster.RemoveText(g)
}

func (g *GUIText) getFont() *FontType {
	return g.font
}

func (g *GUIText) SetColour(rC, gC, bC float32) {
	g.colour[0] = rC
	g.colour[1] = gC
	g.colour[2] = bC
}

func (g *GUIText) SetColourRGB(rC, gC, bC int) {
	g.colour[0] = (1.0 / 255.0) * float32(rC)
	g.colour[1] = (1.0 / 255.0) * float32(gC)
	g.colour[2] = (1.0 / 255.0) * float32(bC)
}

func (g *GUIText) getMesh() uint32 {
	return g.textMeshVao
}

/**
 * Set the VAO and vertex count for this text.
 *
 * @param vao
 *            - the VAO containing all the vertex data for the quads on
 *            which the text will be rendered.
 * @param verticesCount
 *            - the total number of vertices in all of the quads.
 */
func (g *GUIText) setMeshInfo(vao uint32, verticesCount int32) {
	g.textMeshVao = vao
	g.vertexCount = verticesCount
}

func (g *GUIText) GetlineMaxSize() float32 {
	return g.lineMaxSize
}
