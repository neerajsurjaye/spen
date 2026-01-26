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

- For 2D. The width is x. So the camera is centered at (0,0). Basically the cameras range is from $-x/2$ -> $x/2$. Same for height.

- For objects
    - viewPos = worldPos - cameraPos

## Screen Space

- A grid of pixels. Defines location of pixels.
- It is resolution dependent.
- Answers which pixels should light up.

## Converting between spaces

- Matrix multiplication is used to convert between spaces.

## Local to World

- LocalSpace: Say a square with local center at $0,0$ will have its vertices at $x , y$.
- World Space: The cube is actually at $4,5$.
- To get the vertices to world space. We need to account for the scale, rotation and transform on the cube.

- To do this construct the transform matrix.
- $M = T \cdot{} R \cdot{} S$
- Apply to the squres vertices.
    - $v_{world} = M \cdot{} v_{local}$

## World to local

- To move from world to local. We could calculate the inverse of $M$ the transformation matrix.$M^{-1}$

## Parenting

- Child inherit parents transformations. Which could also have parents.
    - $M_{world} = M_{grandparent} \cdot{} M_{parent} \cdot{} M_{local}$

    - This basically means apply ones own transformations. Then apply parent transformation then grandparent transformation.

## Camera

- The camera is not moved. The world is moved opposite to the camera.
    - The world is transformed so camera stays at $0,0,0$ and looks down at $-z$.

- 2D camera defines. Where am I looking how zoomed am I and rotation.
    - The camera is not moved. Ex: if we need to move the camera to right the world is moved to left.

- Note: The objects are themselves are not actually moved they stay in their original location in world space.
    - But while rendering the camera transform. Transforms the whole world so the camera stays at 0,0,0.
- **Camera transform** : Describes camera as a normal object in the world.
    - The inverse of it is taken while rendering to put it in center or $0,0,0$.
- **View Matrix** : Describes how world is seend from the camera. It is the inverse of camera transform.

## World to Screen Conversion 2D

- For 2d the pipeline is as follows.
    - Local Space -> World Space -> View Space -> NDC -> Screen(Pixels)

### NDC: Normalized device coordinates

- NDC is normalized screen space where
    - Left to right is from -1 to 1
    - Bottom to up is from -1 to 1
    - Center is at 0,0
- Visible area is same box regardless of screen size
    - This acts as a common layer between different screen sizes and camera space

### View -> NDC

- Say we have viewWidth and viewHeight.

- So to calculate a point.
    - ndc.x = viewPos.x / (viewWidth / 2)
    - ndx.y = viewPos.y / (viewHeight / 2)
        - Whene viewPos.x ranges from -5 to 5 for view width 10. Similar for viewPos.y
- Ex:
    - Width: 20 units
    - Height: 10 units
    - viewPos: $(5,2.5)$

    Then
    - ndc.x = $5 / 10 = 0.5$
    - ndx.y = $2.5/ 5 = 0.5$

### NDC -> Screen(Pixels)

- Maps -1 -> 1 to 0, W for width and -1 -> to H, 0 (reversed).
- Say screen width = W and screen height = H

Then

- $screenx = (ndc.x + 1) * 0.5 * W$
- $screeny = (1 - ndc.y) * 0.5 * H$
    - **Y is flipped as screen goes downwards. While NDC assues up is positive**

- So fo a (1920 x 1080) screen.
    - NDC (0,0) is (960, 540). which will be the center.

Basically

- NDC (-1 -> 1) and width is 1920
    - When ndc.x + 1 is applied it is (0 -> 2)
    - When 0.5 is multiplied it is (0 -> 1)
    - When 1080 is multiplied it is 0 -> 1920.

### Mouse Picking

- Basically means mouse position on screen to world position.
    - Basically mouse position $(mx, my)$ to world position $(wx, wy)$
- Screen -> NDC -> View -> World

#### Screen to NDC

- If NDC to screen is
    - $screenx = (ndc.x + 1) * 0.5 * W$
    - $screeny = (1 - ndc.y) * 0.5 * H$
- Then
    - $ndc.x = (screenX / W) * 2 - 1$
    - $ndc.y = 1 - (screenY / H) * 2$

#### NDC to View

- IF View to NDC is
    - $ndc.x = viewPos.x / (viewWidth / 2)$
    - $ndc.y = viewPos.y / (viewHeight / 2)$

- Then
    - $viewPos.x = ndc.x * (viewWidth / 2)$
    - $viewPos.y = ndc.y * (viewHeight / 2)$

#### View to world

- If view to world is
    - $viewPos = worldPos - cameraPos$
- Then
    - worldPos = viewPos + cameraPos
