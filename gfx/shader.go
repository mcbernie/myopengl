package gfx

import (
	"fmt"
	"io/ioutil"
	"strings"

	//"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/gl/v2.1/gl"
)

//Shader is the structure for Shader
type Shader struct {
	handle uint32
}

//Program is the structure to hold the gl Program with all Shaders
type Program struct {
	handle             uint32
	shaders            []*Shader
	uniformLocations   map[string]int32
	attributeLocations map[string]uint32
}

//Delete remove a shader from gl.Program, shader
func (shader *Shader) Delete() {
	gl.DeleteShader(shader.handle)
}

//Delete remove a Program from OpenGL
func (prog *Program) Delete() {
	for _, shader := range prog.shaders {
		shader.Delete()
	}
	gl.DeleteProgram(prog.handle)
}

//Attach add multiple Shaders to a Program
func (prog *Program) Attach(shaders ...*Shader) {
	for _, shader := range shaders {
		gl.AttachShader(prog.handle, shader.handle)
		prog.shaders = append(prog.shaders, shader)
	}
}

//Use Enable the Program in OpenGL
func (prog *Program) Use() {
	gl.UseProgram(prog.handle)
}

//Link Linking the Program with OpenGL
func (prog *Program) Link() error {
	gl.LinkProgram(prog.handle)
	err := getGlError(prog.handle, gl.LINK_STATUS, gl.GetProgramiv, gl.GetProgramInfoLog,
		"PROGRAM::LINKING_FAILURE")
	return err
}

//AddAttribute add AttributeLocation to an map for easier and faster access
func (prog *Program) AddAttribute(name string) {
	prog.attributeLocations[name] = uint32(gl.GetAttribLocation(prog.handle, gl.Str(name+"\x00")))
}

//AddUniform add Uniform Location to an map for easier and faster access
func (prog *Program) AddUniform(name string) {
	prog.uniformLocations[name] = gl.GetUniformLocation(prog.handle, gl.Str(name+"\x00"))
}

//GetUniform get the Uniform Location from Map object
func (prog *Program) GetUniform(name string) int32 {
	//log.Println("getUniformLocations:", prog.uniformLocations[name], name, gl.GetUniformLocation(prog.handle, gl.Str(name+"\x00")))
	//return prog.uniformLocations[name] //
	return gl.GetUniformLocation(prog.handle, gl.Str(name+"\x00"))
}

//GetAttribute get the Attribute Location from Map object
func (prog *Program) GetAttribute(name string) int32 {
	//return prog.attributeLocations[name]
	return gl.GetAttribLocation(prog.handle, gl.Str(name+"\x00"))
}

//NewProgram initialize a Program with shaders
func NewProgram(shaders ...*Shader) (*Program, error) {
	prog := &Program{
		handle:             gl.CreateProgram(),
		uniformLocations:   make(map[string]int32),
		attributeLocations: make(map[string]uint32),
	}

	prog.Attach(shaders...)

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

	handle := gl.CreateShader(sType)
	glSrc, freeFn := gl.Strs(src + "\x00")
	defer freeFn()
	gl.ShaderSource(handle, 1, glSrc, nil)
	gl.CompileShader(handle)
	err := getGlError(handle, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog,
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
	handle := gl.CreateShader(sType)
	glSrc, freeFn := gl.Strs(string(src) + "\x00")
	defer freeFn()
	gl.ShaderSource(handle, 1, glSrc, nil)
	gl.CompileShader(handle)
	err = getGlError(handle, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog,
		"SHADER::COMPILE_FAILURE::"+file)
	if err != nil {
		return nil, err
	}
	return &Shader{handle: handle}, nil
}

type getObjIv func(uint32, uint32, *int32)
type getObjInfoLog func(uint32, int32, *int32, *uint8)

func getGlError(glHandle uint32, checkTrueParam uint32, getObjIvFn getObjIv,
	getObjInfoLogFn getObjInfoLog, failMsg string) error {

	var success int32
	getObjIvFn(glHandle, checkTrueParam, &success)

	if success == gl.FALSE {
		var logLength int32
		getObjIvFn(glHandle, gl.INFO_LOG_LENGTH, &logLength)

		log := gl.Str(strings.Repeat("\x00", int(logLength)))
		getObjInfoLogFn(glHandle, logLength, nil, log)

		return fmt.Errorf("%s: %s", failMsg, gl.GoStr(log))
	}

	return nil
}
