package fonts

type word struct {
	characters []*character
	width      float32
	fontSize   float32
}

func createWord(fontSize float32) *word {
	return &word{
		fontSize: fontSize,
	}
}

func (w *word) addCharacter(char *character) {
	w.characters = append(w.characters, char)
	w.width += char.xAdvance * w.fontSize
}

func (w *word) getCharacters() []*character {
	return w.characters
}

func (w *word) getWordWidth() float32 {
	return w.width
}
