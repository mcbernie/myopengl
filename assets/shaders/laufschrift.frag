#version 120

uniform vec4 color;
varying vec2 _uv;
uniform sampler2D renderedTexture;
uniform float time;

void main()
{
  vec2 a = _uv;
  vec4 t = texture2D( renderedTexture, vec2(a.x + time * 0.18, a.y));
  vec4 nc = color;
  nc.a = t.a;
  gl_FragColor = nc;
}