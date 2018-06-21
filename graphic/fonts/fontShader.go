package fonts

import (
	"log"

	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/glHelper"
)

type FontShader struct {
	*gfx.Program
}

func CreateFontShader() *FontShader {
	// create a shader and put it in the thing here
	vertShader, err := gfx.NewShaderFromFile("assets/fonts/font.vert", gfx.VertexShaderType)
	if err != nil {
		panic("VertexShader error:" + err.Error())
	}

	fragShader, err := gfx.NewShaderFromFile("assets/fonts/font.frag", gfx.FragmentShaderType)
	if err != nil {
		panic("FragmentShader error:" + err.Error())
	}

	//create a Shader Program and...
	program, err := gfx.NewProgram(vertShader, fragShader)
	if err != nil {
		panic("Program Error:" + err.Error())
	}

	//Put the ShaderProgram to Abstraction of Program -> FontShader
	f := &FontShader{program}

	f.Use()
	f.AddUniform("colour")
	f.AddUniform("translation")

	f.BindAttribute(0, "position")
	f.BindAttribute(1, "textureCoords")

	f.UnUse()

	log.Println("called.. createfontShader...")

	return f
}

func (f *FontShader) SetColur(colour [3]float32) {
	glHelper.Uniform3f(f.GetUniform("colour"), colour)
}

func (f *FontShader) SetTranslation(translation [2]float32) {
	glHelper.Uniform2f(f.GetUniform("translation"), translation[0], translation[1])
}
