package graphic

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/glHelper"
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
	loader      *objects.Loader
	rawModel    *objects.RawModel
	entity      *objects.Entity
	fpsText     *fonts.GUIText
	font        *fonts.FontType
	laufschrift *fonts.GUIText
}

var rendTexture uint32
var rendFrameBuff uint32
var laufschriftModel *objects.RawModel
var laufschriftEntity *objects.Entity
var lShader *gfx.Program
var setSpecialViewPoint func()

//InitDisplay initialize a Display object
func InitDisplay(windowWidth int, windowHeight int, defaultDelay, defaultDuration float64) *Display {
	d := &Display{
		windowHeight: float32(windowHeight),
		windowWidth:  float32(windowWidth),
	}

	d.loader = objects.MakeLoader()
	d.renderer = objects.MakeRenderer()

	//gl.Viewport(1000, 500, 40000, int32(windowHeight))
	gl.Viewport(0, 0, int32(windowWidth), int32(windowHeight))

	/*gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-10, 10, -10, 10, -10, 10)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()*/
	/**
		My Testing Area
	**/
	vertices := []float32{
		-0.5, 0.5, 0, //V0
		-0.5, 0.45, 0, //V1
		-0.4, 0.45, 0, //V2
		-0.4, 0.5, 0, //V3
	}
	indicies := []int32{
		0, 1, 3, //Top Left triangle (V0, V1, V3)
		3, 1, 2, //Bottom Right triangle (V3, V1, V2)
	}

	rawModel := d.loader.LoadToVAO(vertices, indicies)
	d.entity = objects.MakeEntity(rawModel, mgl32.Vec3{0, 0, -1.0}, 0, 0, 0, 1.0)

	// Simple QUAD <--> for laufschrift
	/*simpleQuad := []float32{
		-1.0, 1.0, 0, //V0
		-1.0, -1.0, 0, //V1
		-1.0, 1.0, 0, //V2
		-1.0, -1.0, 0, //V3
	}*/

	bottomMax := float32(-0.8)
	//bottomMax := float32(1.0)
	simpleQuad := []float32{
		-1.0, bottomMax, 0, //V0
		-1.0, -1.0, 0, //V1
		1.0, -1.0, 0, //V2
		1.0, bottomMax, 0, //V3
	}

	/*simpleTexture := []float32{
		0.1, 0.4,
		0.1, 0.2,
		0.8, 0.2,
	}*/

	simpleTexture := []float32{
		0, 1, //V0
		0, 0, //V1
		1, 0, //V2
		1, 1, //V3
	}

	simpleIndicies := []int32{
		0, 1, 3, //Top Left triangle (V0, V1, V3)
		3, 1, 2, //Bottom Right triangle (V3, V1, V2)
	}
	laufschriftModel = d.loader.LoadToVAOWithTexCoords(simpleQuad, simpleIndicies, simpleTexture)
	//laufschriftModel = d.loader.LoadToVAO(simpleQuad, simpleIndicies)
	laufschriftEntity = objects.MakeEntity(laufschriftModel, mgl32.Vec3{0, 0, 0}, 0, 0, 0, 1.0)

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

	//laufschrift test..
	d.laufschrift = fonts.CreateGuiText("Hallo I'bims 1 Laufschrift mit ganz viel Text!", 3, d.font, [2]float32{-1.0, -0.8}, -1, false)
	d.laufschrift.SetColour(1, 1, 1)

	tsize := d.laufschrift.GetSize()
	lwidth := tsize[0]
	lheight := tsize[1]

	pixelWidth := (float32(windowWidth) / 2.0) * lwidth
	pixelHeight := (float32(windowHeight) / 2.0) * lheight

	log.Println("Screen: width:", windowWidth, " height:", windowHeight)
	log.Println(" Pixel: width:", pixelWidth, " height:", pixelHeight)
	log.Println("opengl: width:", lwidth, " height:", lheight)

	// SlideShowSpecific
	d.slideshow = slideshow.MakeSlideshow(defaultDelay, defaultDuration, d.loader)
	d.slideshow.LoadTransitions("./assets/transitions", d.renderer.GetProjection())
	elapsed = 0.0

	setSpecialViewPoint = func() {
		gl.Viewport(0, 1100, 2000, 100)
	}

	/* Generate Framebuffer for my experiment */
	gl.GenFramebuffers(1, &rendFrameBuff)
	gl.BindFramebuffer(gl.FRAMEBUFFER, rendFrameBuff)

	gl.GenTextures(1, &rendTexture)
	glHelper.ActiveTexture(glHelper.GlTexture0)
	gl.BindTexture(gl.TEXTURE_2D, rendTexture)

	//gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, 2000, 100, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB, 4000, 100, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE|gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE|gl.REPEAT)

	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, rendTexture, 0)
	gl.DrawBuffer(gl.COLOR_ATTACHMENT0)
	/* Framebuffer experiment.... */

	var err error
	lShader, err = createLaufschriftShader()
	if err != nil {
		log.Println("create Laufschrift shader error:", err)
	}

	lShader.Use()
	lShader.AddUniform("projectionMatrix")
	glHelper.UniformMatrix4(lShader.GetUniform("projectionMatrix"), d.renderer.GetProjection())
	lShader.UnUse()
	/*dbuffers :=
	gl.drawbu*/

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.BindTexture(gl.TEXTURE_2D, 0)
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

		d.fpsText.Remove()
		d.fpsText = fonts.CreateGuiText(fmt.Sprintf("FPS:%.3f", fps), 0.7, d.font, [2]float32{-1.0, 1.0}, 4, false)
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
	//log.Println("rendFrameBuff:", rendFrameBuff)
	gl.BindFramebuffer(gl.FRAMEBUFFER, rendFrameBuff)
	gl.Disable(gl.DEPTH_TEST)
	gl.ClearColor(1.0, 0.2, 1.0, 0.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.Viewport(0, 0, 4000, int32(d.windowHeight))
	fonts.TextMaster.Render()
	/*d.renderer.UseDefaultShader()

	d.renderer.RenderEntity(d.entity, d.renderer.Shader)*/
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)

	//d.entity.IncreasePosition(0.09*duration, -0.02*duration, 0)
	//d.slideshow.SlideShowEntity.IncreasePosition(0.05*duration, -0.02*duration, 0)

	gl.Viewport(0, 0, int32(d.windowWidth), int32(d.windowHeight))
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0.5, 1.0, 1.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	d.slideshow.Render(time, d.renderer)

	//d.renderer.UseDefaultShader()

	//d.renderer.RenderEntity(d.entity, d.renderer.Shader)

	lShader.Use()
	//d.renderer.UseDefaultShader()
	glHelper.ActiveTexture(glHelper.GlTexture0)
	glHelper.BindTexture(gl.TEXTURE_2D, rendTexture)
	glHelper.Uniform1i(lShader.GetUniform("renderedTexture"), 0)
	glHelper.Uniform1f(lShader.GetUniform("time"), float32(time))
	d.renderer.RenderEntity(laufschriftEntity, lShader)
	lShader.UnUse()

	//glHelper.BindTexture(gl.TEXTURE_2D, 0)
	/*gl.Disable(gl.DEPTH_TEST)
	glHelper.UseProgram(0)
	gl.Color4f(1.0, 1.0, 1.0, 0.9)
	d.fonts[10].Printf(0, 0, "Hallo Test")*/

}

//Delete unload all data from gpu
func (d *Display) Delete() {
	d.loader.CleanUP()
	d.slideshow.CleanUP()

}

func createLaufschriftShader() (*gfx.Program, error) {
	vert, err := gfx.NewShaderFromFile("assets/shaders/laufschrift.vert", gfx.VertexShaderType)
	if err != nil {
		return nil, err
	}

	frag, err := gfx.NewShaderFromFile("assets/shaders/laufschrift.frag", gfx.FragmentShaderType)
	if err != nil {
		return nil, err
	}

	shader, err := gfx.NewProgram(vert, frag)

	if err != nil {
		return nil, err
	}

	shader.Use()
	shader.AddUniform("renderedTexture")
	shader.AddUniform("projectionMatrix")
	shader.AddUniform("transformationMatrix")
	shader.AddUniform("time")
	shader.AddAttribute("position")
	shader.AddAttribute("texCoord")
	shader.BindAttribute(0, "position")
	shader.BindAttribute(1, "texCoord")
	shader.UnUse()

	return shader, nil
}
