package objects

import (
	"log"

	//"github.com/go-gl/gl/v4.1-core/gl" // OR:

	"github.com/go-gl/mathgl/mgl32"
	"github.com/mcbernie/myopengl/graphic/helper"
)

//type RenderFunction func(interface{}, float64)

type ObjectInterface interface {
	Render(r *Renderer, time float64)
}

type Object struct {
	Name     string
	Textures []*Texture
	Model    *Model
	Entity   *Entity
	Shader   *Program
}

func (obj *Object) Render(r *Renderer, time float64) {

	for t, tex := range obj.Textures {
		tex.Bind(uint32(t))
		defer tex.UnBind()
	}
	obj.Shader.Use()
	defer obj.Shader.UnUse()
	helper.Uniform1i(obj.Shader.GetUniform("renderedTexture"), 0)
	helper.Uniform1f(obj.Shader.GetUniform("time"), float32(time/2.5))
	r.RenderEntity(obj.Entity, obj.Shader)
}

type ObjectsList struct {
	objects  []ObjectInterface
	renderer *Renderer
}

func CreateObjectList(r *Renderer) ObjectsList {
	o := ObjectsList{
		renderer: r,
	}

	return o
}

func (o *ObjectsList) AddRenderer(r ObjectInterface) {
	o.objects = append(o.objects, r)
}

func (o *ObjectsList) Render(time float64) {
	for _, r := range o.objects {
		r.Render(o.renderer, time)
	}
}

type Renderer struct {
	Shader           *Program
	projectionMatrix mgl32.Mat4
}

func MakeRenderer() *Renderer {
	//projectionMatrix := mgl32.Perspective(70, 1, .01, 1000)
	projectionMatrix := mgl32.Ortho2D(-1, 1, -1, 1)
	shader, err := createDefaultShader()

	shader.Use()
	shader.AddUniform("projectionMatrix")
	helper.UniformMatrix4(shader.GetUniform("projectionMatrix"), projectionMatrix)
	shader.UnUse()

	if err != nil {
		log.Println("Error on create default shader:", err)
	}

	return &Renderer{
		Shader:           shader,
		projectionMatrix: projectionMatrix,
	}
}

func (r *Renderer) prepare() {
}

func (r *Renderer) GetProjection() [16]float32 {
	return r.projectionMatrix
}

func (r *Renderer) UseDefaultShader() {
	r.Shader.Use()
}

func (r *Renderer) Render(model *RawModel) {
	helper.BindVertexArray(model.GetVao())

	helper.EnableVertexAttribArray(0)
	helper.DrawElements(helper.GlTriangles, model.GetVertexCount(), helper.GlUnsignedInt, helper.PtrOffset(0))
	helper.DisableVertexAttribArray(0)

	helper.BindVertexArray(0)
	helper.UseProgram(0)
}

func (r *Renderer) RenderEntity(e *Entity, shader *Program) {
	e.Model.Bind()

	helper.EnableVertexAttribArray(0)
	helper.EnableVertexAttribArray(1)
	//glHelper.EnableVertexAttribArray(2)

	rotMatrix := mgl32.HomogRotate3DX(e.Rx)
	rotMatrix = rotMatrix.Mul4(mgl32.HomogRotate3DY(e.Ry))
	rotMatrix = rotMatrix.Mul4(mgl32.HomogRotate3DZ(e.Rz))

	scaleMatrix := mgl32.Scale3D(e.Scale, e.Scale, e.Scale)
	translationMatrix := rotMatrix.Mul4(mgl32.Translate3D(e.Position.X(), e.Position.Y(), e.Position.Z()))
	translationMatrix = translationMatrix.Mul4(scaleMatrix)
	helper.Uniform4f(shader.GetUniform("color"), e.color)
	helper.UniformMatrix4(shader.GetUniform("transformationMatrix"), translationMatrix)

	//gl.DrawElements(gl.TRIANGLES, e.Model.vertexCount, gl.UNSIGNED_INT, gl.PtrOffset(0))
	e.Model.Draw()

	helper.DisableVertexAttribArray(0)
	helper.DisableVertexAttribArray(1)

	e.Model.UnBind()
}
