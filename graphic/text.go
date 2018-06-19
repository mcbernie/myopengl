package graphic

import (
	"log"
	"os"

	"github.com/andrebq/gas"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/mcbernie/myopengl/gltext"
)

var fonts [16]*gltext.Font

func initFont() {
	file, err := gas.Abs("assets/fonts/luxisr.ttf")
	if err != nil {
		log.Printf("Find font file: %v", err)
		return
	}

	//defer file.

	for i := range fonts {
		fonts[i], err = loadFont(file, int32(12+i))
		if err != nil {
			log.Printf("LoadFont: %v", err)
			return
		}

		//defer fonts[i].Release()
	}
}

// loadFont loads the specified font at the given scale.
func loadFont(file string, scale int32) (*gltext.Font, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer fd.Close()

	return gltext.LoadTruetype(fd, scale, 32, 127, gltext.LeftToRight)
}

// drawString draws the same string for each loaded font.
func (d *Display) drawString(x, y float32, str string, size int) error {
	_, h := fonts[size].GlyphBounds()
	y = y + float32(size*h)

	// Draw a rectangular backdrop using the string's metrics.
	sw, sh := fonts[size].Metrics(str)
	gl.Color4f(0.5, 0.1, 0.1, 0.7)
	gl.Rectf(x, y, x+float32(sw), y+float32(sh))

	// Render the string.
	gl.Color4f(0.5, 0.7, 0.1, 1)
	err := fonts[size].Printf(x, y, str)
	if err != nil {
		//log.Println("error:", err)
		return err
	}

	return nil
}
