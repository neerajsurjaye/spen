#version 330 core

layout (location = 0) in vec2 aPos;

out vec2 vPos;

uniform mat4 uMVP;
uniform float uRadius;

// uniform mat4 uMVP;

void main(){
    vec2 scaled = aPos * uRadius;
    vPos = aPos;

    // gl_Position = uMVP * vec4(aPos, 0.0, 1.0);
    gl_Position = uMVP * vec4(scaled, 0.0, 1.0);
}