#version 120
	
attribute vec2 position;
attribute vec2 texCoord;

varying vec2 _uv;

uniform mat4 transformationMatrix;
uniform mat4 projectionMatrix;



void main() {
  //gl_Position = vec4(position * vec2(2.0, -2.0), 0.0, 1.0);
  //gl_Position = projectionMatrix * transformationMatrix * vec4(position,0.0,0.5);
	//gl_Position = projectionMatrix * transformationMatrix * vec4(position,0.2,0.5);
  //vec4 tpos = vec4(position,1.0,1.0);
  gl_Position = projectionMatrix * transformationMatrix * vec4(position * vec2(1.0, -1.0),1.0,1.0);

	//_uv = texCoord;

	//vec2 uv = tpos.xy;
  //gl_Position = vec4(position,0.0,1.0);
  
  //vec2 uv = position * 0.5 + 0.5;
	//_uv = vec2(uv.x / 2.0 + 0.5, uv.y  / 2.0 + 0.5);
  _uv = texCoord;
}