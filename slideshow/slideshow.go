package slideshow

import (
	"math"
	"math/rand"

	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/graphic/objects"
)

//Slideshow the Struct for the slideshow
type Slideshow struct {
	model *objects.RawModel

	loaders []*loader

	slides      []*gfx.Slide
	transitions []*gfx.Transition
	box         []float32

	currentIndex int

	currentSlide int
	nextSlide    int

	currentTransition int
	nextTransition    int

	delay    float64
	duration float64

	/*windowWidth  float32
	windowHeight float32*/
}

//MakeSlideshow Generates the slideshow
func MakeSlideshow(defaultDelay, defaultDuration float64, loader *objects.Loader) *Slideshow {

	verts := []float32{
		-1.0, 1.0, 0, //V0
		-1.0, -1.0, 0, //V1
		1.0, -1.0, 0, //V2
		1.0, 1.0, 0, //V3
	}

	inds := []int32{
		0, 1, 3,
		3, 1, 2,
	}

	model := loader.LoadToVAO(verts, inds)

	s := &Slideshow{
		model: model,

		currentIndex:      0,
		currentTransition: 0,
		currentSlide:      0,
		nextSlide:         1,

		delay:    defaultDelay,
		duration: defaultDuration,
	}

	return s
}

//Render Render the transitions
func (s *Slideshow) Render(time float64, renderer *objects.Renderer) {
	s.renderTransition(time)
	renderer.Render(s.model)
}

//UpdateWindowSize get Called if the window would resized
/*func (s *Slideshow) UpdateWindowSize(width, height float32) {
	s.windowWidth = width
	s.windowHeight = height
}*/

func (s *Slideshow) renderTransition(time float64) {
	index := s.index(time)

	aviableSlides := s.onlyAviableSlides()

	if index != s.currentIndex {
		s.currentIndex = index
		s.currentTransition = rand.Intn(len(s.transitions))
		//log.Println("current transition:", d.transitions[d.currentTransition].Name)
	}

	s.currentSlide = index % (len(aviableSlides))
	s.nextSlide = (index + 1) % (len(aviableSlides))

	from := aviableSlides[s.currentSlide]
	to := aviableSlides[s.nextSlide]

	for _, slide := range s.slides {
		slide.Update()
	}

	transition := s.transitions[s.currentTransition]
	transition.Draw(s.progress(time), from.Tex, to.Tex)

}

func (s *Slideshow) onlyAviableSlides() []*gfx.Slide {

	r := make([]*gfx.Slide, 0)
	for _, s := range s.slides {
		if !s.IsLoading() {
			r = append(r, s)
		}
	}

	return r
}

func (s *Slideshow) total() float64 {
	return s.delay + s.duration
}

func (s *Slideshow) progress(time float64) float32 {
	return float32(math.Max(0, (time-float64(s.index(time))*s.total()-s.delay)/s.duration))
}

func (s *Slideshow) index(time float64) int {
	return int(math.Floor(time / (s.delay + s.duration)))
}

//Delete remove all transitions and all slides from memory
func (s *Slideshow) Delete() {

	for _, transition := range s.transitions {
		transition.Delete()
	}

	for _, slide := range s.slides {
		slide.Delete()
	}

}
