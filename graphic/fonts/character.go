package fonts

type character struct {
	id               int32
	xTextureCoord    float32
	yTextureCoord    float32
	xMaxTextureCoord float32
	yMaxTextureCoord float32
	xOffset          float32
	yOffset          float32
	sizeX            float32
	sizeY            float32
	xAdvance         float32
}

func createCharacter(id int32,
	xTextureCoord, yTextureCoord, xTexSize, yTexSize,
	xOffset, yOffset, sizeX, sizeY, xAdvance float32) *character {

	return &character{
		id:               id,
		xTextureCoord:    xTextureCoord,
		yTextureCoord:    yTextureCoord,
		xOffset:          xOffset,
		yOffset:          yOffset,
		sizeX:            sizeX,
		sizeY:            sizeY,
		xMaxTextureCoord: xTexSize + xTextureCoord,
		yMaxTextureCoord: yTexSize + yTextureCoord,
		xAdvance:         xAdvance,
	}

}
