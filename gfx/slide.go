package gfx

import (
	"image"
)

type Slide interface {
	GetDelay() float64
	GetUid() string
	IsLoading() bool

	Display() *Texture
	Update()
	Play()
	SetFrame(img image.Image)
	Delete()
	CleanUP()

	GoToNextSlide(currentDuration float64) bool

	BackgroundThread()
}
