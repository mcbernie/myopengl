#version 120

uniform vec4 color;
varying vec2 _uv;
uniform sampler2D renderedTexture;
uniform float time;

void main()
{
  
  //gl_FragColor = vec4(colour, texture2D(fontAtlas, pass_textureCoords).a);
  //gl_FragColor = texture2D(renderedTexture, _uv.x, _uv.y);
  vec2 a = _uv;
  //a.x = 0.1;
  vec4 t = texture2D( renderedTexture, vec2(a.x + time * 0.18, a.y));
  vec4 nc = color;
  nc.a = t.a;
  gl_FragColor = nc;
  //gl_FragColor = texture2D( renderedTexture, vec2(a.x, a.y));
  //gl_FragColor = vec4(texture2D(renderedTexture, _uv).x, texture2D(renderedTexture, _uv).y, texture2D(renderedTexture, _uv).z, 1.0);
  //gl_FragColor = texture2D(renderedTexture, _uv);
  //gl_FragColor = vec4(0.5,0.8,1.0,1.0);
}