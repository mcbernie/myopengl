package objects

import (
	"log"

	"github.com/go-gl/mathgl/mgl32"
)

//Entity struct saves raw model and add position, rotation and scale
type Entity struct {
	Model    *Model
	Position mgl32.Vec3
	Rx       float32
	Ry       float32
	Rz       float32
	Scale    float32
	color    [4]float32
}

//MakeEntity creates a new Entity Structure Object
func MakeEntity(model *Model, position mgl32.Vec3, rX, rY, rZ, scale float32) *Entity {
	return &Entity{
		Model:    model,
		Position: position,
		Rx:       rX,
		Ry:       rY,
		Rz:       rZ,
		Scale:    scale,
		color:    [4]float32{1, 1, 1, 1},
	}
}

//IncreasePosition increase Position otation by x,y,z float32
func (e *Entity) IncreasePosition(dx, dy, dz float32) {
	e.Position = e.Position.Add(mgl32.Vec3{dx, dy, dz})
}

//IncreaseRotation increase Rotation by x,y,z float32
func (e *Entity) IncreaseRotation(dx, dy, dz float32) {
	e.Rx += dx
	e.Ry += dy
	e.Rz += dz
}

//SetColour set the color of entity by red, green, blue, alpha float32
func (e *Entity) SetColour(rC, gC, bC, a float32) {
	e.color[0] = rC
	e.color[1] = gC
	e.color[2] = bC
	e.color[3] = a
}

//SetColourRGB set the color of entity by red, green, blue, alpha integer
func (e *Entity) SetColourRGB(rC, gC, bC, a int) {
	e.color[0] = (1.0 / 255.0) * float32(rC)
	e.color[1] = (1.0 / 255.0) * float32(gC)
	e.color[2] = (1.0 / 255.0) * float32(bC)
	e.color[3] = (1.0 / 255.0) * float32(a)
}

//Delete cleanup entity memory
func (e *Entity) Delete() {
	log.Println("delete entity")
	e.Model.Delete()
}
