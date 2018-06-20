package objects

import (
	"github.com/go-gl/mathgl/mgl32"
	//"github.com/go-gl/gl/v4.1-core/gl" // OR:
)

type Entity struct {
	Model    *RawModel
	Position mgl32.Vec3
	Rx       float32
	Ry       float32
	Rz       float32
	Scale    float32
}

func MakeEntity(model *RawModel, position mgl32.Vec3, rX, rY, rZ, scale float32) *Entity {
	return &Entity{
		Model:    model,
		Position: position,
		Rx:       rX,
		Ry:       rY,
		Rz:       rZ,
		Scale:    scale,
	}
}

func (e *Entity) IncreasePosition(dx, dy, dz float32) {
	e.Position = e.Position.Add(mgl32.Vec3{dx, dy, dz})
}

func (e *Entity) IncreaseRotation(dx, dy, dz float32) {
	e.Rx += dx
	e.Ry += dy
	e.Rz += dz
}
