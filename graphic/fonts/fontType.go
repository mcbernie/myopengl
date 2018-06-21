package fonts

type FontType struct {
	textureAtlas uint32
	loader       *textMeshCreator
}

/**
 * Creates a new font and loads up the data about each character from the
 * font file.
 *
 * @param textureAtlas
 *            - the ID of the font atlas texture.
 * @param fontFile
 *            - the font file containing information about each character in
 *            the texture atlas.
 */
func MakeFontType(textureAtlas uint32, fontFile string) *FontType {
	loader, err := makeTextMeshCreator(fontFile)

	if err != nil {
		panic("Error on create MeshCreator")
	}

	return &FontType{
		textureAtlas: textureAtlas,
		loader:       loader,
	}
}

func (f *FontType) loadText(text *GUIText) *textMeshData {
	return f.loader.createTextMesh(text)
}
