package graphic

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/mcbernie/myopengl/features"
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

	renderer    *objects.Renderer
	Loader      *objects.Loader
	rawModel    *objects.RawModel
	entity      *objects.Entity
	fpsText     *fonts.GUIText
	font        *fonts.FontType
	laufschrift *features.LaufschriftObject
	objectsList objects.ObjectsList
}

var tex *gfx.Texture

//var setSpecialViewPoint func()

//InitDisplay initialize a Display object
func InitDisplay(windowWidth int, windowHeight int, defaultDelay, defaultDuration float64) *Display {
	d := &Display{
		windowHeight: float32(windowHeight),
		windowWidth:  float32(windowWidth),
	}

	d.Loader = objects.MakeLoader()
	d.renderer = objects.MakeRenderer()
	d.objectsList = objects.CreateObjectList(d.renderer)

	gl.Viewport(0, 0, int32(windowWidth), int32(windowHeight))

	/**
		My Testing Area
	**/
	/*vertices := []float32{
		-0.5, 0.5, 0, //V0
		-0.5, 0.45, 0, //V1
		-0.4, 0.45, 0, //V2
		-0.4, 0.5, 0, //V3
	}
	indicies := []int32{
		0, 1, 3, //Top Left triangle (V0, V1, V3)
		3, 1, 2, //Bottom Right triangle (V3, V1, V2)
	}*/
	//model := objects.CreateModelWithData(indicies, vertices)
	//d.entity = objects.MakeEntity(model, mgl32.Vec3{0, 0, -1.0}, 0, 0, 0, 1.0)

	fonts.InitTextMaster(d.Loader)
	d.font = fonts.MakeFontType(d.Loader.LoadTexture("assets/fonts/verdana.png"), "assets/fonts/verdana.fnt")

	// SlideShowSpecific
	d.slideshow = slideshow.MakeSlideshow(defaultDelay, defaultDuration, d.Loader)
	d.slideshow.LoadTransitions("./assets/transitions", d.renderer.GetProjection())
	elapsed = 0.0

	d.objectsList.AddRenderer(d.slideshow)

	d.laufschrift = features.CreateLaufschrift("Hallo ein sehr kleiner super TEst für mich und mein Freund der .... Hühnerdieb .....!")
	d.objectsList.AddRenderer(d.laufschrift)
	return d
}

//SetWindowSize set new windows size on resize event
func (d *Display) SetWindowSize(width, height int) {
	d.windowWidth = float32(width)
	d.windowHeight = float32(height)
	gl.Viewport(0, 0, int32(width), int32(height))
	d.font.ReplaceMeshCreator()
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

		/*d.fpsText.Remove()
		d.fpsText = fonts.CreateGuiText(fmt.Sprintf("FPS:%.3f", fps), 0.7, d.font, [2]float32{-1.0, 1.0}, 4, false)
		d.fpsText.SetColourRGB(246, 122, 140)

		d.entity.SetColourRGB(255, 0, 10, 80)*/
		d.laufschrift.SetColor(255, 0, 10)
		if fps < 60 {
			//d.fpsText.SetColour(0.8, 0.8, 0.8)
			d.laufschrift.SetColor(200, 200, 200)
			if fps < 30 {
				d.laufschrift.SetColor(200, 130, 130)
				//d.fpsText.SetColour(0.8, 0.5, 0.5)
			}
		}

		frameCount = 0
		lastTime = time
	}

	//d.entity.IncreasePosition(0.09*duration, -0.02*duration, 0)
	//d.slideshow.SlideShowEntity.IncreasePosition(0.05*duration, -0.02*duration, 0)

	//gl.Viewport(0, 0, int32(d.windowWidth), int32(d.windowHeight))
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0.0, 0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	//d.slideshow.Render(time, d.renderer)

	//d.renderer.UseDefaultShader()
	//d.renderer.RenderEntity(d.entity, d.renderer.Shader)

	d.objectsList.Render(time)

}

//Delete unload all data from gpu
func (d *Display) Delete() {
	d.Loader.CleanUP()
	d.slideshow.CleanUP()

}
