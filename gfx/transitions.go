package gfx

import (
	"errors"
	"fmt"

	"github.com/mcbernie/myopengl/glHelper"
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
	Shader *Program
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
	void main() {
	gl_Position = vec4(position,0.0,1.0);
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
func MakeTransition(resizeMode ResizeMode, glsl string, name string) *Transition {

	// create a shader and put it in the thing here
	vertShader, err := NewShader(vert, glHelper.GlVertexShader)
	if err != nil {
		panic("VertexShader error:" + err.Error())
	}
	fragShader, err := NewShader(makeFrag(glsl, resizeMode), glHelper.GlFragmentShader)
	if err != nil {
		panic("FragmentShader error:" + err.Error())
	}

	program, err := NewProgram(vertShader, fragShader)
	if err != nil {
		panic("Program Error:" + err.Error())
	}
	//program.Use()
	program.AddAttribute("position")
	program.AddUniform("ratio")
	program.AddUniform("progress")
	program.AddUniform("from")
	program.AddUniform("to")
	program.AddUniform("_fromR")
	program.AddUniform("_toR")

	return &Transition{
		Shader: program,
		glsl:   glsl,
		Name:   name,
	}
}

//Draw draws a transition
func (transition *Transition) Draw(progress float32, from *Texture, to *Texture, width float32, height float32 /*, params map[string]interface{}*/) {
	shader := transition.Shader
	shader.Use()
	glHelper.Uniform1f(shader.GetUniform("ratio"), width/height)
	glHelper.Uniform1f(shader.GetUniform("progress"), progress)

	from.Bind(0)
	to.Bind(1)
	glHelper.Uniform1i(shader.GetUniform("from"), 0)
	glHelper.Uniform1i(shader.GetUniform("to"), 1)
	glHelper.Uniform1i(shader.GetUniform("_fromR"), from.width/from.height)
	glHelper.Uniform1i(shader.GetUniform("_toR"), to.width/to.height)
	// other...
	//shader.Delete()
}

//Delete remove shader program from memory
func (transition *Transition) Delete() {
	transition.Shader.Delete()
}
