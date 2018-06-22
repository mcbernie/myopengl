package gfx

import "github.com/go-gl/glfw/v3.2/glfw"

var globalWindow *glfw.Window

//CreateWindow create a glfw window and save pointer to global variable
func CreateWindow(width, height int, title string) (*glfw.Window, error) {
	win, err := glfw.CreateWindow(width, height, title, nil, nil)

	globalWindow = win

	return globalWindow, err
}

//GetWindow returns a global static glfw.Window pointer
func GetWindow() *glfw.Window {
	return globalWindow
}
