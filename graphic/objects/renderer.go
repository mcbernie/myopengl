package objects

import (
	"github.com/mcbernie/myopengl/glHelper"
	//"math"
	//"math/rand"
	"log"

	//"github.com/go-gl/gl/v4.1-core/gl" // OR:
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/mcbernie/myopengl/gfx"
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
	glHelper.BindVertexArray(model.GetVao())

	gl.EnableVertexAttribArray(0)
	//gl.DrawArrays(gl.TRIANGLES, 0, model.GetVertexCount())
	gl.DrawElements(gl.TRIANGLES, model.GetVertexCount(), gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.DisableVertexAttribArray(0)

	glHelper.BindVertexArray(0)
	gl.UseProgram(0)
}
