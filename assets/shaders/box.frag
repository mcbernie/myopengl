#version 120

uniform vec4 color;
varying vec2 _uv;

uniform sampler2D tex;

uniform vec2 u_dimensions;
uniform vec2 u_border;

float map(float value, float originalMin, float originalMax, float newMin, float newMax) {
    return (value - originalMin) / (originalMax - originalMin) * (newMax - newMin) + newMin;
}

// Helper function, because WET code is bad code
// Takes in the coordinate on the current axis and the borders 
float processAxis(float coord, float textureBorder, float windowBorder) {
    if (coord < windowBorder)
        return map(coord, 0, windowBorder, 0, textureBorder) ;
    if (coord < 1 - windowBorder) 
        return map(coord,  windowBorder, 1 - windowBorder, textureBorder, 1 - textureBorder);
    return map(coord, 1 - windowBorder, 1, 1 - textureBorder, 1);
}

void main(void) {
    vec2 newUV = vec2(
        processAxis(_uv.x, u_border.x, u_dimensions.x),
        processAxis(_uv.y, u_border.y, u_dimensions.y)
    );

    gl_FragColor = texture2D(tex,newUV);
}