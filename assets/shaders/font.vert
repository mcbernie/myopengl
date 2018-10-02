#version 150

in vec2 position;
in vec2 textureCoords;

out vec2 pass_textureCoords;

uniform vec2 translation;

uniform mat4 transformationMatrix;
uniform mat4 projectionMatrix;


void main(void){

	gl_Position = projectionMatrix * vec4((translation.x + position.x), (translation.y + position.y), 1.0, 1.0) ;
	//gl_Position = vec4(position,0.0,1.0);

	pass_textureCoords = textureCoords;

}