package slideshow

import (
	"image"

	"github.com/mcbernie/myopengl/graphic/objects"
)

type Slide interface {
	GetDelay() float64
	GetUid() string
	IsLoading() bool

	Display() *objects.Texture
	Update()
	Play()
	SetFrame(img image.Image)
	Delete()
	CleanUP()

	GoToNextSlide(currentDuration float64) bool

	BackgroundThread()
}
