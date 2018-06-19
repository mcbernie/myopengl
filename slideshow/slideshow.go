package slideshow

import (
	"math"
	"math/rand"
	//"log"

	//"github.com/go-gl/gl/v4.1-core/gl" // OR:
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/mcbernie/myopengl/gfx"
)

//Slideshow the Struct for the slideshow
type Slideshow struct {
	vbo uint32
	vao uint32

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

	windowWidth  float32
	windowHeight float32
}

//MakeSlideshow Generates the slideshow
func MakeSlideshow(defaultDelay, defaultDuration float64) *Slideshow {

	box := []float32{
		-1.0, -1.0,
		1.0, -1.0,
		-1.0, 1.0,
		-1.0, 1.0,
		1.0, -1.0,
		1.0, 1.0,
	}

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)

	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(box)*4, gl.Ptr(box), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))
	gl.BindBuffer(VBO, 0)

	gl.BindVertexArray(0)

	s := &Slideshow{
		vao: VAO,
		vbo: VBO,
		//box: box,

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
func (s *Slideshow) Render(time float64) {
	gl.BindVertexArray(s.vao)
	gl.EnableVertexAttribArray(0)
	s.renderTransition(time)
	gl.DisableVertexAttribArray(0)
	gl.BindVertexArray(0)
}

//UpdateWindowSize get Called if the window would resized
func (s *Slideshow) UpdateWindowSize(width, height float32) {
	s.windowWidth = width
	s.windowHeight = height
}

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
	transition.Draw(s.progress(time), from.Tex, to.Tex, s.windowWidth, s.windowHeight)

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
	gl.BindVertexArray(0)
	gl.DeleteBuffers(1, &s.vbo)
	gl.DeleteVertexArrays(1, &s.vao)

	for _, transition := range s.transitions {
		transition.Delete()
	}

	for _, slide := range s.slides {
		slide.Delete()
	}

}
