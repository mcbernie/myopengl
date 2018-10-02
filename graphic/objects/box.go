package objects

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mcbernie/myopengl/graphic/helper"
)

type BoxObject struct {
	x       float32
	y       float32
	width   float32
	height  float32
	shader  *Program
	texture *Texture
	entity  *Entity
}

func (b *BoxObject) Render(r *Renderer, time float64) {
	b.texture.Bind(0)
	b.shader.Use()

	helper.UniformMatrix4(b.shader.GetUniform("projectionMatrix"), r.GetProjection())
	helper.Uniform1i(b.shader.GetUniform("tex"), 0)

	// clip = 24 * 24 // 372
	// slice == 8 * 8 // 124

	// u_dimensions = vec2(slice/box.w, slice/box.h)
	// u_border = vec2(slice/clip.w, slice/clip.h)

	// u_dimensions = vec2(124/box.w, 124/box.h)
	// u_border = vec2(124/372, 124/372)

	helper.Uniform2f(b.shader.GetUniform("u_dimensions"), 0.009, 0.1)           // windowBorder
	helper.Uniform2f(b.shader.GetUniform("u_border"), 124.0/372.0, 124.0/372.0) // textureborder

	r.RenderEntity(b.entity, b.shader)

	b.shader.UnUse()
	b.texture.UnBind()
}

func (b *BoxObject) SetPosition(x, y float32) {
	b.entity.IncreasePosition(x, y, 0.0)
}

func CreateBox(x, y, width, height float32, texturePath string) *BoxObject {

	vert, err := NewShaderFromFile("assets/shaders/box.vert", VertexShaderType)
	if err != nil {
		panic(err)
	}
	frag, err := NewShaderFromFile("assets/shaders/box.frag", FragmentShaderType)
	if err != nil {
		panic(err)
	}
	shader, _ := NewProgram(vert, frag)

	shader.Use()
	shader.AddUniform("projectionMatrix")
	shader.AddUniform("transformationMatrix")
	shader.AddUniform("tex")
	shader.AddUniform("u_dimensions")
	shader.AddUniform("u_border")
	shader.AddAndBindAttribute(0, "position")
	shader.AddAndBindAttribute(1, "texCoord")
	shader.UnUse()

	texture := NewTextureFromFile(texturePath)

	b := BoxObject{
		x:       x,
		y:       y,
		width:   width,
		height:  height,
		texture: texture,
		shader:  shader,
	}

	simpleQuad := b._setModelPositionAndSize(x, y, width, height)

	simpleTexture := []float32{
		0, 1, //V0 (x,y)
		0, 0, //V1 (x,y)
		1, 0, //V2 (x,y)
		1, 1, //V3 (x,y)
	}

	simpleIndicies := []int32{
		0, 1, 3, //Top Left triangle (V0, V1, V3)
		3, 1, 2, //Bottom Right triangle (V3, V1, V2)
	}

	model := CreateModelWithDataTexture(simpleIndicies, simpleQuad, simpleTexture)
	entity := MakeEntity(model, mgl32.Vec3{0, 0, 0}, 0, 0, 0, 1.0)
	b.entity = entity

	return &b
}

func (b *BoxObject) _setModelPositionAndSize(x, y, width, height float32) []float32 {
	b.x = x
	b.y = y
	b.width = width
	b.height = height

	return []float32{
		b.x, b.y + b.height, 0, //V0
		b.x, b.y, 0, //V1
		b.x + b.width, b.y, 0, //V2
		b.x + b.width, b.y + b.height, 0, //V3
	}

}

func (b *BoxObject) SetModelPositionAndSize(x, y, width, height float32) {
	b.entity.Model.SetPositions(b._setModelPositionAndSize(x, y, width, height))
}
