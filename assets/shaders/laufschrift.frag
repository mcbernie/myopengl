#version 150

uniform vec4 color;
in vec2 _uv;
out vec4 fragColor;

uniform sampler2D renderedTexture;
uniform float time;


void main()
{
  vec2 a = _uv;
  vec4 t = texture( renderedTexture, vec2(a.x + time * 0.18, a.y));
  vec4 nc = color;
  nc.a = t.a;
  fragColor = nc;
}