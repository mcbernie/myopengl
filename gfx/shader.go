package gfx

import (
	"io/ioutil"

	"github.com/mcbernie/myopengl/glHelper"
)

const (
	//VertexShaderType for openGL Program / Shader
	VertexShaderType = glHelper.GlVertexShader
	//FragmentShaderType for openGL Program / Shader
	FragmentShaderType = glHelper.GlFragmentShader
)

type ShaderProgram interface {
	Attach(shaders ...*Shader)
	Use()
	UnUse()
	Link() error
	AddAttribute(name string)
	AddUniform(name string)
	GetUniform(name string) int32
	GetAttribute(name string) int32
	BindAttribute(index uint32, name string)
}

//Shader is the structure for Shader
type Shader struct {
	handle uint32
}

//Program is the structure to hold the gl Program with all Shaders
type Program struct {
	ShaderProgram
	handle             uint32
	shaders            []*Shader
	uniformLocations   map[string]int32
	attributeLocations map[string]uint32
}

//Delete remove a shader from gl.Program, shader
func (shader *Shader) Delete() {
	glHelper.DeleteShader(shader.handle)
}

//Delete remove a Program from OpenGL
func (prog *Program) Delete() {
	for _, shader := range prog.shaders {
		shader.Delete()
	}
	glHelper.DeleteProgram(prog.handle)
}

//Attach add multiple Shaders to a Program
func (prog *Program) Attach(shaders ...*Shader) {
	for _, shader := range shaders {
		glHelper.AttachShader(prog.handle, shader.handle)
		prog.shaders = append(prog.shaders, shader)
	}
}

//Use Enable the Program in OpenGL
func (prog *Program) Use() {
	glHelper.UseProgram(prog.handle)
}

//UnUse Disable the Program in OpenGL
func (prog *Program) UnUse() {
	glHelper.UseProgram(0)
}

//Link Linking the Program with OpenGL
func (prog *Program) Link() error {
	glHelper.LinkProgram(prog.handle)
	err := glHelper.GetGlError(prog.handle, glHelper.GlLinkStatus, glHelper.GetProgramiv, glHelper.GetProgramInfoLog,
		"PROGRAM::LINKING_FAILURE")
	return err
}

//AddAttribute add AttributeLocation to an map for easier and faster access
func (prog *Program) AddAttribute(name string) {
	prog.attributeLocations[name] = uint32(glHelper.GetAttribLocation(prog.handle, name))
}

//AddUniform add Uniform Location to an map for easier and faster access
func (prog *Program) AddUniform(name string) {
	prog.uniformLocations[name] = glHelper.GetUniformLocation(prog.handle, name)
}

//GetUniform get the Uniform Location from Map object
func (prog *Program) GetUniform(name string) int32 {
	//log.Println("getUniformLocations:", prog.uniformLocations[name], name, gl.GetUniformLocation(prog.handle, gl.Str(name+"\x00")))
	//return prog.uniformLocations[name] //
	return glHelper.GetUniformLocation(prog.handle, name)
}

//GetAttribute get the Attribute Location from Map object
func (prog *Program) GetAttribute(name string) int32 {
	//return prog.attributeLocations[name]
	return glHelper.GetAttribLocation(prog.handle, name)
}

//BindAttribute bind vertex variable to vertex shader
func (prog *Program) BindAttribute(index uint32, name string) {
	glHelper.BindAttribLocation(prog.handle, index, name)

}

//NewProgram initialize a Program with shaders
func NewProgram(shaders ...*Shader) (*Program, error) {
	prog := &Program{
		handle:             glHelper.CreateProgram(),
		uniformLocations:   make(map[string]int32),
		attributeLocations: make(map[string]uint32),
	}

	prog.Attach(shaders...)
	prog.BindAttribute(0, "position")
	if err := prog.Link(); err != nil {
		return nil, err
	}

	//After linking the Program. delete the shader...
	for _, s := range prog.shaders {
		s.Delete()
	}

	return prog, nil
}

//NewShader Creates a new Shader with in sType specified Shader Type from Source String
func NewShader(src string, sType uint32) (*Shader, error) {

	handle := glHelper.CreateShader(sType)
	glSrc, freeFn := glHelper.Strs(src + "\x00")
	defer freeFn()
	glHelper.ShaderSource(handle, 1, glSrc, nil)
	glHelper.CompileShader(handle)
	err := glHelper.GetGlError(handle, glHelper.GlCompileStatus, glHelper.GetShaderiv, glHelper.GetShaderInfoLog,
		"SHADER::COMPILE_FAILURE::")
	if err != nil {
		return nil, err
	}
	return &Shader{handle: handle}, nil
}

//NewShaderFromFile creates a new Shader from File
func NewShaderFromFile(file string, sType uint32) (*Shader, error) {
	src, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	handle := glHelper.CreateShader(sType)
	glSrc, freeFn := glHelper.Strs(string(src) + "\x00")
	defer freeFn()
	glHelper.ShaderSource(handle, 1, glSrc, nil)
	glHelper.CompileShader(handle)
	err = glHelper.GetGlError(handle, glHelper.GlCompileStatus, glHelper.GetShaderiv, glHelper.GetShaderInfoLog,
		"SHADER::COMPILE_FAILURE::"+file)
	if err != nil {
		return nil, err
	}
	return &Shader{handle: handle}, nil
}
