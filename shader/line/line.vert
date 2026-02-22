#version 330 core

layout (location = 0) in vec2 aPos;

uniform mat4 uMVP;

void main(){
    //Why its a mat4
    gl_Position = uMVP * vec4(aPos, 0.0, 1.0);
}