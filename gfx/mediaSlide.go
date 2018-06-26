package gfx

import (
	"image"
	"log"
	"sync"
	"sync/atomic"

	"github.com/mcbernie/myopengl/glHelper"
)

//MediaSlide A default Slide for all types of Media
type MediaSlide struct {
	name string
	uid  string

	delay float64

	//Tex CurrentTexture
	Tex      *Texture
	isLoaded int32

	imageMux             sync.Mutex
	imageReadyForReplace bool
	img                  image.Image

	IsVideo bool
}

func (s *MediaSlide) GetDelay() float64 {
	return s.delay
}

//GetUid Retruns own uid
func (s *MediaSlide) GetUid() string {
	return s.uid
}

func createSlide(uid string, isVideo bool) *MediaSlide {

	tex := NewTexture(glHelper.GlClampToEdge, glHelper.GlClampToEdge)
	return &MediaSlide{
		uid:      uid,
		Tex:      tex,
		isLoaded: 0,
		IsVideo:  isVideo,
	}

}

func (s *MediaSlide) Play() {
}

func (s *MediaSlide) setIsLoading(loading bool) {
	var i int32
	if loading {
		i = 1
	}
	atomic.StoreInt32(&(s.isLoaded), int32(i))
}

//IsLoading check if texture is loading
func (s *MediaSlide) IsLoading() bool {
	if atomic.LoadInt32(&(s.isLoaded)) != 0 {
		return true
	}
	return false

}

func (s *MediaSlide) Display() *Texture {
	return s.Tex
}

//SetFrame replace And set frame
func (s *MediaSlide) SetFrame(img image.Image) {
	s.imageMux.Lock()
	s.img = img
	s.imageReadyForReplace = true
	s.imageMux.Unlock()
}

//Delete remove texture from memory
func (s *MediaSlide) Delete() {
	log.Println("delete texture from ", s.uid)
	glHelper.AddFunction(func() {
		s.Tex.Delete()
	})

}

//CleanUP remove texture from memory after closing
func (s *MediaSlide) CleanUP() {
	log.Println("delete texture from ", s.uid)
	s.Tex.Delete()
}

func (s *MediaSlide) Update() {

}

func (s *MediaSlide) BackgroundThread() {

}
