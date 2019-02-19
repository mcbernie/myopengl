package gui

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/golang-ui/nuklear/nk"
	"github.com/mcbernie/myopengl/graphic/objects"
)

const (
	winWidth  = 400
	winHeight = 500

	maxVertexBuffer  = 512 * 1024
	maxElementBuffer = 128 * 1024
)

type Option uint8

const (
	Easy Option = 0
	Hard Option = 1
)

type State struct {
	bgColor   nk.Color
	prop      int32
	opt       Option
	pauseText string

	delay_bytes    []byte
	delay_len      int32
	duration_bytes []byte
	duration_len   int32
}

type GuiSystem struct {
	ctx             *nk.Context
	win             *glfw.Window
	state           *State
	currentDuration float64
	font            *nk.Font
	fontHandle      *nk.UserFont
	visible         bool
}

func (gui *GuiSystem) Render(r *objects.Renderer, deltaTime float64) {
	if gui.visible != true {
		return
	}

	nk.NkPlatformNewFrame()
	nk.NkStylePushFont(gui.ctx, gui.fontHandle)

	bounds := nk.NkRect(50, 50, 230, 250)

	update := nk.NkBegin(gui.ctx, "Star Entertainer 2.0", bounds,
		nk.WindowBorder|nk.WindowTitle|nk.WindowMovable)

	if update > 0 {
		nk.NkLayoutRowStatic(gui.ctx, 30, 80, 1)
		{
			if nk.NkButtonLabel(gui.ctx, "button") > 0 {
				log.Println("[INFO] button pressed!")
			}
		}
		/*nk.NkLayoutRowDynamic(gui.ctx, 30, 2)
		{
			if nk.NkOptionLabel(gui.ctx, "easy", flag(gui.state.opt == Easy)) > 0 {
				gui.state.opt = Easy
			}
			if nk.NkOptionLabel(gui.ctx, "hard", flag(gui.state.opt == Hard)) > 0 {
				gui.state.opt = Hard
			}
		}*/
		nk.NkLayoutRowDynamic(gui.ctx, 25, 1)
		{
			nk.NkPropertyInt(gui.ctx, "Compression:", 0, &gui.state.prop, 100, 10, 1)
		}
		nk.NkLayoutRowDynamic(gui.ctx, 20, 1)
		{
			nk.NkLabel(gui.ctx, "background:", nk.TextLeft)
		}
		nk.NkLayoutRowDynamic(gui.ctx, 25, 1)
		{
			/*size := nk.NkVec2(nk.NkWidgetWidth(gui.ctx), 400)
			if nk.NkComboBeginColor(gui.ctx, gui.state.bgColor, size) > 0 {
				nk.NkLayoutRowDynamic(gui.ctx, 120, 1)
				gui.state.bgColor = nk.NkColorPicker(gui.ctx, gui.state.bgColor, nk.ColorFormatRGBA)
				nk.NkLayoutRowDynamic(gui.ctx, 25, 1)
				r, g, b, a := gui.state.bgColor.RGBAi()
				r = nk.NkPropertyi(gui.ctx, "#R:", 0, r, 255, 1, 1)
				g = nk.NkPropertyi(gui.ctx, "#G:", 0, g, 255, 1, 1)
				b = nk.NkPropertyi(gui.ctx, "#B:", 0, b, 255, 1, 1)
				a = nk.NkPropertyi(gui.ctx, "#A:", 0, a, 255, 1, 1)
				gui.state.bgColor.SetRGBAi(r, g, b, a)
				nk.NkComboEnd(gui.ctx)
			}*/

			nk.NkLabel(gui.ctx, "Pause in Sek:", nk.TextLeft)
			nk.NkEditString(gui.ctx, nk.EditSimple, gui.state.delay_bytes, &gui.state.delay_len, 64, nk.NkFilterDecimal)

			nk.NkLabel(gui.ctx, "Übergang in Sek:", nk.TextLeft)
			nk.NkEditString(gui.ctx, nk.EditSimple, gui.state.duration_bytes, &gui.state.duration_len, 64, nk.NkFilterDecimal)
		}
	}
	nk.NkEnd(gui.ctx)

	nk.NkStylePopFont(gui.ctx)

	// Render
	bg := make([]float32, 4)
	nk.NkColorFv(bg, gui.state.bgColor)

	nk.NkPlatformRender(nk.AntiAliasingOn, maxVertexBuffer, maxElementBuffer)

}

func (gui *GuiSystem) Delete() {
	log.Println("shutdown gui")
	nk.NkPlatformShutdown()
}

func CreateGui(win *glfw.Window) *GuiSystem {

	log.Println("start gui")
	gui := GuiSystem{
		win: win,
		ctx: nk.NkPlatformInit(win, nk.PlatformInstallCallbacks),
		state: &State{
			bgColor:        nk.NkRgba(28, 48, 62, 255),
			duration_bytes: []byte{0, 64},
			duration_len:   0,
			delay_bytes:    []byte{0, 64},
			delay_len:      0,
		},
		visible: false,
	}

	atlas := nk.NewFontAtlas()
	nk.NkFontStashBegin(&atlas)
	gui.font = nk.NkFontAtlasAddFromFile(atlas, "assets/fonts/verdana.ttf", 16, nil)
	nk.NkFontStashEnd()
	gui.fontHandle = gui.font.Handle()

	return &gui
}

func (gui *GuiSystem) GetPause() int32 {
	s := string(gui.state.delay_bytes[:gui.state.delay_len])
	r, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return 0
	}
	return int32(r)
}

func (gui *GuiSystem) GetDuration() int32 {
	s := string(gui.state.duration_bytes[:gui.state.duration_len])
	r, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return 0
	}
	return int32(r)
}

func (gui *GuiSystem) SetDuration(duration int32) {
	s := fmt.Sprintf("%d", duration)
	b := []byte(s)

	gui.state.duration_bytes = b
	gui.state.duration_len = int32(len(s))
}

func (gui *GuiSystem) SetVisible(visible bool) {
	gui.visible = visible
}

func (gui *GuiSystem) ToggleVisible() {
	gui.visible = !gui.visible
}

func b(v int32) bool {
	return v == 1
}

func flag(v bool) int32 {
	if v {
		return 1
	}
	return 0
}
