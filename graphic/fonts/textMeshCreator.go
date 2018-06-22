package fonts

const (
	lineHeight = 0.03
	spaceASCII = 32
)

type textMeshCreator struct {
	metaData      *metaFile
	maxLineHeight float32
}

func makeTextMeshCreator(metaFile string) (*textMeshCreator, error) {

	metaData, err := loadMetaFile(metaFile)

	if err != nil {
		panic("Error On Loading Metafile!")
	}

	return &textMeshCreator{
		metaData: metaData,
	}, nil
}

func (t *textMeshCreator) createTextMesh(text *GUIText) *textMeshData {
	lines := t.createStructure(text)
	data := t.createQuadVertices(text, lines)
	return data
}

func (t *textMeshCreator) createStructure(text *GUIText) []*line {
	chars := text.textString

	var lines []*line

	currentLine := createLine(t.metaData.spaceWidth, text.fontSize, text.lineMaxSize)
	currentWord := createWord(text.fontSize)

	for _, c := range chars {
		ascii := int32(c)
		if ascii == spaceASCII {
			if added := currentLine.attemptToAddWord(currentWord); !added {
				lines = append(lines, currentLine)
				currentLine = createLine(t.metaData.spaceWidth, text.fontSize, text.lineMaxSize)
				currentLine.attemptToAddWord(currentWord)
			}

			currentWord = createWord(text.fontSize)
			continue
		}

		if ascii == int32('\n') {
			lines = append(lines, currentLine)
			currentLine = createLine(t.metaData.spaceWidth, text.fontSize, text.lineMaxSize)
			currentWord = createWord(text.fontSize)
			continue
		}

		cc := t.metaData.getCharacter(ascii)
		currentWord.addCharacter(cc)
	}

	t.completeStructure(&lines, currentLine, currentWord, text)
	return lines
}

func (t *textMeshCreator) completeStructure(lines *[]*line, currentLine *line, currentWord *word, text *GUIText) {
	if added := currentLine.attemptToAddWord(currentWord); !added {
		*lines = append(*lines, currentLine)
		currentLine = createLine(t.metaData.spaceWidth, text.fontSize, text.lineMaxSize)
		currentLine.attemptToAddWord(currentWord)
	}
	*lines = append(*lines, currentLine)
}

func (t *textMeshCreator) createQuadVertices(text *GUIText, lines []*line) *textMeshData {

	text.SetNumberOfLines(int32(len(lines)))

	var curserX, curserY float32

	var vertices []float32
	var textureCoords []float32

	for _, l := range lines {
		if text.centerText {
			curserX = (l.maxLength - l.getLineLength()) / 2.0
		}

		for _, w := range l.getWords() {
			for _, letter := range w.getCharacters() {
				//log.Println("add tex coords -> ", letter.xTextureCoord, letter.yTextureCoord)
				addVerticesForCharacter(curserX, curserY, letter, text.fontSize, &vertices)
				addTexCoords(&textureCoords, letter.xTextureCoord, letter.yTextureCoord,
					letter.xMaxTextureCoord, letter.yMaxTextureCoord)
				curserX += letter.xAdvance * text.fontSize
			}
			curserX += t.metaData.spaceWidth * text.fontSize
		}
		curserX = 0
		curserY += lineHeight * text.fontSize
	}

	//log.Println("TextureCoords:->", textureCoords)
	return createTextMeshData(vertices, textureCoords)
}

func addVerticesForCharacter(curserX, curserY float32, letter *character, fontSize float32, vertices *[]float32) {
	x := curserX + (letter.xOffset * fontSize)
	y := curserY + (letter.yOffset * fontSize)
	maxX := x + (letter.sizeX * fontSize)
	maxY := y + (letter.sizeY * fontSize)
	properX := float32(2.0*x) - 1.0
	properY := float32(-2.0*y) + 1.0
	properMaxX := float32(2.0*maxX) - 1.0
	properMaxY := float32(-2.0*maxY) + 1.0
	addVertices(vertices, properX, properY, properMaxX, properMaxY)

}

func addVertices(vertices *[]float32, x, y, maxX, maxY float32) {
	*vertices = append(*vertices, x)
	*vertices = append(*vertices, y)
	*vertices = append(*vertices, x)
	*vertices = append(*vertices, maxY)
	*vertices = append(*vertices, maxX)
	*vertices = append(*vertices, maxY)
	*vertices = append(*vertices, maxX)
	*vertices = append(*vertices, maxY)
	*vertices = append(*vertices, maxX)
	*vertices = append(*vertices, y)
	*vertices = append(*vertices, x)
	*vertices = append(*vertices, y)
}

func addTexCoords(texCoords *[]float32, x, y, maxX, maxY float32) {
	*texCoords = append(*texCoords, x)
	*texCoords = append(*texCoords, y)
	*texCoords = append(*texCoords, x)
	*texCoords = append(*texCoords, maxY)
	*texCoords = append(*texCoords, maxX)
	*texCoords = append(*texCoords, maxY)
	*texCoords = append(*texCoords, maxX)
	*texCoords = append(*texCoords, maxY)
	*texCoords = append(*texCoords, maxX)
	*texCoords = append(*texCoords, y)
	*texCoords = append(*texCoords, x)
	*texCoords = append(*texCoords, y)
}
