# Motion

## Units and Consistency

- Units should be consistent for the system.
- Ex:
    - Distance: meters
    - Time: seconds
    - Mass: kilogram\
       Then
    - Velocity: meters/second
    - Acceleration: meters/second^2
    - Force: Kg \* m / s^2

## Velocity

- Rate of change in position.
    - How position changes over time.
- $velocity = \Delta position / \Delta time$
- $v = dx / dt$

- So to get updated position of objects in space.
    - $postion \mathrel{+}= velocity * dt$
        - $postion = position + velocity * dt$

### Constant Velocity Motion

- Velocity doesn't change over time
    - $position(t) = position_0 + velocity * t$
- Straight line motion.
- No curves.
- Equal distance over time.
- **Assume constant acceleration in one physics step. Even if acceleration exists.**

### Direction & Magnitude

- $velocity = direction * speed$
- Magnitude
    - $speed = |velocity|$
- Direction
    - $direction = velocity / |velocity|$

- This can help in change direction but keep speed constant.

## Accleration

- Defines how velocity changes over time.
- $acceleration = \Delta velocity / \Delta time$
- $a = dv / dt$

So

- $velocity = velocity + accleration * dt$
- Accleration usually comes from forces.
    - Gravity -> constant acceleration
    - Thrust -> Controlled Acceleration

### Constant Acceleration

- Acceleration stays same over time.
    - Velocity changes linearly.
    - $velocity = velocity + acceleration * dt$
    - $position = position + velocity * dt$

### Acceleration and Curvature

- Acceleration perpendicular to motion leads to curvature.
- Acceleration parallel to motion leads to changes to speed.

## Time

- $\Delta t$ is the amount of simulated time passed at every physics step.
    - Without dt. Physics is dependent on how fast the cpu is not on physics.
    - Gives frame rate independence.
- Fixed dt gives deterministic physics.

## Newtons First Law (Intertia)

- A object will remain at rest, or continue moving in a straight line at a constant velocity. Unless aceted upon by an net external force.
    - Motion doesnn't require force to continue.
    - Force is only required to change motion.
    - **Veclotiy doesn't require force**

## Newtons Second Law

- Force on object equals mass times acceleration.
- $F = m \cdot a$

or

- $a = \frac{F}{m}$
    - Force causes acceleration.
    - Mass resists acceleration.

### Inverse mass

- Easier to compute. Multiplication is cheper to compute.
    - So instead of storing mass $m$. Store inverse mass $m^{-1}$
    - So $a = F \times invMass$

- Easy to have infinite mass.
    - If $m = \infty$ then $m^{-1} = 0$
- Avoids divide by 0 bugs.

## Force

- An interaction that changes motion.
- A force is an input that contributes to acceleration over time.

$$a = F \cdot m^{-1}$$

### Continuous Force

- Applied every frame over duration.
    - Gravity
    - Thrust
    - Wind
- Changes velocity gradually

### Instataneous Effects (Impulses)

- Applied once, Changes velocity immediatly
    - Collision
    - Explosions
- $J = \Delta v \cdot m$
- $velocity += impulse * invmass$
- Some forces are applied at very small time step even smaller than dt. So we don't calculate over dt. we directly set the output value.

### Force Accumulation

- Net force = sum of force applied at current timestep.
- Forces are added linearly.
  $$F_{net} = F_1 + F_2 + F_3 + ...$$
- Forces should be summed at a time step then applied to object.
- Old forces from prev frame should not be used. Or say the sum of forces applied on earlier frame should be cleared and new sum should be used.
- Impulse should not be placed in the force sum.

## Gravity

- $g = (0, -9.81, 0) m / s^2$
    - An acceleration not a force.
- Gravity is constant acceleration applied per frame.

### Gravity as a force

- $F_{gravity} = m \cdot g$

The mass cancels out

- $a = F \cdot m^{-1}$

then

- $a = (m \cdot g) m^{-1}$
- $a = g$

## Motion Update Order

- Force -> Acceleration -> Velocity -> Position
- Impulses should be applied before all these.
