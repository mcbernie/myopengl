package slideshow

import (
	"errors"
	"fmt"

	"github.com/mcbernie/myopengl/graphic/helper"
	"github.com/mcbernie/myopengl/graphic/objects"
)

const fragementShaderTemplate = `
#version 120
varying vec2 _uv;
uniform sampler2D from;
uniform sampler2D to;
uniform float progress;
uniform float ratio;
uniform float _fromR;
uniform float _toR;

vec4 getFromColor(vec2 uv) {
  return texture2D(from, %s);
}

vec4 getToColor(vec2 uv) {
  return texture2D(to, %s);
}

%s

void main() {
  gl_FragColor = transition(_uv);
}
`

//Transition simple Type Structure for a Transition
type Transition struct {
	glsl          string
	name          string
	defaultParams map[string]interface{}
	paramsTypes   map[string]string
	//Shader Holds the program for this Transition
	Shader *objects.Program
	//Name save the transition name
	Name string
}

//ResizeMode simple Enum which ResizeMode should this transition is using
type ResizeMode int

const (
	//Cover Fill Whole Image
	Cover ResizeMode = iota
	//Contain Keep ratio
	Contain
	//Stretch Stretch image
	Stretch
)

const (
	vert = `
	#version 120
	attribute vec2 position;
	varying vec2 _uv;

	uniform mat4 transformationMatrix;
	uniform mat4 projectionMatrix;

	void main() {
	gl_Position = projectionMatrix * transformationMatrix * vec4(position,0.0,1.0);
	vec2 uv = position * 0.5 + 0.5;
	_uv = vec2(uv.x, 1.0 - uv.y);
	}`
)

func resize(resizeMode ResizeMode) (func(r ...string) string, error) {
	switch resizeMode {
	case Cover:
		return func(r ...string) string {
			return ".5+(uv-.5)*vec2(min(ratio/" + r[0] + ",1.),min(" + r[0] + "/ratio,1.))"
		}, nil
	case Contain:
		return func(r ...string) string {
			return ".5+(uv-.5)*vec2(max(ratio/" + r[0] + ",1.),max(" + r[0] + "/ratio,1.))"
		}, nil
	case Stretch:
		return func(r ...string) string { return "uv" }, nil
	default:
		return nil, errors.New("unknown ResizeMode")
	}
}

func makeFrag(transitionGlsl string, resizeMode ResizeMode) string {
	r, err := resize(resizeMode)
	if err != nil {
		panic("invalid resizeMod=" + err.Error())
	}

	return fmt.Sprintf(fragementShaderTemplate, r("_fromR"), r("_toR"), transitionGlsl)
}

//MakeTransition generate a Transition with glsl
func MakeTransition(resizeMode ResizeMode, glsl string, name string, projection [16]float32) *Transition {

	// create a shader and put it in the thing here
	vertShader, err := objects.NewShader(vert, helper.GlVertexShader)
	if err != nil {
		panic("VertexShader error:" + err.Error())
	}
	fragShader, err := objects.NewShader(makeFrag(glsl, resizeMode), helper.GlFragmentShader)
	if err != nil {
		panic("FragmentShader error:" + err.Error())
	}

	program, err := objects.NewProgram(vertShader, fragShader)
	if err != nil {
		panic("Program Error:" + err.Error())
	}
	program.Use()
	program.AddAttribute("position")
	program.AddUniform("transformationMatrix")
	program.AddUniform("projectionMatrix")
	program.AddUniform("ratio")
	program.AddUniform("progress")
	program.AddUniform("from")
	program.AddUniform("to")
	program.AddUniform("_fromR")
	program.AddUniform("_toR")

	program.UnUse()
	return &Transition{
		Shader: program,
		glsl:   glsl,
		Name:   name,
	}
}

//Draw draws a transition
func (transition *Transition) Draw(progress float32, from *objects.Texture, to *objects.Texture, projection [16]float32 /*, width float32, height float32*/ /*, params map[string]interface{}*/) {
	shader := transition.Shader
	shader.Use()
	//glHelper.Uniform1f(shader.GetUniform("ratio"), width/height)
	helper.Uniform1f(shader.GetUniform("progress"), progress)

	from.Bind(0)
	to.Bind(1)

	helper.UniformMatrix4(shader.GetUniform("projectionMatrix"), projection)
	helper.Uniform1i(shader.GetUniform("from"), 0)
	helper.Uniform1i(shader.GetUniform("to"), 1)
	helper.Uniform1i(shader.GetUniform("_fromR"), from.Width/from.Height)
	helper.Uniform1i(shader.GetUniform("_toR"), to.Width/to.Height)
	// other...
	//shader.Delete()
}

//Delete remove shader program from memory
func (transition *Transition) CleanUP() {
	transition.Shader.Delete()
}
