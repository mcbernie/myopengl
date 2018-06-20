package glHelper

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/mathgl/mgl32"
)

//Init Initialize OpenGL
func Init() error {
	return gl.Init()
}

func ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

func Clear(mask uint32) {
	gl.Clear(mask)
}

//DeleteTextures removes texture from opengl device memory
func DeleteTextures(n int32, handle *uint32) {
	gl.DeleteTextures(n, handle)
}

//DeleteShader removes shader from opengl device memory
func DeleteShader(shader uint32) {
	gl.DeleteShader(shader)
}

//DeleteProgram removes shader program from opengl device memory
func DeleteProgram(program uint32) {
	gl.DeleteProgram(program)
}

//AttachShader Add Shader to Program
func AttachShader(program uint32, shader uint32) {
	gl.AttachShader(program, shader)
}

//UseProgram Set Program to current context
func UseProgram(program uint32) {
	gl.UseProgram(program)
}

//LinkProgram Link a Program
func LinkProgram(program uint32) {
	gl.LinkProgram(program)
}

type getObjIv func(uint32, uint32, *int32)
type getObjInfoLog func(uint32, int32, *int32, *uint8)

//GetGlError returns current gl error
func GetGlError(glHandle uint32, checkTrueParam uint32, getObjIvFn getObjIv,
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

//CreateShader Creates a new empty shader for sType and returns an shader handle
func CreateShader(sType uint32) uint32 {
	return gl.CreateShader(sType)
}

//ShaderSource Set the uncompiles sourcecode to an shader
func ShaderSource(shader uint32, count int32, glSrc **uint8, length *int32) {
	gl.ShaderSource(shader, count, glSrc, length)
	//gl.ShaderSource
}

//CompileShader Compiles an Shader
func CompileShader(shader uint32) {
	gl.CompileShader(shader)
}

//Strs Converts golang string to opengl string
func Strs(src string) (**uint8, func()) {
	return gl.Strs(string(src) + "\x00")
}

//GetShaderiv returns an iv?
func GetShaderiv(shader uint32, pname uint32, params *int32) {
	gl.GetShaderiv(shader, pname, params)
}

//GetShaderInfoLog returns Log Data for Shader
func GetShaderInfoLog(shader uint32, bufSize int32, length *int32, infoLog *uint8) {
	gl.GetShaderInfoLog(shader, bufSize, length, infoLog)
}

//GetProgramiv returns iv ?
func GetProgramiv(program uint32, pname uint32, params *int32) {
	gl.GetProgramiv(program, pname, params)
}

//GetProgramInfoLog Returns al Log Data from Opengl for an Program
func GetProgramInfoLog(program uint32, bufSize int32, length *int32, infoLog *uint8) {
	gl.GetProgramInfoLog(program, bufSize, length, infoLog)
}

//GetAttribLocation Returns the Location of an Shader Attribute
func GetAttribLocation(program uint32, src string) int32 {
	return gl.GetAttribLocation(program, gl.Str(src+"\x00"))
}

//GetUniformLocation add Uniform Location to an map for easier and faster access
func GetUniformLocation(program uint32, src string) int32 {
	return gl.GetUniformLocation(program, gl.Str(src+"\x00"))
}

//BindAttribLocation bind vertex variable to vertex shader
func BindAttribLocation(program uint32, index uint32, src string) {
	gl.BindAttribLocation(program, index, gl.Str(src+"\x00"))
}

//CreateProgram Creates a new Shader Program and returns handle
func CreateProgram() uint32 {
	return gl.CreateProgram()
}

//Uniform1f Set 1 float32 value in location
func Uniform1f(location int32, v0 float32) {
	gl.Uniform1f(location, v0)
}

//Uniform1i set 1 int32 value in location
func Uniform1i(location int32, v0 int32) {
	gl.Uniform1i(location, v0)
}

//GenTextures generate x textures and returns an handle or array of handles
func GenTextures(n int32, handle *uint32) {
	gl.GenTextures(n, handle)
}

//TexParameteri set parameter pname with value param in texture handle
func TexParameteri(handle uint32, pname uint32, param int32) {
	gl.TexParameteri(handle, pname, param)
}

//TexImage2D setup the texture2d
func TexImage2D(handle uint32, level int32, internalFmt int32, w int32, h int32,
	border int32, format uint32, xtype uint32, pixels unsafe.Pointer) {
	gl.TexImage2D(handle, level, internalFmt, w, h, border, format, xtype, pixels)
}

//GenerateMipmap generates a MipMap
func GenerateMipmap(handle uint32) {
	gl.GenerateMipmap(handle)
}

func BindTexture(target uint32, texture uint32) {
	gl.BindTexture(target, texture)
}

func ActiveTexture(texture uint32) {
	gl.ActiveTexture(texture)
}

//Ptr get an unsafe Pointer for gl objects
func Ptr(data interface{}) unsafe.Pointer {
	return gl.Ptr(data)
}

//gl.UniformMatrix4fv(shader.GetUniform("transformationMatrix"), 1, false, &translationMatrix[0])

func UniformMatrix4(location int32, matrix mgl32.Mat4) {
	gl.UniformMatrix4fv(location, 1, false, &matrix[0])
}
