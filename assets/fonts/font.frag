#version 120

varying vec2 pass_textureCoords;

uniform vec3 colour;
uniform sampler2D fontAtlas;

void main(void){

	//gl_FragColor = vec4(colour, texture2D(fontAtlas, pass_textureCoords).a);
	gl_FragColor = texture2D(fontAtlas, pass_textureCoords);
}