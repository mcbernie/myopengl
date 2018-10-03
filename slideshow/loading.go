package slideshow

import (
	"io/ioutil"
	"log"
	"strings"
)

type loader struct {
	uid    string
	path   string
	remote bool
}

//CreateNewSlideFromImageFile create a new slide from a image
func (s *Slideshow) CreateNewSlideFromImageFile(path string, name string, duration float64) (Slide, error) {
	slide, err := NewSlideFromImageFile(path, name, duration)
	if err != nil {
		return nil, err
	}

	s.slides = append(s.slides, slide)
	return slide, nil

}

//CreateNewSlideFromRemote create a new slide from a image url
func (s *Slideshow) CreateNewSlideFromRemote(url string, name string, duration float64) (Slide, error) {
	slide, err := NewSlideFromRemoteImage(url, name, duration)
	if err != nil {
		return nil, err
	}

	s.slides = append(s.slides, slide)
	return slide, nil
}

//AppendNewSlideFromRemote Add a new Slide to Slideshow
func (s *Slideshow) AppendNewSlideFromRemote(url string, uid string) {
	s.loaders = append(s.loaders, &loader{
		uid:    uid,
		path:   url,
		remote: true,
	})
}

//LoadImageFromRemote take currentslide an load new image
/*func (s *Slideshow) LoadImageFromRemote(url string) {
	go func() {

		update := s.currentSlide - 1
		if update < 0 {
			update = len(s.slides)
		}

		s.slides[update].LoadImageFromRemote(url)
	}()
}*/

//LoadTransitions load all transitions with .glsl file extension from an specified path
func (s *Slideshow) LoadTransitions(path string, projection [16]float32) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() && strings.Contains(f.Name(), ".glsl") {
			sa, _ := loadFromFile(path + "/" + f.Name())
			s.transitions = append(s.transitions, MakeTransition(Stretch, sa, f.Name(), projection))
		}
	}
}

//RemoveSlide removes a slide from slideshow
func (s *Slideshow) RemoveSlide(uid string) {
	for i, slide := range s.slides {
		if slide.GetUid() == uid {
			slide.Delete()
			newslides := append(s.slides[:i], s.slides[i+1:]...)

			s.slides = newslides
		}
	}
}
