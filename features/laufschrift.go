package features

import (
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/glHelper"
	"github.com/mcbernie/myopengl/graphic/objects"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type LaufschriftObject struct {
	image   image.Image
	text    string
	texture *gfx.Texture
	shader  *gfx.Program
	entity  *objects.Entity
	width   float32
	height  float32
	x       float32
	y       float32
}

func (l *LaufschriftObject) Render(r *objects.Renderer, time float64) {
	l.texture.Bind(0)
	l.shader.Use()

	glHelper.UniformMatrix4(l.shader.GetUniform("projectionMatrix"), r.GetProjection())
	glHelper.Uniform1i(l.shader.GetUniform("renderedTexture"), 0)
	glHelper.Uniform1f(l.shader.GetUniform("time"), float32(time/2.5))

	r.RenderEntity(l.entity, l.shader)

	l.shader.UnUse()
	l.texture.UnBind()
}

func (l *LaufschriftObject) SetText(text string) {

	l.text = text
	l.image = addLabel(text)
	l.texture.SetDefaultImage(l.image)

}

func (l *LaufschriftObject) SetColor(r, g, b int) {
	l.entity.SetColourRGB(r, g, b, 1)
}

func (l *LaufschriftObject) UpdateEntity() {
	l.width = 2.0
	l.height = 0.2

	iheight := float64(l.image.Bounds().Dy()) // height of laufschrift texture font
	iwidth := float64(l.image.Bounds().Dx())  // width of laufschrift texture font

	ratioOfiWidth := iwidth / iheight // to hold the correct ratio between height and width, width needs to bee the x factor of it
	modelWidth := float64(l.width)    // 2.0 param 2.0 as screenwidth

	maxSizeWidth := ratioOfiWidth // 100% == uv.x = 1
	pValue := maxSizeWidth / modelWidth

	log.Println("Image width:", iwidth, " Image height:", iheight)
	log.Println("Image ration:", ratioOfiWidth)
	log.Println("maxSizeWidth:", maxSizeWidth)
	log.Println("pValue:", pValue)

	newuvx := float32((1.0 / ratioOfiWidth) * pValue)
	log.Println("newuvx:", newuvx)
	bottomMax := float32(-0.8)
	//bottomMax := float32(1.0)
	simpleQuad := []float32{
		-1.0, bottomMax, 0, //V0
		-1.0, -1.0, 0, //V1
		1.0, -1.0, 0, //V2
		1.0, bottomMax, 0, //V3
	}

	simpleTexture := []float32{
		0, 1, //V0 (x,y)
		0, 0, //V1 (x,y)
		newuvx, 0, //V2 (x,y)
		newuvx, 1, //V3 (x,y)
	}

	simpleIndicies := []int32{
		0, 1, 3, //Top Left triangle (V0, V1, V3)
		3, 1, 2, //Bottom Right triangle (V3, V1, V2)
	}

	model := objects.CreateModelWithDataTexture(simpleIndicies, simpleQuad, simpleTexture)
	l.entity = objects.MakeEntity(model, mgl32.Vec3{0, 0, 0}, 0, 0, 0, 1.0)

}

func CreateLaufschrift(text string) *LaufschriftObject {

	lShader, err := createLaufschriftShader()
	if err != nil {
		log.Println("create Laufschrift shader error:", err)
	}

	l := LaufschriftObject{
		texture: gfx.NewTextureDefault(),
		shader:  lShader,
	}
	l.SetText(text)

	l.texture.Bind(0)
	glHelper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	glHelper.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, gl.CLAMP_TO_EDGE)
	l.texture.UnBind()

	l.UpdateEntity()

	return &l
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

/*func createLaufschrift(d *Display, text string, width, height float32) {

	d.objectsList.AddRenderer(&o)
}*/

func addLabel(label string) image.Image {

	// Read the font data.
	fontBytes, err := ioutil.ReadFile("assets/fonts/luximr.ttf")
	if err != nil {
		log.Println(err)
		return nil
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return nil
	}

	// Variables
	fontSize := float64(120.0)
	dpi := float64(72)
	lineSpacing := 1.0

	face := truetype.NewFace(f, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingNone})

	fg, bg := image.Black, image.Transparent

	d := &font.Drawer{Dst: nil, Src: fg, Face: face}

	var width, height int
	metrics := face.Metrics()
	lineHeight := (metrics.Ascent + metrics.Descent).Ceil()
	lineGap := int((lineSpacing - float64(1)) * float64(lineHeight))

	// if lines then here
	lines := strings.Split(label, "\n")
	for i, s := range lines {
		d.Dot = fixed.P(0, height)
		lineWidth := d.MeasureString(s).Ceil()
		if lineWidth > width {
			width = lineWidth
		}
		height += lineHeight
		if i > 1 {
			height += lineGap
		}
	}

	// now i have width and height...
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	//now draw text on image
	d = &font.Drawer{Dst: rgba, Src: fg, Face: face}
	// i need metrices and py
	// and lineHeigt, lineGap and the lines splittet from string label by \n
	py := 0 + metrics.Ascent.Round() // 0 should be y if function call

	for i, s := range lines {
		d.Dot = fixed.P(0, py) // 0 should x on function call
		d.DrawString(s)
		py += lineHeight
		if i > 1 {
			py += lineGap
		}
	}

	/*outputFile, err := os.Create("test1.png")
	if err != nil {
		// Handle error
		log.Println("ERROR:", err)
	}

	png.Encode(outputFile, rgba)
	outputFile.Close()*/

	return rgba
}
