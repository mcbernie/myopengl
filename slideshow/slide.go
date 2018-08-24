package slideshow

import (
	"image"

	"github.com/mcbernie/myopengl/gfx"
)

type Slide interface {
	GetDelay() float64
	GetUid() string
	IsLoading() bool

	Display() *gfx.Texture
	Update()
	Play()
	SetFrame(img image.Image)
	Delete()
	CleanUP()

	GoToNextSlide(currentDuration float64) bool

	BackgroundThread()
}
