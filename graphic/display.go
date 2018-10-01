package graphic

import (
	"log"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/mcbernie/myopengl/graphic/fonts"
	"github.com/mcbernie/myopengl/graphic/helper"
	"github.com/mcbernie/myopengl/graphic/objects"
	"github.com/mcbernie/myopengl/slideshow"
)

//Display all Main current states
type Display struct {
	windowWidth  float32
	windowHeight float32

	defaultShader *objects.Program
	slideshow     *slideshow.Slideshow

	renderer    *objects.Renderer
	Loader      *objects.Loader
	rawModel    *objects.RawModel
	entity      *objects.Entity
	fpsText     *fonts.GUIText
	font        *fonts.FontType
	laufschrift *objects.LaufschriftObject
	objectsList objects.ObjectsList
}

var tex *objects.Texture

//InitDisplay initialize a Display object
func InitDisplay(windowWidth int, windowHeight int, defaultDelay, defaultDuration float64) *Display {
	d := &Display{
		windowHeight: float32(windowHeight),
		windowWidth:  float32(windowWidth),
	}

	//Setup Scoping
	helper.InitScoping()

	//Init BlendFunction
	d.EnableBlendFunction()

	log.Print("init loader, renderer and createObjectList...")
	d.Loader = objects.MakeLoader()
	d.renderer = objects.MakeRenderer()
	d.objectsList = objects.CreateObjectList(d.renderer)

	log.Print("init Viewport")
	helper.Viewport(0, 0, int32(windowWidth), int32(windowHeight))

	log.Print("init fonts")
	fonts.InitTextMaster(d.Loader)
	d.font = fonts.MakeFontType(d.Loader.LoadTexture("assets/fonts/verdana.png"), "assets/fonts/verdana.fnt")

	log.Print("init slideshow")
	// SlideShowSpecific
	d.slideshow = slideshow.MakeSlideshow(defaultDelay, defaultDuration, d.Loader)
	d.slideshow.LoadTransitions("./assets/transitions", d.renderer.GetProjection())
	elapsed = 0.0
	d.objectsList.AddRenderer(d.slideshow)

	/*go func() {
		time.Sleep(5 * time.Second)
		log.Println("Test laufschirft replacing")
		d.laufschrift.SetTextSafe("Hallo Mallo")
	}()*/

	//d.fpsText = fonts.CreateGuiText("init", 0.7, d.font, [2]float32{-1.0, 1.0}, 4, false)
	d.laufschrift = objects.CreateLaufschrift(
		"Ganz kurzer Text!8n ug ztg ztgvi gviv izvizviztvizviztgfiufiztfz",
		-0.8, -0.8, 1.6, 0.2)
	d.objectsList.AddRenderer(d.laufschrift)

	return d
}

func (d *Display) SetProjection() {
	helper.MatrixMode(helper.GlProjection)
	helper.LoadIdentity()
	helper.Viewport(0, 0, int32(d.windowWidth), int32(d.windowHeight))
	helper.Ortho(0.0, float64(d.windowWidth), 0.0, float64(d.windowHeight), 0.0, 1.0)
	helper.MatrixMode(helper.GlModelView)
	helper.LoadIdentity()
}

func (d *Display) SetKeyCallback(key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	step := 0.01

	if mods == glfw.ModShift {
		step = 0.1
	}

	if key == glfw.KeyDown {
		d.laufschrift.SetPosition(0, -float32(step))
	}

	if key == glfw.KeyUp {
		d.laufschrift.SetPosition(0, float32(step))
	}

	if key == glfw.KeyLeft {
		d.laufschrift.SetPosition(-float32(step), 0)
	}

	if key == glfw.KeyRight {
		d.laufschrift.SetPosition(float32(step), 0)
	}

}

//EnableBlendFunction are required to display alpha in fragmentshaders / ex. Laufschrift
func (d *Display) EnableBlendFunction() {
	helper.Enable(helper.GlBlend)
	helper.BlendFunc(helper.GlSrcAlpha, helper.GlOneMinusSrcAlpha)
}

//SetWindowSize set new windows size on resize event
func (d *Display) SetWindowSize(width, height int) {
	d.windowWidth = float32(width)
	d.windowHeight = float32(height)

	d.SetProjection()
	d.font.ReplaceMeshCreator()
}

var elapsed float64
var lastTime float64
var frameCount int

//Render make all updates for rendering
func (d *Display) Render(time float64) {
	//Run all functions...
	helper.RunFunctions()

	delta := time - lastTime
	frameCount++
	if delta >= 1 {
		fps := float64(frameCount) / delta

		/*d.fpsText.Remove()
		d.fpsText = fonts.CreateGuiText(fmt.Sprintf("FPS:%.3f", fps), 0.7, d.font, [2]float32{-1.0, 1.0}, 4, false)
		d.fpsText.SetColourRGB(246, 122, 140)*/

		/*d.entity.SetColourRGB(255, 0, 10, 80)*/
		d.laufschrift.SetColor(255, 0, 10)
		if fps < 60 {
			//d.fpsText.SetColour(0.8, 0.8, 0.8)
			//d.laufschrift.SetColor(200, 200, 200)
			if fps < 30 {
				//d.laufschrift.SetColor(200, 130, 130)
				//d.fpsText.SetColour(0.8, 0.5, 0.5)
			}
		}

		frameCount = 0
		lastTime = time
	}

	//helper.Enable(helper.GlDepthTest)
	helper.ClearColor(0.0, 0.5, 1.0, 1.0)
	helper.Clear(helper.GlColorBufferBit | helper.GlDepthBufferBit)

	d.objectsList.Render(time)
}

//Delete unload all data from gpu
func (d *Display) Delete() {
	d.Loader.CleanUP()
	d.slideshow.CleanUP()
}
