#version 120

attribute vec2 position;
attribute vec2 textureCoords;

varying vec2 pass_textureCoords;

uniform vec2 translation;

void main(void){

	//gl_Position = vec4(position,0.0,1.0);
	gl_Position = vec4(position + translation * vec2(2.0, -2.0), 0.0, 1.0);

	/*gl_Position = projectionMatrix * transformationMatrix * vec4(position,0.0,1.0);*/
	//vec2 uv = position * 0.5 + 0.5;
	//pass_textureCoords = vec2(uv.x, 1.0 - uv.y);
	//pass_textureCoords = vec2(0.00390625,0.00390625);
	pass_textureCoords = textureCoords;

}