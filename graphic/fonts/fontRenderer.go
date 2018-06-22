package fonts

import (
	"github.com/mcbernie/myopengl/glHelper"
)

type fontRenderer struct {
	shader *FontShader
}

func createFontRenderer() *fontRenderer {
	shader := CreateFontShader()
	//
	return &fontRenderer{
		shader: shader,
	}
}

func (fr *fontRenderer) render(texts TextList) {
	fr.prepare()

	for font, innerTexts := range texts {
		glHelper.ActiveTexture(glHelper.GlTexture0)
		glHelper.BindTexture(glHelper.GlTexture2D, font.textureAtlas)
		for _, text := range innerTexts {
			fr.renderText(text)
		}
		glHelper.BindTexture(glHelper.GlTexture2D, 0)
	}

	fr.endRendering()
}

func (fr *fontRenderer) prepare() {
	//glHelper.Enable(glHelper.GlBlend)
	//glHelper.BlendFunc(glHelper.GlSrcAlpha, glHelper.GlOneMinusSrcAlpha)
	//glHelper.Disable(glHelper.GlDepthTest)
	fr.shader.Use()

}

func (fr *fontRenderer) renderText(text *GUIText) {
	glHelper.BindVertexArray(text.getMesh())
	glHelper.EnableVertexAttribArray(0)
	glHelper.EnableVertexAttribArray(1)

	fr.shader.SetColur(text.colour)
	fr.shader.SetTranslation(text.position)
	glHelper.Uniform1i(fr.shader.GetUniform("fontAtlas"), 0)

	glHelper.DrawTrianglesArray(0, text.vertexCount)

	glHelper.DisableVertexAttribArray(0)
	glHelper.DisableVertexAttribArray(1)

	glHelper.BindVertexArray(0)

}

func (fr *fontRenderer) endRendering() {
	fr.shader.UnUse()
	//glHelper.Disable(glHelper.GlBlend)
	//glHelper.Enable(glHelper.GlDepthTest)
}
