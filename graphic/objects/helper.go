package objects

import (
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/glHelper"
)

const defaultFrag = `
#version 120
//varying vec4 vColor;
uniform vec4 color;
void main()
{
  gl_FragColor = color;
}
`

const defaultVert = `
#version 120

attribute vec4 position;

uniform mat4 transformationMatrix;
uniform mat4 projectionMatrix;


void main()
{
	gl_Position = projectionMatrix * transformationMatrix * position;//vec4(position.zyx, 1.0);
}
`

func createDefaultShader() (*gfx.Program, error) {

	// create a shader and put it in the thing here
	vertShader, err := gfx.NewShader(defaultVert, glHelper.GlVertexShader)
	if err != nil {
		panic("VertexShader error:" + err.Error())
	}

	fragShader, err := gfx.NewShader(defaultFrag, glHelper.GlFragmentShader)
	if err != nil {
		panic("FragmentShader error:" + err.Error())
	}

	program, err := gfx.NewProgram(vertShader, fragShader)
	if err != nil {
		panic("Program Error:" + err.Error())
	}

	program.AddUniform("color")
	/*program.AddUniform("u_projView")
	program.AddAttribute("position")
	program.AddAttribute("Color")*/

	return program, nil
}
