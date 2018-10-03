package slideshow

import (
	"log"
	"math/rand"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/mcbernie/myopengl/graphic/helper"
	"github.com/mcbernie/myopengl/graphic/objects"
)

//Slideshow the Struct for the slideshow
type Slideshow struct {
	//model *objects.RawModel
	SlideShowEntity *objects.Entity

	loaders []*loader

	slides      []Slide
	transitions []*Transition
	box         []float32

	currentIndex int

	currentSlide string
	nextSlide    string

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

	max := float32(1.0)

	verts := []float32{
		-max, max, -0.1, //V0
		-max, -max, -0.1, //V1
		max, -max, -0.1, //V2
		max, max, -0.1, //V3
	}

	uvs := []float32{
		0, 1, //V0 (x,y)
		0, 0, //V1 (x,y)
		1, 0, //V2 (x,y)
		1, 1, //V3 (x,y)
	}

	inds := []int32{
		0, 1, 3,
		3, 1, 2,
	}

	//m := objects.CreateModelWithData(inds, verts)
	m := objects.CreateModelWithDataTexture(inds, verts, uvs)
	entity := objects.MakeEntity(m, mgl32.Vec3{0, 0.0, -0.2}, 0, 0, 0, 1.0)
	s := &Slideshow{
		SlideShowEntity: entity,

		currentIndex:      0,
		currentTransition: 0,
		currentSlide:      "",
		nextSlide:         "",

		delay:        defaultDelay,
		duration:     defaultDuration,
		defaultDelay: defaultDelay,
	}

	return s
}

var previousTime float64
var progress float64

/*
* need to optimized... all var's outside ?!
 */
var currentDuration float64
var pause bool
var delayForTransition float64
var from, to Slide
var transition *Transition
var transitionId int

//Render Render the transitions
func (s *Slideshow) Render(renderer *objects.Renderer, time float64) {
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

			pause = true // pause aktivieren

			if to != nil {
				log.Println("in pause mode, currently shows:", to.GetUid())
			}
			s.setIndexSlides(len(aviableSlides))
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

	from, to = s.getFromAndTo()

	for _, slide := range s.slides {
		slide.Update()
	}

	to.Play()

	transition.Draw(float32(progress),
		from.Display(),
		to.Display(),
		renderer.GetProjection(),
	)
	if err := helper.ErrorCheck(); err != "" {
		log.Println(err)
	}
	//begin render Entity after all shader processing is done!
	renderer.RenderEntity(s.SlideShowEntity, transition.Shader)
	if err := helper.ErrorCheck(); err != "" {
		log.Println(err)
	}

}

var lastIndex int

func (s *Slideshow) setIndexSlides(aviableSlides int) {
	s.currentSlide = s.nextSlide
	current, _ := getSlideByUID(s.slides, s.currentSlide)

	next := current + 1

	if next >= aviableSlides {
		next = 0
	}

	s.nextSlide = s.slides[next].GetUid()

}

func (s *Slideshow) getFromAndTo() (Slide, Slide) {

	idFrom, from := getSlideByUID(s.slides, s.currentSlide)
	if idFrom == -1 {
		//from not found use index zero!
		from = s.slides[0]
		idFrom = 0
		s.currentSlide = from.GetUid()
	}

	idTo, to := getSlideByUID(s.slides, s.nextSlide)
	if idTo == -1 {
		to = s.slides[idFrom+1]
		idTo = idFrom + 1
		s.nextSlide = to.GetUid()
	}

	return from, to
}

func getSlideByUID(slides []Slide, uid string) (int, Slide) {

	for i, s := range slides {
		if s.GetUid() == uid {
			return i, s
		}
	}

	return -1, nil

}

func (s *Slideshow) onlyAviableSlides() []Slide {

	r := make([]Slide, 0)
	for _, s := range s.slides {
		if !s.IsLoading() {
			r = append(r, s)
		}
	}

	return r
}

func (s *Slideshow) SetDuration(duration int32) {
	s.duration = float64(duration)
}

func (s *Slideshow) GetDuration() int32 {
	return int32(s.duration)
}

func (s *Slideshow) SetDelay(delay int32) {
	s.delay = float64(delay)
}

//CleanUP remove all transitions and all slides from memory
func (s *Slideshow) CleanUP() {

	for _, transition := range s.transitions {
		transition.CleanUP()
	}

	for _, slide := range s.slides {
		slide.CleanUP()
	}

	s.SlideShowEntity.Model.Delete()
}
