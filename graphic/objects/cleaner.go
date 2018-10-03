package objects

import "github.com/mcbernie/myopengl/graphic/helper"

//TextureCleaner simple helper structure to keep infos abouot Textures, and remove after close
type TextureCleaner struct {
	handle  uint32
	counter int32
	name    string
}

//CreateTextureCleaner create a TextureCleaner, set a name and the handler from opengl
func CreateTextureCleaner(name string, handle uint32) *TextureCleaner {
	return &TextureCleaner{
		name:   name,
		handle: handle,
	}
}

//AddHandle Add a Handle
func (tc *TextureCleaner) AddHandle(handle uint32) {
	tc.handle = handle
}

//AddUser increase the "in use"-counter by 1
func (tc *TextureCleaner) AddUser() {
	tc.counter++
}

//RemoveUser decrease "in use"-counter by 1. if counter < 1 then Remove get called
func (tc *TextureCleaner) RemoveUser() {
	tc.counter--
	if tc.counter < 1 {
		tc.Remove()
	}
}

//Remove removes the Texture from opengl memory
func (tc *TextureCleaner) Remove() {
	helper.DeleteTextures(1, &tc.handle)
}
