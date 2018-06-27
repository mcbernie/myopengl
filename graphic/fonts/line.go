package fonts

type line struct {
	maxLength         float32
	spaceSize         float32
	words             []*word
	currentLineLength float32
}

func createLine(spaceWidth, fontSize, maxLength float32) *line {
	return &line{
		spaceSize: spaceWidth * fontSize,
		maxLength: maxLength,
	}
}

func (l *line) attemptToAddWord(w *word) bool {
	additionalLength := w.getWordWidth()
	if len(l.words) < 1 {
		additionalLength += 0
	} else {
		additionalLength += l.spaceSize
	}

	if l.maxLength == -1 {
		l.words = append(l.words, w)
		return true
	}

	if l.currentLineLength+additionalLength <= l.maxLength {
		l.words = append(l.words, w)
		return true
	}
	return false
}

/**
 * @return The current screen-space length of the line.
 */
func (l *line) getLineLength() float32 {
	return l.currentLineLength
}

/**
 * @return The list of words in the line.
 */
func (l *line) getWords() []*word {
	return l.words
}
