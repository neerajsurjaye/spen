#version 330 core

layout (location = 0) in vec2 aPos;
layout (location = 1) in vec4 aColor;

uniform mat4 uMVP;

out vec4 vColor;

void main(){
    //Why its a mat4
    vColor = aColor;

    gl_Position = uMVP * vec4(aPos, 0.0, 1.0);
}