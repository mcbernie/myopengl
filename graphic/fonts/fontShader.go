package fonts

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mcbernie/myopengl/graphic/helper"
	"github.com/mcbernie/myopengl/graphic/objects"
)

type FontShader struct {
	*objects.Program
}

func CreateFontShader() *FontShader {
	// create a shader and put it in the thing here
	vertShader, err := objects.NewShaderFromFile("assets/shaders/font.vert", objects.VertexShaderType)
	if err != nil {
		panic("VertexShader error:" + err.Error())
	}

	fragShader, err := objects.NewShaderFromFile("assets/shaders/font.frag", objects.FragmentShaderType)
	if err != nil {
		panic("FragmentShader error:" + err.Error())
	}

	//create a Shader Program and...
	program, err := objects.NewProgram(vertShader, fragShader)
	if err != nil {
		panic("Program Error:" + err.Error())
	}

	projectionMatrix := mgl32.Ortho2D(-5, 5, -5, 5)

	//Put the ShaderProgram to Abstraction of Program -> FontShader
	f := &FontShader{program}

	f.Use()
	f.AddUniform("fontAtlas")
	f.AddUniform("colour")
	f.AddUniform("translation")

	f.AddUniform("projectionMatrix")
	helper.UniformMatrix4(f.GetUniform("projectionMatrix"), projectionMatrix)

	f.BindAttribute(0, "position")
	f.BindAttribute(1, "textureCoords")

	f.UnUse()

	return f
}

func (f *FontShader) SetColur(colour [3]float32) {
	helper.Uniform3f(f.GetUniform("colour"), colour)
}

func (f *FontShader) SetTranslation(translation [2]float32) {
	helper.Uniform2f(f.GetUniform("translation"), translation[0]+1.0, translation[1]-1.0)
}
