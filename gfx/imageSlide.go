package gfx

import (
	"image"
	"log"
	"net/http"
	"os"

	"github.com/mcbernie/myopengl/glHelper"

	_ "image/jpeg" // Import JPEG Decoding
	_ "image/png"  // Import PNG Decoding
)

type ImageSlide struct {
	*MediaSlide
}

func newImageSlide(uid string) *ImageSlide {
	i := &ImageSlide{}
	i.MediaSlide = createSlide(uid, false)
	return i

}

//NewSlideFromImageFile Create slide from image
func NewSlideFromImageFile(path, uid string, duration float64) (*ImageSlide, error) {

	s := newImageSlide(uid)
	s.delay = duration
	err := s.LoadImageFromFile(path)

	return s, err

}

//NewSlideFromRemoteImage Create a slide from remote image
func NewSlideFromRemoteImage(url string, uid string, duration float64) (*ImageSlide, error) {

	ret := make(chan *ImageSlide)
	glHelper.AddFunction(func() {
		ret <- newImageSlide(uid)
	})
	s := <-ret
	s.delay = duration
	err := s.LoadImageFromRemote(url)

	return s, err
}

//LoadImageFromRemote load an image from a remote location
func (s *ImageSlide) LoadImageFromRemote(url string) error {
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
		return e
	}

	defer response.Body.Close()
	img, _, err := image.Decode(response.Body)
	if err != nil {
		log.Println("error on loading new image from remote", err)
		return err
	}

	s.setIsLoading(true)
	s.SetFrame(img)
	s.setIsLoading(false)
	return nil
}

//LoadImageFromFile Load an image from Path
func (s *ImageSlide) LoadImageFromFile(path string) error {
	img, err := LoadImageFromFile(path)
	if err != nil {
		log.Println("failed to load image from path:" + path)
		return err
	}

	s.setIsLoading(true)

	s.SetFrame(img)
	s.setIsLoading(false)
	return nil
}

//LoadImageFromFile Loads an image from file
func LoadImageFromFile(file string) (image.Image, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer imgFile.Close()

	// Decode detexts the type of image as long as its image/<type> is imported
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	return img, nil
}

//Update lock if there is a news image
func (s *ImageSlide) Update() {
	s.imageMux.Lock()
	if s.imageReadyForReplace {
		s.Tex.SetImage(s.img, glHelper.GlClampToEdge, glHelper.GlClampToEdge)
		s.imageReadyForReplace = false
	}
	s.imageMux.Unlock()
}

func (s *ImageSlide) GoToNextSlide(currentDuration float64) bool {
	if currentDuration >= s.delay {
		return true
	}
	return false
}
