package gfx

import (
	"image"
	"log"
	"net/http"
	"os"
	"sync"
	"sync/atomic"

	"github.com/mcbernie/myopengl/glHelper"

	_ "image/jpeg" // Import JPEG Decoding
	_ "image/png"  // Import PNG Decoding
)

//Slide A Simple Slideshow element
type Slide struct {
	name string
	uid  string

	//Tex CurrentTexture
	Tex      *Texture
	isLoaded int32

	imageMux             sync.Mutex
	imageReadyForReplace bool
	img                  image.Image
	IsVideo              bool
}

//GetUid Retruns own uid
func (s *Slide) GetUid() string {
	return s.uid
}

//NewSlideFromImageFile Create slide from image
func NewSlideFromImageFile(path string, uid string) (*Slide, error) {

	s := createSlide(uid, false)
	err := s.LoadImageFromFile(path)

	return s, err

}

//NewSlideFromRemoteImage Create a slide from remote image
func NewSlideFromRemoteImage(url string, uid string) (*Slide, error) {

	ret := make(chan *Slide)
	glHelper.AddFunction(func() {
		ret <- createSlide(uid, false)
	})
	s := <-ret
	err := s.LoadImageFromRemote(url)

	return s, err
}

//NewSlideForVideo Create a new Slide for Video content
func NewSlideForVideo(uid string) *Slide {

	return createSlide(uid, true)
}

func createSlide(uid string, isVideo bool) *Slide {

	tex := NewTexture(glHelper.GlClampToEdge, glHelper.GlClampToEdge)
	return &Slide{
		uid:      uid,
		Tex:      tex,
		isLoaded: 0,
		IsVideo:  isVideo,
	}

}

//LoadImageFromRemote load an image from a remote location
func (s *Slide) LoadImageFromRemote(url string) error {
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
func (s *Slide) LoadImageFromFile(path string) error {
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

func (s *Slide) setIsLoading(loading bool) {
	var i int32
	if loading {
		i = 1
	}
	atomic.StoreInt32(&(s.isLoaded), int32(i))
}

//IsLoading check if texture is loading
func (s *Slide) IsLoading() bool {
	if atomic.LoadInt32(&(s.isLoaded)) != 0 {
		return true
	}
	return false

}

//Update lock if there is a news image
func (s *Slide) Update() {
	// check if there is a new image ready for replacing...
	s.imageMux.Lock()
	defer s.imageMux.Unlock()

	if s.IsVideo {
		s.Tex.SetImage(s.img, glHelper.GlClampToEdge, glHelper.GlClampToEdge)
	} else {
		if s.imageReadyForReplace {
			s.Tex.SetImage(s.img, glHelper.GlClampToEdge, glHelper.GlClampToEdge)
			s.imageReadyForReplace = false
		}
	}
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

//SetFrame replace And set frame
func (s *Slide) SetFrame(img image.Image) {
	s.imageMux.Lock()
	s.img = img
	s.imageReadyForReplace = true
	s.imageMux.Unlock()
}

//Delete remove texture from memory
func (s *Slide) Delete() {
	log.Println("delete texture from ", s.uid)
	s.Tex.Delete()
}
