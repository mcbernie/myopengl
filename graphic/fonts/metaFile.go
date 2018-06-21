package fonts

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	padTop = iota
	padLeft
	padBottom
	padRight

	desiredPadding = padRight
)

const (
	splitter        = " "
	numberSeparator = ","
)

type metaFile struct {
	aspectRatio            float32
	verticalPerPixelSize   float32
	horizontalPerPixelSize float32
	spaceWidth             float32
	padding                []int32
	paddingWidth           int32
	paddingHeight          int32
	metaData               map[int32]*character
	values                 map[string]string
	scanner                *bufio.Scanner
}

func loadMetaFile(path string) (*metaFile, error) {
	m := metaFile{
		aspectRatio: 1,
		values:      make(map[string]string),
		metaData:    make(map[int32]*character),
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	m.scanner = scanner

	m.loadPaddingData()
	m.loadLineSizes()
	imageWidth := m.getValueOfVariable("scaleW")
	m.loadCharacterData(imageWidth)

	return &m, nil
}

func (m *metaFile) processNextLine() bool {
	m.values = make(map[string]string)

	if inRun := m.scanner.Scan(); !inRun {
		// End of File
		return false
	}

	line := m.scanner.Text()
	for _, part := range strings.Split(line, splitter) {
		valuePairs := strings.Split(part, "=")
		if len(valuePairs) == 2 {
			m.values[valuePairs[0]] = valuePairs[1]
		}
	}

	return true
}

func (m *metaFile) getValueOfVariable(variable string) int32 {
	t := m.values[variable]

	//log.Println("T:", t)
	conv, err := strconv.ParseInt(t, 10, 32)
	if err != nil {
		log.Printf("Error on Parseing %s with Value %s\n", variable, t)
		return -1
	}

	return int32(conv)
}

func (m *metaFile) getValuesOfVariable(variable string) []int32 {
	numbers := strings.Split(m.values[variable], numberSeparator)
	var actualValues []int32

	for _, number := range numbers {

		if conv, err := strconv.ParseInt(number, 10, 32); err != nil {
			log.Printf("Error on Parseing %s with Value %s\n", variable, number)
		} else {
			actualValues = append(actualValues, int32(conv))
		}
	}

	return actualValues
}

/**
 * Loads the data about how much padding is used around each character in
 * the texture atlas.
 */
func (m *metaFile) loadPaddingData() {
	m.processNextLine()
	m.padding = m.getValuesOfVariable("padding")

	m.paddingWidth = m.padding[padLeft] + m.padding[padRight]
	m.paddingHeight = m.padding[padTop] + m.padding[padBottom]
}

/**
* Loads information about the line height for this font in pixels, and uses
* this as a way to find the conversion rate between pixels in the texture
* atlas and screen-space.
 */
func (m *metaFile) loadLineSizes() {
	m.processNextLine()
	lineHeightPixels := m.getValueOfVariable("lineHeight") - m.paddingHeight
	m.verticalPerPixelSize = lineHeight / float32(lineHeightPixels)
	m.horizontalPerPixelSize = m.verticalPerPixelSize / m.aspectRatio

}

/**
* Loads in data about each character and stores the data in the
* {@link Character} class.
*
* @param imageWidth
*            - the width of the texture atlas in pixels.
 */
func (m *metaFile) loadCharacterData(imageWidth int32) {
	m.processNextLine()
	m.processNextLine()

	for m.processNextLine() {
		c := m.loadCharacter(imageWidth)
		if c != nil {
			m.metaData[c.id] = c
		}
	}

}

/**
 * Loads all the data about one character in the texture atlas and converts
 * it all from 'pixels' to 'screen-space' before storing. The effects of
 * padding are also removed from the data.
 *
 * @param imageSize
 *            - the size of the texture atlas in pixels.
 * @return The data about the character.
 */
func (m *metaFile) loadCharacter(imageSize int32) *character {
	id := m.getValueOfVariable("id")
	if id == spaceASCII {
		m.spaceWidth = float32((m.getValueOfVariable("xadvance") - m.paddingWidth)) * m.horizontalPerPixelSize
		return nil
	}

	myX := m.getValueOfVariable("x")
	//log.Println("The Xtex??", imageSize)
	xTex := float32(myX+(m.padding[padLeft]-desiredPadding)) / float32(imageSize)
	yTex := float32(m.getValueOfVariable("y")+(m.padding[padTop]-desiredPadding)) / float32(imageSize)

	//log.Println("xTex:", xTex, " yTex:", yTex)
	width := m.getValueOfVariable("width") - (m.paddingWidth - (2 * desiredPadding))
	height := m.getValueOfVariable("height") - ((m.paddingHeight) - (2 * desiredPadding))

	quadWidth := float32(width) * m.horizontalPerPixelSize
	quadHeight := float32(height) * m.verticalPerPixelSize
	xTexSize := float32(width / imageSize)
	yTexSize := float32(height / imageSize)

	xOff := float32(m.getValueOfVariable("xoffset")+m.padding[padLeft]-desiredPadding) * m.horizontalPerPixelSize
	yOff := float32(m.getValueOfVariable("yoffset")+(m.padding[padTop]-desiredPadding)) * m.verticalPerPixelSize
	xAdvance := float32(m.getValueOfVariable("xadvance")-m.paddingWidth) * m.horizontalPerPixelSize

	return createCharacter(id, xTex, yTex, xTexSize, yTexSize, xOff, yOff, quadWidth, quadHeight, xAdvance)
}

func (m *metaFile) getCharacter(ascii int32) *character {
	return m.metaData[ascii]
}
