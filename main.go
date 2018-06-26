package main

import (
	"log"
	"runtime"
	"time"

	//"github.com/go-gl/gl/v4.1-core/gl" // OR:

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/glHelper"
	"github.com/mcbernie/myopengl/graphic"
)

const windowWidth = 800
const windowHeight = 600

const delay float64 = 5.0    //8.0
const duration float64 = 1.5 // 2.0

func init() {
	// GLFW event handling must be run on the main OS thread
	runtime.LockOSThread()
}

func main() {

	//Setup Scoping
	glHelper.InitScoping()

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to inifitialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	//glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	//glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	//glfw.CreateWindow(windowWidth, windowHeight, "basic slideshow", nil, nil)
	window, err := gfx.CreateWindow(windowWidth, windowHeight, "basic slideshow")
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := glHelper.Init(); err != nil {
		panic(err)
	}

	log.Println("OpenGL Version:", gl.GoStr(gl.GetString(gl.VERSION)))
	log.Println("OpenGL Shading Version:", gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION)))

	window.SetKeyCallback(keyCallback)

	err = programLoop(window)
	if err != nil {
		log.Fatal(err)
	}
}

func programLoop(window *glfw.Window) error {

	width, height := window.GetSize()

	display := graphic.InitDisplay(width, height, delay, duration)
	display.LoadImagesFromPath("./assets/images")

	window.SetSizeCallback(func(w *glfw.Window, width int, height int) {
		display.SetWindowSize(width, height)
	})

	go func() {

		time.Sleep(10 * time.Second)
		log.Println("now begin loading a image...")
		display.LoadRemoteImage("http://wacogmbh.de:3999/index.php?m=fb&o=image&name=med_1421768202_45415200.jpg", "lkih76555")

	}()

	go func() {
		time.Sleep(16 * time.Second)
		log.Println("now begin loading 2. a image...")
		display.LoadRemoteImage("http://wacogmbh.de:3999/index.php?m=fb&o=image&name=med_1275283013_60102000.jpg", "asduhfudh")
	}()

	/*go func() {
		time.Sleep(21 * time.Second)
		log.Println("Now testing the removeing of an texture")
		display.RemoveSlide("lkih76555")
	}()*/

	/*go func() {
		time.Sleep(23 * time.Second)
		log.Println("now begin loading 2. a video...")
		display.LoadVideoAtRuntime("assets/video/tr5_event_bally.mp4", "tr5_bally_event")
	}()*/

	//display.LoadVideo("assets/video/big_buck_bunny.mp4", "Big_Buck_Bunny")

	display.LoadVideo("assets/video/tr5_event_bally.mp4", "Big_Buck_Bunny")
	defer display.Delete()

	//gl.Enable(gl.DEPTH_TEST)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	for !window.ShouldClose() {
		//scoping...
		glHelper.RunFunctions()

		// poll events and call their registered callbacks
		glfw.PollEvents()

		// background color
		glHelper.ClearColor(0.0, 0.0, 0.0, 1.0)
		glHelper.Clear(glHelper.GlColorBufferBit)
		//gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		display.Render(glfw.GetTime())

		window.SwapBuffers()
	}

	return nil
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action,
	mods glfw.ModifierKey) {

	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}
