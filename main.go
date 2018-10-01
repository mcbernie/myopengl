package main

import (
	"log"
	"runtime"
	"time"

	//"github.com/go-gl/gl/v4.1-core/gl" // OR:
	"github.com/go-gl/gl/v2.1/gl"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/mcbernie/myopengl/graphic"
	"github.com/mcbernie/myopengl/graphic/helper"
	"github.com/mcbernie/myopengl/graphic/objects"

	"github.com/pkg/profile"
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
	//cpu profiling
	//defer profile.Start().Stop()
	//mem profiling
	defer profile.Start(profile.MemProfile).Stop()

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to inifitialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := objects.CreateWindow(windowWidth, windowHeight, "SlideShow Test")
	if err != nil {
		panic(err)
	}
	//m := window.GetMonitor()
	//log.Println("m:", m)

	if err := helper.Init(); err != nil {
		panic(err)
	}

	log.Println("OpenGL Version:", gl.GoStr(gl.GetString(gl.VERSION)))
	log.Println("OpenGL Shading Version:", gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION)))

	window.SetKeyCallback(keyCallback)
	window.MakeContextCurrent()

	err = programLoop(window)
	if err != nil {
		log.Fatal(err)
	}
}

func programLoop(window *glfw.Window) error {

	width, height := window.GetFramebufferSize()
	display := graphic.InitDisplay(width, height, delay, duration)
	window.SetFramebufferSizeCallback(func(w *glfw.Window, width, height int) {
		display.SetWindowSize(width, height)
	})

	display.LoadImagesFromPath("./assets/images")

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
		display.LoadVideoAtRuntime("assets/video/tr5_event_bally.mp4", "tr5_bally_event", 10)
	}()*/

	//mw, mh := glfw.GetMonitors()[1].GetPhysicalSize()
	//log.Println("Monitors: ", mw, " h:", mh)

	defer display.Delete()
	display.SetProjection()

	mac_moved := false

	for !window.ShouldClose() {
		display.Render(glfw.GetTime())
		window.SwapBuffers()
		glfw.PollEvents()

		if mac_moved == false {
			x, y := window.GetPos()

			window.SetPos(x+1, y)
			mac_moved = true
		}
	}

	return nil
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action,
	mods glfw.ModifierKey) {

	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}
