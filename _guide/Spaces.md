## Space

## Local Space

- Object own corrdinate system
- As objects move rotate scale. Their local space changes.
- Helps in parent child relationship.

Ex: Object roates around its origin. Not the world origin.

## World Space

- Global Coordinate system

## View Space(Camera Space)

- Camera is at origin looking at fixed direction.
    - Everything else is transformed relative to the camera.
- Eases calculation for rendering.
  Ex: Say for ray tracking the ray originates from origin and moves towards the objects.

- It is used because rendering is camera centeric not world centeric.

## Screen Space

- A grid of pixels. Defines location of pixels.
- It is resolution dependent.
- Answers which pixels should light up.

## Converting between spaces

- Matrix multiplication is used to convert between spaces.
