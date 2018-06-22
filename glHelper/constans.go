package glHelper

import (
	"github.com/go-gl/gl/v2.1/gl"
)

const (
	GlColorBufferBit = gl.COLOR_BUFFER_BIT

	//GlFragmentShader Define Shader as Fragment Shader
	GlFragmentShader     = gl.FRAGMENT_SHADER
	GlVertexShader       = gl.VERTEX_SHADER
	GlCompileStatus      = gl.COMPILE_STATUS
	GlLinkStatus         = gl.LINK_STATUS
	GlClampToEdge        = gl.CLAMP_TO_EDGE
	GlTextureWrapR       = gl.TEXTURE_WRAP_R
	GlTextureWrapS       = gl.TEXTURE_WRAP_S
	GlTextureMinFilter   = gl.TEXTURE_MIN_FILTER
	GlTextureMagFilter   = gl.TEXTURE_MAG_FILTER
	GlLinearMipmapLinear = gl.LINEAR_MIPMAP_LINEAR
	GlTextureLodBias     = gl.TEXTURE_LOD_BIAS
	GlLinear             = gl.LINEAR

	//GlTexture0 all Textures..
	GlTexture0 = gl.TEXTURE0
	GlTexture1 = gl.TEXTURE1
	GlTexture2 = gl.TEXTURE2
	GlTexture3 = gl.TEXTURE3
	GlTexture4 = gl.TEXTURE4
	GlTexture5 = gl.TEXTURE5
	GlTexture6 = gl.TEXTURE6
	GlTexture7 = gl.TEXTURE7

	GlTexture2D = gl.TEXTURE_2D
	GlSrgbAlpha = gl.SRGB_ALPHA
	GlRgbA      = gl.RGBA

	GlUnsignedByte = gl.UNSIGNED_BYTE

	GlBlend            = gl.BLEND
	GlSrcAlpha         = gl.SRC_ALPHA
	GlOneMinusSrcAlpha = gl.ONE_MINUS_SRC_ALPHA
	GlDepthTest        = gl.DEPTH_TEST

	GlTriangles = gl.TRIANGLES
)
