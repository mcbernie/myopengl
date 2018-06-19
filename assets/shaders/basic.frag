#version 410 core

//in vec2 TexCoord;
in VertexData {
  vec2 TexCoord;
} VertexIn;

out vec4 color;

uniform sampler2D ourTexture0;
uniform sampler2D ourTexture1;
uniform sampler2D ourTexture2;


void main()
{
   color = texture(ourTexture0, VertexIn.TexCoord);
    //color = vec4(1,1,1,1);
}