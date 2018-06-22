package graphic

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/graphic/fonts"
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
	entity   *objects.Entity
	fpsText  *fonts.GUIText
	font     *fonts.FontType
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
		-0.5, 0.45, 0, //V1
		-0.4, 0.45, 0, //V2
		-0.4, 0.5, 0, //V3
	}

	// Important->
	/*
		v0 oben links
		v1 unten links
		v2 unten rechts
		v3 oben rechts
	*/

	indicies := []int32{
		0, 1, 3, //Top Left triangle (V0, V1, V3)
		3, 1, 2, //Bottom Right triangle (V3, V1, V2)
	}

	rawModel := d.loader.LoadToVAO(vertices, indicies)

	d.entity = objects.MakeEntity(rawModel, mgl32.Vec3{0, 0, -1.0}, 0, 0, 0, 1.0)

	// --->>>
	fonts.InitTextMaster(d.loader)
	d.font = fonts.MakeFontType(d.loader.LoadTexture("assets/fonts/verdana.png"), "assets/fonts/verdana.fnt")
	/*text := fonts.CreateGuiText("", 1, d.font, [2]float32{0.0, 0.0}, 4, false)
	text.SetColour(1.0, 1.0, 1)
	d.fpsText = text*/

	// <<<----

	/**
		End of My Testing Area
	**/

	// SlideShowSpecific
	d.slideshow = slideshow.MakeSlideshow(defaultDelay, defaultDuration, d.loader)
	d.slideshow.LoadTransitions("./assets/transitions", d.renderer.GetProjection())
	elapsed = 0.0
	return d
}

//SetWindowSize set new windows size on resize event
func (d *Display) SetWindowSize(width, height int) {
	d.windowWidth = float32(width)
	d.windowHeight = float32(height)
}

var elapsed float64
var lastTime float64
var frameCount int

//Render make all updates for rendering
func (d *Display) Render(time float64) {
	delta := time - lastTime
	frameCount++

	if delta >= 1 {
		fps := float64(frameCount) / delta

		d.fpsText.Remove()
		d.fpsText = fonts.CreateGuiText(fmt.Sprintf("FPS:%.3f", fps), 0.7, d.font, [2]float32{0.0, 0.0}, 4, false)
		d.fpsText.SetColourRGB(246, 122, 140)

		d.entity.SetColourRGB(255, 0, 10, 80)
		if fps < 60 {
			d.fpsText.SetColour(0.8, 0.8, 0.8)
			if fps < 30 {
				d.fpsText.SetColour(0.8, 0.5, 0.5)
			}
		}

		frameCount = 0
		lastTime = time
	}

	//d.entity.IncreasePosition(0.09*duration, -0.02*duration, 0)
	//d.slideshow.SlideShowEntity.IncreasePosition(0.05*duration, -0.02*duration, 0)

	d.slideshow.Render(time, d.renderer)

	d.renderer.UseDefaultShader()

	d.renderer.RenderEntity(d.entity, d.renderer.Shader)

	/*gl.Disable(gl.DEPTH_TEST)
	glHelper.UseProgram(0)
	gl.Color4f(1.0, 1.0, 1.0, 0.9)
	d.fonts[10].Printf(0, 0, "Hallo Test")*/

	fonts.TextMaster.Render()

}

//Delete unload all data from gpu
func (d *Display) Delete() {
	d.loader.CleanUP()
	d.slideshow.Delete()

}
