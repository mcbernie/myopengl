#version 410 core

layout (location = 0) in vec3 position;
layout (location = 1) in vec2 texCoord;

out VertexData {
  vec2 TexCoord;
} VertexOut;

uniform mat4 world;
uniform mat4 view;
uniform mat4 project;

uniform float posx;
uniform float posy;

void main()
{
    gl_Position = project * view * world * vec4(position, 1.0);
    VertexOut.TexCoord = texCoord;    // pass the texture coords on to the fragment shader      
}