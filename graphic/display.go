package graphic

import (
	"log"

	//"github.com/go-gl/gl/v4.1-core/gl" // OR:
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/slideshow"
)

//Display all Main current states
type Display struct {
	windowWidth  float32
	windowHeight float32

	defaultShader *gfx.Program
	gui           *GUI
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

func (d *Display) setupVbp() {
	//var VBO uint32
	/*gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(d.box)*4, gl.Ptr(d.box), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))*/
}

//Render make all updates for rendering
func (d *Display) Render(time float64) {

	// wait for incoming gl commands and run...

	d.slideshow.Render(time)
}

//GUI show struct
type GUI struct {
	width  float64
	height float64

	x float64
	y float64
}

func (d *Display) MakeGUI(x float64, y float64, width float64, height float64) *GUI {

	calcX := (x / float64(d.windowWidth))
	calcY := (y * -1 / float64(d.windowHeight)) - height

	log.Println(calcY)
	//calcX := xPointZero
	//calcY := yPointZero

	return &GUI{
		width:  width / float64(d.windowWidth),
		height: height / float64(d.windowHeight),
		x:      calcX,
		y:      calcY,
	}

}

func (g *GUI) draw() {
	//d.defaultShader.Use()
	gl.UseProgram(0)
	//d.drawString(0.0, 0.0, "Hallo Mallo Welt und Geld", 12)

	gl.Disable(gl.TEXTURE_2D)

	gl.Disable(gl.LIGHTING)
	gl.Color3d(1.0, 1.0, 1.0)
	gl.Begin(gl.QUADS)
	gl.Vertex3d(g.x, g.y+g.height, 0)
	gl.Vertex3d(g.x+g.width, g.y+g.height, 0)
	gl.Vertex3d(g.x+g.width, g.y, 0)
	gl.Vertex3d(g.x, g.y, 0)
	gl.End()
	gl.Enable(gl.TEXTURE_2D)
}

//Delete unload all data from gpu
func (d *Display) Delete() {
	d.slideshow.Delete()
}
