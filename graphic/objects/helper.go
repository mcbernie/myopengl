package objects

import (
	"github.com/mcbernie/myopengl/gfx"
	"github.com/mcbernie/myopengl/glHelper"
)

const defaultFrag = `
#version 120
//varying vec4 vColor;

void main()
{
  gl_FragColor = vec4(0.5, 0.8, 1.0, 1.0);
}
`

const defaultVert = `
#version 120

//uniform mat4 u_projView;

attribute vec4 position;
//attribute vec4 Color;

//varying vec4 vColor;
//varying vec2 vTexCoord;

void main()
{
	//vColor = Color;
	//vTexCoord = TexCoord;
	//gl_Position = u_projView * vec4(Position, 0.0, 1.0);
	gl_Position = position;//vec4(position.zyx, 1.0);
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

	/*program.AddUniform("u_projView")
	program.AddAttribute("position")
	program.AddAttribute("Color")*/

	return program, nil
}