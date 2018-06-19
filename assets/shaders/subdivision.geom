#version 410 core

layout (triangles) in;
layout (triangle_strip, max_vertices=256) out; 

in VertexData {
  vec2 TexCoord;
} VertexIn[];

out VertexData {
  vec2 TexCoord;
} VertexOut;

uniform int sub_divisions;
uniform mat4 world;

void main() {
  vec4 v0 = gl_in[0].gl_Position;
  vec2 t0 = VertexIn[0].TexCoord;
  
  vec4 v1 = gl_in[1].gl_Position;
  vec2 t1 = VertexIn[1].TexCoord;

  vec4 v2 = gl_in[2].gl_Position;
  vec2 t2 = VertexIn[2].TexCoord;

  float dx = abs(v0.x-v2.x)/sub_divisions;
  float dz = abs(v0.z-v1.z)/sub_divisions;

  float dtx = abs(t0.x-t2.x)/sub_divisions;
  float dty = abs(t0.y-t1.y)/sub_divisions;

  float x=v0.x;
  float z=v0.z;

  float tx=t0.x;
  float ty=t0.y;

  for(int j=0;j<sub_divisions*sub_divisions;j++) {
    gl_Position =  world * vec4(x,0,z,1);
    VertexOut.TexCoord = vec2(tx , ty);
    EmitVertex();
    gl_Position =  world * vec4(x,0,z+dz,1);
    VertexOut.TexCoord = vec2(tx , ty+dty);
    EmitVertex();
    gl_Position =  world * vec4(x+dx,0,z,1);
    VertexOut.TexCoord = vec2(tx +dtx , ty);

    EmitVertex();
    gl_Position =  world * vec4(x+dx,0,z+dz,1);
    VertexOut.TexCoord = vec2(tx +dtx , ty + dty);

    EmitVertex();
    EndPrimitive();
    x+=dx;
    tx+=dtx;
    if((j+1) %sub_divisions == 0) {
      x=v0.x;
      tx=t0.x;
      
      z+=dz;
      ty+=dty;
    }
  }
}