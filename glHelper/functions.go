// +build !darwin

package glHelper

import (
	//"github.com/go-gl/gl/v4.1-core/gl" // OR:
	"log"

	"github.com/go-gl/gl/v2.1/gl"
)

//GenerateVertexArray Generate an VAO
func GenerateVertexArray(n int32) uint32 {
	var vao uint32
	gl.GenVertexArrays(n, &vao)
	if err := gl.GetError(); err != 0 {
		log.Println("Error in GenVertexArrays!", err)
	}

	return vao
}

//BindVertexArray binds an VAO
func BindVertexArray(vao uint32) {
	gl.BindVertexArray(vao)
}

//DeleteVertexArrary Removes an Vertex Array Object from Memory
func DeleteVertexArrary(n int32, arrays *uint32) {
	gl.DeleteVertexArrays(n, arrays)
}
