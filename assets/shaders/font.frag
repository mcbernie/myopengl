#version 150

in vec2 pass_textureCoords;

uniform vec3 colour;
uniform sampler2D fontAtlas;

out vec4 color;

void main(void){

	color = vec4(colour, texture(fontAtlas, pass_textureCoords).a);
	//gl_FragColor = vec4(colour, texture2D(fontAtlas, pass_textureCoords).a + 0.2);
	//gl_FragColor = texture2D(fontAtlas, vec2(0.0,0.0));
  //gl_FragColor = texture2D(fontAtlas,pass_textureCoords);
	
}