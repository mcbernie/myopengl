package graphic

import (
	"io/ioutil"
	"log"

	"github.com/mcbernie/myopengl/gfx"
)

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
