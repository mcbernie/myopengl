package graphic

import (
	"log"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/graphic/fonts"
	"github.com/mcbernie/myopengl/graphic/gltext"
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

	fonts [16]*gltext.Font
}

//InitDisplay initialize a Display object
func InitDisplay(windowWidth int, windowHeight int, defaultDelay, defaultDuration float64) *Display {
	d := &Display{
		windowHeight: float32(windowHeight),
		windowWidth:  float32(windowWidth),
	}

	d.loader = objects.MakeLoader()
	d.renderer = objects.MakeRenderer()

	//d.InitFont()
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

	rawModel := d.loader.LoadToVAO(vertices, indicies)

	d.entity = objects.MakeEntity(rawModel, mgl32.Vec3{-0.5, 0.5, -0.1}, 0, 0, 0, 1.0)

	// --->>>
	log.Println("InitTextMAster-->")
	fonts.InitTextMaster(d.loader)

	log.Println("Add Font->")
	font := fonts.MakeFontType(d.loader.LoadTexture("assets/images/index.php-3.jpeg"), "assets/fonts/verdana.fnt")

	log.Println("CreateGUIText->")
	text := fonts.CreateGuiText("Dies ist ein Test", 3, font, [2]float32{0, 0}, 1, true)
	text.SetColour(1, 0, 0)

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

//Render make all updates for rendering
func (d *Display) Render(time float64) {
	/*duration := float32(time - elapsed)
	elapsed = time
	d.entity.IncreasePosition(0.09*duration, -0.02*duration, 0)
	d.slideshow.SlideShowEntity.IncreasePosition(0.05*duration, -0.02*duration, 0)*/

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

	for i := range d.fonts {
		d.fonts[i].Release()
	}
}
