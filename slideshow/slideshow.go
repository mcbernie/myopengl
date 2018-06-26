package slideshow

import (
	"log"
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/graphic/objects"
)

//Slideshow the Struct for the slideshow
type Slideshow struct {
	//model *objects.RawModel
	SlideShowEntity *objects.Entity

	loaders []*loader

	slides      []gfx.Slide
	transitions []*gfx.Transition
	box         []float32

	currentIndex int

	currentSlide int
	nextSlide    int

	currentTransition int
	nextTransition    int

	delay        float64
	duration     float64
	defaultDelay float64
	/*windowWidth  float32
	windowHeight float32*/
}

//MakeSlideshow Generates the slideshow
func MakeSlideshow(defaultDelay, defaultDuration float64, loader *objects.Loader) *Slideshow {

	verts := []float32{
		-1.0, 1.0, -0.1, //V0
		-1.0, -1.0, -0.1, //V1
		1.0, -1.0, -0.1, //V2
		1.0, 1.0, -0.1, //V3
	}

	inds := []int32{
		0, 1, 3,
		3, 1, 2,
	}

	model := loader.LoadToVAO(verts, inds)
	entity := objects.MakeEntity(model, mgl32.Vec3{0, 0, -2.10}, 0, 0, 0, 1.0)

	s := &Slideshow{
		SlideShowEntity: entity,

		currentIndex:      0,
		currentTransition: 0,
		currentSlide:      0,
		nextSlide:         1,

		delay:        defaultDelay,
		duration:     defaultDuration,
		defaultDelay: defaultDelay,
	}

	return s
}

var previousTime float64
var progress float64
var index int
var currentDuration float64
var pause bool
var delayForTransition float64

var from, to gfx.Slide
var transition *gfx.Transition
var transitionId int

//Render Render the transitions
func (s *Slideshow) Render(time float64, renderer *objects.Renderer) {
	delayForTransition = 5.0
	aviableSlides := s.onlyAviableSlides()

	durationBetweenFrames := time - previousTime
	previousTime = time

	currentDuration += durationBetweenFrames
	if pause == false {
		//Animationsschleife

		progress = currentDuration / s.duration

		if progress >= 1 {
			// animation ist abgeschlossen jetzt gilt der delay!
			progress = 0.0      // Progress zurücksetzen
			currentDuration = 0 // dauer zurücksetzen

			index++ // index einen aufzählen
			//überlaufschutz für den index. es kann keinen höheren index als verfügbare slides geben
			if index > len(aviableSlides)-1 {
				index = 0
			}
			log.Println("currentSlide:", index)
			s.currentSlide = index
			s.nextSlide = (index + 1) % (len(aviableSlides))
			pause = true // pause aktivieren

			log.Println("in pause mode, currently shows:", to.GetUid())

		}

	} else {
		//Delay checker
		if from.GoToNextSlide(currentDuration) == true {
			log.Println(" from:", from.GetUid(), " to:", to.GetUid(), " duration:", currentDuration)
			pause = false       // delay ist abgelaufen, pause beenden
			currentDuration = 0 // dauer zurücksetzen
			transitionId = rand.Intn(len(s.transitions))
		} else {
			from.Play()
		}
	}

	transition = s.transitions[transitionId]
	from = aviableSlides[s.currentSlide]
	to = aviableSlides[s.nextSlide]

	for _, slide := range s.slides {
		slide.Update()
	}

	to.Play()

	transition.Draw(float32(progress),
		from.Display(),
		to.Display())

	//begin render Entity after all shader processing is done!
	renderer.RenderEntity(s.SlideShowEntity, transition.Shader)
}

func (s *Slideshow) onlyAviableSlides() []gfx.Slide {

	r := make([]gfx.Slide, 0)
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

// bei jedem update schauen...
/*
progress geht von 0.0 bis 1.0

zeit -
zeit / delay + duration
*
delay + duration

- delay
/ duration


*/

func (s *Slideshow) progress(time float64) float32 {
	return float32(math.Max(0, (time-float64(s.index(time))*s.total()-s.delay)/s.duration))
}

func (s *Slideshow) index(time float64) int {
	return int(math.Floor(time / (s.delay + s.duration)))
}

//CleanUP remove all transitions and all slides from memory
func (s *Slideshow) CleanUP() {

	for _, transition := range s.transitions {
		transition.CleanUP()
	}

	for _, slide := range s.slides {
		slide.CleanUP()
	}

}
