package graphic

import (
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/graphic/objects"
	"github.com/mcbernie/myopengl/slideshow"
)

//Display all Main current states
type Display struct {
	windowWidth  float32
	windowHeight float32

	defaultShader *gfx.Program
	slideshow     *slideshow.Slideshow

	renderer *objects.Renderer
	loader   *objects.Loader
	rawModel *objects.RawModel
}

//InitDisplay initialize a Display object
func InitDisplay(windowWidth int, windowHeight int, defaultDelay, defaultDuration float64) *Display {
	d := &Display{
		windowHeight: float32(windowHeight),
		windowWidth:  float32(windowWidth),
	}

	d.loader = objects.MakeLoader()
	d.renderer = objects.MakeRenderer()

	/**
		My Testing Area
	**/
	vertices := []float32{
		-0.5, 0.5, 0, //V0
		-0.5, -0.5, 0, //V1
		0.5, -0.5, 0, //V2
		0.5, 0.5, 0, //V3
	}

	indicies := []int32{
		0, 1, 3, //Top Left triangle (V0, V1, V3)
		3, 1, 2, //Bottom Right triangle (V3, V1, V2)
	}

	d.rawModel = d.loader.LoadToVAO(vertices, indicies)

	/**
		End of My Testing Area
	**/

	// SlideShowSpecific
	d.slideshow = slideshow.MakeSlideshow(defaultDelay, defaultDuration, d.loader)
	d.slideshow.LoadTransitions("./assets/transitions")

	//initFont()
	return d
}

//SetWindowSize set new windows size on resize event
func (d *Display) SetWindowSize(width, height int) {
	d.windowWidth = float32(width)
	d.windowHeight = float32(height)
}

//Render make all updates for rendering
func (d *Display) Render(time float64) {

	d.slideshow.Render(time, d.renderer)
	d.renderer.UseDefaultShader()
	d.renderer.Render(d.rawModel)
}

//Delete unload all data from gpu
func (d *Display) Delete() {
	d.loader.CleanUP()
	d.slideshow.Delete()
}
