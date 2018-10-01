package fonts

import "github.com/mcbernie/myopengl/graphic/helper"

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
		helper.ActiveTexture(helper.GlTexture0)
		helper.BindTexture(helper.GlTexture2D, font.textureAtlas)
		for _, text := range innerTexts {
			fr.renderText(text)
		}
		helper.BindTexture(helper.GlTexture2D, 0)
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
	helper.BindVertexArray(text.getMesh())
	helper.EnableVertexAttribArray(0)
	helper.EnableVertexAttribArray(1)

	fr.shader.SetColur(text.colour)
	fr.shader.SetTranslation(text.position)
	helper.Uniform1i(fr.shader.GetUniform("fontAtlas"), 0)

	helper.DrawTrianglesArray(0, text.vertexCount)

	helper.DisableVertexAttribArray(0)
	helper.DisableVertexAttribArray(1)

	helper.BindVertexArray(0)

}

func (fr *fontRenderer) endRendering() {
	fr.shader.UnUse()
	//glHelper.Disable(glHelper.GlBlend)
	//glHelper.Enable(glHelper.GlDepthTest)
}
