#version 150
	
in vec2 position;
in vec2 texCoord;

out vec2 _uv;

uniform mat4 transformationMatrix;
uniform mat4 projectionMatrix;

void main() {
  gl_Position = projectionMatrix * transformationMatrix * vec4(position * vec2(1.0, -1.0),1.0,1.0);
  _uv = texCoord;
}