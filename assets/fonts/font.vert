#version 120

attribute vec2 position;
attribute vec2 textureCoords;

varying vec2 pass_textureCoords;

uniform vec2 translation;

void main(void){

	gl_Position = vec4(position + translation * vec2(2.0, -2.0), 0.0, 1.0);
	pass_textureCoords = textureCoords;

}