package graphic

import (
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/slideshow"
)

//Display all Main current states
type Display struct {
	windowWidth  float32
	windowHeight float32

	defaultShader *gfx.Program
	slideshow     *slideshow.Slideshow
}

//InitDisplay initialize a Display object
func InitDisplay(windowWidth int, windowHeight int, defaultDelay, defaultDuration float64) *Display {
	display := &Display{
		windowHeight: float32(windowHeight),
		windowWidth:  float32(windowWidth),
	}

	display.slideshow = slideshow.MakeSlideshow(defaultDelay, defaultDuration)
	display.slideshow.UpdateWindowSize(float32(windowWidth), float32(windowHeight))

	display.slideshow.LoadTransitions("./assets/transitions")

	initFont()
	return display
}

//SetWindowSize set new windows size on resize event
func (d *Display) SetWindowSize(width, height int) {
	d.windowWidth = float32(width)
	d.windowHeight = float32(height)
	d.slideshow.UpdateWindowSize(float32(width), float32(height))
}

//Render make all updates for rendering
func (d *Display) Render(time float64) {
	d.slideshow.Render(time)
}

//Delete unload all data from gpu
func (d *Display) Delete() {
	d.slideshow.Delete()
}
