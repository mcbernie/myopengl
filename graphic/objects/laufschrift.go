package objects

import (
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/mcbernie/myopengl/graphic/helper"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type LaufschriftObject struct {
	image         image.Image
	text          string
	texture       *Texture
	shader        *Program
	entity        *Entity
	width         float32
	height        float32
	x             float32
	y             float32
	boxForeground *BoxObject
	boxBackground *BoxObject
}

func (l *LaufschriftObject) Render(r *Renderer, time float64) {
	//l.boxBackground.SetModelPositionAndSize(l.x, l.y, l.width, l.height)
	l.boxBackground.Render(r, time)

	l.texture.Bind(0)
	l.shader.Use()

	helper.UniformMatrix4(l.shader.GetUniform("projectionMatrix"), r.GetProjection())
	helper.Uniform1i(l.shader.GetUniform("renderedTexture"), 0)
	helper.Uniform1f(l.shader.GetUniform("time"), float32(time/2.5))

	r.RenderEntity(l.entity, l.shader)

	l.shader.UnUse()
	l.texture.UnBind()

	//l.boxForeground.SetModelPositionAndSize(l.x, l.y, l.width, l.height)
	l.boxForeground.Render(r, time)
}

func (l *LaufschriftObject) SetText(text string) {

	l.text = text
	l.image = addLabel(text)
	l.texture.SetDefaultImage(l.image)

}

func (l *LaufschriftObject) SetTextSafe(text string) {

	l.text = text
	l.image = addLabel(text)
	helper.AddFunction(func() {
		l.texture.SetDefaultImage(l.image)
		l.texture.Bind(0)
		helper.TexParameteri(helper.GlTexture2D, helper.GlTextureWrapS, helper.GlRepeat)
		helper.TexParameteri(helper.GlTexture2D, helper.GlTextureWrapR, helper.GlClampToEdge)
		l.texture.UnBind()
		l.UpdateEntity()
	})

}

func (l *LaufschriftObject) SetPosition(x, y float32) {
	l.entity.IncreasePosition(x, y, 0)
	l.boxBackground.SetPosition(x, y)
	l.boxForeground.SetPosition(x, y)
}

func (l *LaufschriftObject) SetColor(r, g, b int) {
	l.entity.SetColourRGB(r, g, b, 1)
}

func (l *LaufschriftObject) UpdateEntity() {
	height := l.height - 0.04
	y := l.y + 0.02

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

	simpleQuad := []float32{
		l.x, y + height, 0, //V0
		l.x, y, 0, //V1
		l.x + l.width, y, 0, //V2
		l.x + l.width, y + height, 0, //V3
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
	if l.entity != nil {
		log.Println("Set new data")
		l.entity.Model.SetPositions(simpleQuad)
		l.entity.Model.SetTexCoords(simpleTexture)
	} else {
		model := CreateModelWithDataTexture(simpleIndicies, simpleQuad, simpleTexture)
		l.entity = MakeEntity(model, mgl32.Vec3{0, 0, 0}, 0, 0, 0, 1.0)
	}

}

func CreateLaufschrift(text string, x, y, width, height float32) *LaufschriftObject {

	lShader, err := createLaufschriftShader()
	if err != nil {
		log.Println("create Laufschrift shader error:", err)
	}

	l := LaufschriftObject{
		texture:       NewTextureDefault(),
		shader:        lShader,
		boxForeground: CreateBox(x, y, width, height, "assets/gui/box/RectangleForeground.png"),
		boxBackground: CreateBox(x, y, width, height, "assets/gui/box/RectangleBackground.png"),
		x:             x,
		y:             y,
		width:         width,
		height:        height,
	}
	l.SetText(text)

	l.texture.Bind(0)
	helper.TexParameteri(helper.GlTexture2D, helper.GlTextureWrapS, helper.GlRepeat)
	helper.TexParameteri(helper.GlTexture2D, helper.GlTextureWrapR, helper.GlClampToEdge)
	l.texture.UnBind()

	l.UpdateEntity()

	return &l
}

func createLaufschriftShader() (*Program, error) {
	vert, err := NewShaderFromFile("assets/shaders/laufschrift.vert", VertexShaderType)
	if err != nil {
		return nil, err
	}

	frag, err := NewShaderFromFile("assets/shaders/laufschrift.frag", FragmentShaderType)
	if err != nil {
		return nil, err
	}

	shader, err := NewProgram(vert, frag)

	if err != nil {
		return nil, err
	}

	shader.Use()
	shader.AddUniform("renderedTexture")
	shader.AddUniform("projectionMatrix")
	shader.AddUniform("transformationMatrix")
	shader.AddUniform("time")
	shader.AddAndBindAttribute(0, "position")
	shader.AddAndBindAttribute(1, "texCoord")
	shader.UnUse()

	return shader, nil
}

func addLabel(label string) image.Image {

	// Read the font data.
	fontBytes, err := ioutil.ReadFile("assets/fonts/juice.ttf")
	if err != nil {
		log.Println("Laufschrift ReadFile Error:", err)
		return nil
	}
	f, err := freetype.ParseFont(fontBytes)

	if err != nil {
		log.Println("Laufschrift ParseFont Error:", err)
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
