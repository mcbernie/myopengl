package slideshow

import (
	"image"
	"log"
	"net/http"

	_ "image/jpeg" // Import JPEG Decoding
	_ "image/png"  // Import PNG Decoding

	"github.com/mcbernie/myopengl/graphic/helper"
	"github.com/mcbernie/myopengl/graphic/objects"
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
	helper.AddFunction(func() {
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
	img, err := objects.LoadImageFromFile(path)
	if err != nil {
		log.Println("failed to load image from path:" + path)
		return err
	}

	s.setIsLoading(true)

	s.SetFrame(img)
	s.setIsLoading(false)
	return nil
}

//Update lock if there is a news image
func (s *ImageSlide) Update() {
	s.imageMux.Lock()
	if s.imageReadyForReplace {
		s.Tex.SetImage(s.img, helper.GlClampToEdge, helper.GlClampToEdge)
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
