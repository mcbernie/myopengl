package objects

import (
	//"math"
	//"math/rand"
	//"log"

	"log"

	//"github.com/go-gl/gl/v4.1-core/gl" // OR:
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/mcbernie/myopengl/gfx"
	//"github.com/mcbernie/myopengl/gfx"
)

type Renderer struct {
	shader *gfx.Program
}

func MakeRenderer() *Renderer {
	shader, err := createDefaultShader()

	if err != nil {
		log.Println("Error on create default shader:", err)
	}

	return &Renderer{
		shader: shader,
	}
}

func (r *Renderer) prepare() {
}

func (r *Renderer) UseDefaultShader() {
	r.shader.Use()
}

func (r *Renderer) Render(model *RawModel) {

	//gl.BindVertexArray(model.GetVao())
	gl.BindVertexArrayAPPLE(model.GetVao())

	gl.EnableVertexAttribArray(0)
	gl.DrawArrays(gl.TRIANGLES, 0, model.GetVertexCount())
	gl.DisableVertexAttribArray(0)

	//gl.BindVertexArray(0)
	gl.BindVertexArrayAPPLE(model.GetVao())

	gl.UseProgram(0)
}
