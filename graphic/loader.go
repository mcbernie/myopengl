package graphic

import (
	"io/ioutil"
	"log"
)

//LoadRemoteImage load an remote image for slideshow
func (d *Display) LoadRemoteImage(path string, uid string) {
	d.slideshow.CreateNewSlideFromRemote(path, uid, 5.0)
}

//LoadLocalImage load an local image for slideshow
func (d *Display) LoadLocalImage(path string, uid string) {
	log.Println("LoadLocalImage:", uid)
	d.slideshow.CreateNewSlideFromImageFile(path, uid, 5.0)
}

func (d *Display) LoadVideo(path string, uid string) {
	log.Println("Load a Video:", uid)

	slide, err := d.slideshow.CreateNewSlideForVideo(path, uid)
	if err != nil {
		log.Println("Error on Load Slide")
	} else {
		log.Println("Start Background Thread for ", uid)
		slide.BackgroundThread()
	}

}

func (d *Display) LoadVideoAtRuntime(path, uid string, withDuration float64) {
	log.Println("Load a Video At Runtime:", uid)

	_, err := d.slideshow.CreateNewSlideForVideoRemote(path, uid, withDuration)
	if err != nil {
		log.Println("Error on Load Slide")
	}
}

//LoadImagesFromPath load all images from a specified path and put it in slide array
func (d *Display) LoadImagesFromPath(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("Fatal Error on LoadImagesFromPath:", err)
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
