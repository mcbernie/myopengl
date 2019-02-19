package objects

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

var globalWindow *glfw.Window

func getDPI(m *glfw.Monitor, w *glfw.Window) float32 {
	wi, _ := m.GetPhysicalSize()
	widthframebuffer, _ := w.GetSize()
	dpi := float32(widthframebuffer) / (float32(wi) / 25.4)
	return dpi
}

//CreateWindow create a glfw window and save pointer to global variable
func CreateWindow(width, height int, title string) (*glfw.Window, error) {
	win, err := glfw.CreateWindow(width, height, title, nil, nil)
	//win.Show()
	//m := win.GetMonitor()
	/*m := glfw.GetPrimaryMonitor()
	win.SetMonitor(m, 0, 0, width, height, 60)
	log.Println("m1:", m)*/

	/*win.SetPosCallback(func(w *glfw.Window, xPos int, yPos int) {
		log.Println("move window: x:", xPos, " y:", yPos)
		mons := glfw.GetMonitors()
		mon0x, mon0y := mons[0].GetPos()
		mon1x, mon1y := mons[1].GetPos()

		log.Println("dpi mon0:", getDPI(mons[0], w))
		log.Println("dpi mon1:", getDPI(mons[1], w))
		log.Println("mon0x:", mon0x, " mon0y:", mon0y)
		log.Println("mon1x:", mon1x, " mon1y:", mon1y)
	})*/

	globalWindow = win

	return globalWindow, err
}

//GetWindow returns a global static glfw.Window pointer
func GetWindow() *glfw.Window {
	return globalWindow
}
