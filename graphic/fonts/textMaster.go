package fonts

import (
	"github.com/mcbernie/myopengl/graphic/objects"
)

type textMaster struct {
	loader *objects.Loader
	//for each fonttype a list for all guitexts
	texts    TextList
	renderer *fontRenderer
}

var TextMaster textMaster

func InitTextMaster(theLoader *objects.Loader) {
	TextMaster = textMaster{
		loader:   theLoader,
		renderer: createFontRenderer(),
		texts:    make(TextList),
	}
}

func (t *textMaster) Render() {
	t.renderer.render(t.texts)
}

func (t *textMaster) LoadText(text *GUIText) {
	font := text.font
	data := font.loadText(text)

	vao := t.loader.LoadVertexAndTextureToVAO(data.vertexPositions, data.textureCoords)

	text.setMeshInfo(vao, data.getVertexCount())

	if t.texts[font] == nil {
		t.texts[font] = []*GUIText{}
	}
	t.texts[font] = append(t.texts[font], text)
}

func (t *textMaster) RemoveText(text *GUIText) {
	// Remove A TEXT
	for font, innerTexts := range t.texts {
		var indexToDelete int
		found := false
		for i, t := range innerTexts {
			if t == text {
				indexToDelete = i
				found = true
				break
			}
		}

		if found {
			t.texts[font] = append(t.texts[font][:indexToDelete], t.texts[font][indexToDelete+1:]...)
			break
		}

	}
}

func (t *textMaster) CleanUP() {
	//call cleanup on renderer
	//t.renderer.CleanUP()
}
