#version 150

in vec4 gl_Vertex;
out vec4 gl_Position;

void main()
{
  gl_Position = gl_Vertex;  
}