package graphic

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/andrebq/gas"
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/graphic/gltext"
)

func (d *Display) InitFont() {
	file, err := gas.Abs("assets/fonts/luxisr.ttf")
	if err != nil {
		log.Printf("Find font file: %v", err)
		return
	}

	// Load the same font at different scale factors and directions.
	for i := range d.fonts {
		d.fonts[i], err = loadFont(file, int32(12+i))

		if err != nil {
			log.Printf("LoadFont: %v", err)
			return
		}

		//defer d.fonts[i].Release()
	}

}

//LoadRemoteImage load an remote image for slideshow
func (d *Display) LoadRemoteImage(path string, uid string) {
	d.slideshow.CreateNewSlideFromRemote(path, uid)
}

//LoadLocalImage load an local image for slideshow
func (d *Display) LoadLocalImage(path string, uid string) {
	d.slideshow.CreateNewSlideFromImageFile(path, uid)
}

//CreateVideoSlide Creates an video slide
func (d *Display) CreateVideoSlide(uid string) (*gfx.Slide, error) {
	return d.slideshow.CreateNewSlideForVideoFrames(uid)
}

//LoadImagesFromPath load all images from a specified path and put it in slide array
func (d *Display) LoadImagesFromPath(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			d.LoadLocalImage(path+"/"+f.Name(), f.Name())
		}
	}
}

//RemoveSlide outside in helper to remove a slide
func (d *Display) RemoveSlide(uid string) {
	d.slideshow.RemoveSlide(uid)
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
/*func drawString(x, y float32, str string) error {
	for i := range fonts {
		if fonts[i] == nil {
			continue
		}

		// We need to offset each string by the height of the
		// font. To ensure they don't overlap each other.
		_, h := fonts[i].GlyphBounds()
		y := y + float32(i*h)

		// Draw a rectangular backdrop using the string's metrics.
		sw, sh := fonts[i].Metrics(SampleString)
		gl.Color4f(0.1, 0.1, 0.1, 0.7)
		gl.Rectf(x, y, x+float32(sw), y+float32(sh))

		// Render the string.

		if err != nil {
			return err
		}
	}

	return nil
}*/
