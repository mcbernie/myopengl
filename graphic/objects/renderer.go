package objects

import (
	"github.com/mcbernie/myopengl/glHelper"
	//"math"
	//"math/rand"
	"log"

	//"github.com/go-gl/gl/v4.1-core/gl" // OR:
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mcbernie/myopengl/gfx"
)

type Renderer struct {
	Shader           *gfx.Program
	projectionMatrix mgl32.Mat4
}

func MakeRenderer() *Renderer {
	projectionMatrix := mgl32.Perspective(70, 1, .01, 1000)

	shader, err := createDefaultShader()

	shader.Use()
	shader.AddUniform("projectionMatrix")
	glHelper.UniformMatrix4(shader.GetUniform("projectionMatrix"), projectionMatrix)
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
	glHelper.BindVertexArray(model.GetVao())

	gl.EnableVertexAttribArray(0)
	gl.DrawElements(gl.TRIANGLES, model.GetVertexCount(), gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.DisableVertexAttribArray(0)

	glHelper.BindVertexArray(0)
	gl.UseProgram(0)
}

func (r *Renderer) RenderEntity(e *Entity, shader *gfx.Program) {
	glHelper.BindVertexArray(e.Model.GetVao())
	gl.EnableVertexAttribArray(0)

	rotMatrix := mgl32.HomogRotate3DX(e.Rx)
	rotMatrix = rotMatrix.Mul4(mgl32.HomogRotate3DY(e.Ry))
	rotMatrix = rotMatrix.Mul4(mgl32.HomogRotate3DZ(e.Rz))

	scaleMatrix := mgl32.Scale3D(e.Scale, e.Scale, e.Scale)
	translationMatrix := rotMatrix.Mul4(mgl32.Translate3D(e.Position.X(), e.Position.Y(), e.Position.Z()))
	translationMatrix = translationMatrix.Mul4(scaleMatrix)

	glHelper.UniformMatrix4(shader.GetUniform("transformationMatrix"), translationMatrix)

	gl.DrawElements(gl.TRIANGLES, e.Model.GetVertexCount(), gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.DisableVertexAttribArray(0)

	glHelper.BindVertexArray(0)
	gl.UseProgram(0)
}
