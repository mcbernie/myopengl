package graphic

import (
	"log"

	"github.com/go-gl/glfw/v3.2/glfw"
	//"github.com/mcbernie/myopengl/graphic/gui"
	"github.com/mcbernie/myopengl/graphic/helper"
	"github.com/mcbernie/myopengl/graphic/objects"
	"github.com/mcbernie/myopengl/slideshow"
)

//Display all Main current states
type Display struct {
	windowWidth  float32
	windowHeight float32
	fbWidth      float32
	fbHeiht      float32

	window *glfw.Window

	// Memory, Object and Render Managment
	renderer    *objects.Renderer
	Loader      *objects.Loader
	objectsList objects.ObjectsList

	// GFX Systems
	slideshow   *slideshow.Slideshow
	laufschrift *objects.LaufschriftObject
	//gui         *gui.GuiSystem
}

var tex *objects.Texture

//InitDisplay initialize a Display object
func InitDisplay(window *glfw.Window, defaultDelay, defaultDuration float64) *Display {

	windowWidth, windowHeight := window.GetFramebufferSize()
	d := &Display{
		windowWidth:  float32(windowWidth),
		windowHeight: float32(windowHeight),
		window:       window,
	}

	d.GlfwCallback(window)

	//Setup Scoping
	helper.InitScoping()

	//Init BlendFunction
	d.EnableBlendFunction()

	log.Print("init loader, renderer and createObjectList...")
	d.Loader = objects.MakeLoader()
	d.renderer = objects.MakeRenderer()
	d.objectsList = objects.CreateObjectList(d.renderer)

	log.Print("init slideshow")
	d.slideshow = slideshow.MakeSlideshow(defaultDelay, defaultDuration, d.Loader)
	d.slideshow.LoadTransitions("./assets/transitions", d.renderer.GetProjection())
	elapsed = 0.0
	d.objectsList.AddRenderer(d.slideshow)

	/*d.laufschrift = objects.CreateLaufschrift(
	"Ganz kurzer Text!8n ug ztg ztgvi gviv izvizviztvizviztgfiufiztfz",
	-0.8, -0.8, 1.6, 0.2)*/
	//d.objectsList.AddRenderer(d.laufschrift)

	//d.gui = gui.CreateGui(d.window)
	//d.gui.SetDuration(d.slideshow.GetDuration())
	//d.objectsList.AddRenderer(d.gui)

	return d
}

func (d *Display) SetKeyCallback(key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	//step := 0.01

	/*if mods == glfw.ModShift {
		step = 0.1
	}*/

	if key == glfw.KeyDown {
		//d.laufschrift.SetPosition(0, -float32(step))
	}

	if key == glfw.KeyUp {
		//d.laufschrift.SetPosition(0, float32(step))
	}

	if key == glfw.KeyLeft {
		//d.laufschrift.SetPosition(-float32(step), 0)
	}

	if key == glfw.KeyRight {
		//d.laufschrift.SetPosition(float32(step), 0)
	}

	if key == glfw.KeyS && action == glfw.Press {
		//d.gui.ToggleVisible()
	}

}

//EnableBlendFunction are required to display alpha in fragmentshaders / ex. Laufschrift
func (d *Display) EnableBlendFunction() {
	helper.Enable(helper.GlBlend)
	helper.BlendFunc(helper.GlSrcAlpha, helper.GlOneMinusSrcAlpha)
}

//SetWindowSize set new windows size on resize event
func (d *Display) GlfwCallback(w *glfw.Window) {
	fbWidth, fbHeight := w.GetFramebufferSize()
	winWidth, winHeight := w.GetSize()

	d.windowWidth = float32(winWidth)
	d.windowHeight = float32(winHeight)
	d.fbWidth = float32(fbWidth)
	d.fbHeiht = float32(fbHeight)

	d.SetProjection()
}

func (d *Display) SetProjection() {
	helper.Viewport(0, 0, int32(d.fbWidth), int32(d.fbHeiht))
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
		if fps < 60 {
			if fps < 30 {
			}
		}
		frameCount = 0
		lastTime = time
	}

	helper.Viewport(0, 0, int32(d.fbWidth), int32(d.fbHeiht))

	helper.ClearColor(0.0, 0.2, 1.0, 1.0)
	helper.Clear(helper.GlColorBufferBit)
	helper.RunFunctions()

	/*if pause := d.gui.GetPause(); pause > 0 && pause < 50 {
		d.slideshow.SetDelay(pause)
	}
	if duration := d.gui.GetDuration(); duration > 0 && duration < 50 {
		d.slideshow.SetDuration(duration)
	}*/

	//Run object list
	d.objectsList.Render(time)

}

//Delete unload all data from gpu
func (d *Display) Delete() {
	d.Loader.CleanUP()
	d.objectsList.Delete()
}
