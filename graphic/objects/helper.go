package objects

import (
	"github.com/mcbernie/myopengl/graphic/helper"
)

const defaultFrag = `
#version 150

out vec4 fragColor;
uniform vec4 color;

void main()
{
  fragColor = color;
}
`

const defaultVert = `
#version 150

in vec4 position;

uniform mat4 transformationMatrix;
uniform mat4 projectionMatrix;


void main()
{
	gl_Position = projectionMatrix * transformationMatrix * position;//vec4(position.zyx, 1.0);
}
`

func createDefaultShader() (*Program, error) {

	// create a shader and put it in the thing here
	vertShader, err := NewShader(defaultVert, helper.GlVertexShader)
	if err != nil {
		panic("VertexShader error:" + err.Error())
	}

	fragShader, err := NewShader(defaultFrag, helper.GlFragmentShader)
	if err != nil {
		panic("FragmentShader error:" + err.Error())
	}

	program, err := NewProgram(vertShader, fragShader)
	if err != nil {
		panic("Program Error:" + err.Error())
	}

	program.AddUniform("color")
	return program, nil
}
