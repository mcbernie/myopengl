package graphic

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/mcbernie/myopengl/gfx"
)

const defaultFrag = `
#version 120
varying vec4 vColor;

void main()
{
  gl_FragColor = vColor;
}
`

const defaultVert = `
#version 120

uniform mat4 u_projView;

attribute vec2 Position;
attribute vec4 Color;

varying vec4 vColor;
//varying vec2 vTexCoord;

void main()
{
	vColor = Color;
	//vTexCoord = TexCoord;
	//gl_Position = u_projView * vec4(Position, 0.0, 1.0);
}
`

func createDefaultShader() (*gfx.Program, error) {

	// create a shader and put it in the thing here
	vertShader, err := gfx.NewShader(defaultVert, gl.VERTEX_SHADER)
	if err != nil {
		panic("VertexShader error:" + err.Error())
	}
	fragShader, err := gfx.NewShader(defaultFrag, gl.FRAGMENT_SHADER)
	if err != nil {
		panic("FragmentShader error:" + err.Error())
	}

	program, err := gfx.NewProgram(vertShader, fragShader)
	if err != nil {
		panic("Program Error:" + err.Error())
	}

	program.AddUniform("u_projView")
	program.AddAttribute("Position")
	program.AddAttribute("Color")

	return program, nil
}
