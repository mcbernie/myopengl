package objects

import "github.com/mcbernie/myopengl/graphic/helper"

type TextureCleaner struct {
	handle  uint32
	counter int32
	name    string
}

func CreateTextureCleaner(name string, handle uint32) *TextureCleaner {
	return &TextureCleaner{
		name:   name,
		handle: handle,
	}
}

func (tc *TextureCleaner) AddHandle(handle uint32) {
	tc.handle = handle
}

func (tc *TextureCleaner) AddUser() {
	tc.counter++
}

func (tc *TextureCleaner) RemoveUser() {
	tc.counter--
	if tc.counter < 1 {
		tc.Remove()
	}
}

func (tc *TextureCleaner) Remove() {
	helper.DeleteTextures(1, &tc.handle)
}
