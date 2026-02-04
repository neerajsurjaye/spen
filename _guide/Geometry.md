# Geometry

## Foundation

### Point vs Vector

- Point represents a position in space. Just a location.
    - P = (x , y)
- Vector represents direction and magnitude. It answers how and in which direction.
    - V = (dx, dy)

- Point + Vector : move the point
- Point - Point : A vector representing distance between them.
- Vector + Vector : Combines Movement
- Point + Point : Not possible
- Vector - Vector: Difference in movement

### Coordinate System

- Answers where is (0,0)? Which way is x and y? And what does one unit of x and y mean?

### Squared distance (distance²)

- Distance is calculated by
    - $d = \sqrt{(x_2 - x_1)^2 + (y_2 - y_1)^2}$
- Squared distance removes the square root.
    - $d^2 = (x_2 - x_1)^2 + (y_2 - y_1)^2$

- Square Root is expensive. While comparing distance squre root is not required.
    - $d_1^2 > d_2^2 \implies{} d_1 > d_2$

## Lines and Line Segments

### Lines Parametric form

- Line is represented as
    - $L(t) = P + tD$
    - Where
        - P is a point on line.
        - D Direction vector.
        - t is scaler parameter. Where t can be any real number. Putting a value here gives a point magnitude of D times away from point P.
- Also the equation is written as component form
    - Where\
      $P = (P_x, P_y)$\
      $D = (D_x, D_y)$
    - Then\
       $x(t) = P_x + tD_x$\
       $y(t) = P_y + tD_y$

### Line vs Line Intersection

- Say we have two lines
    - $L_1(t) = P + tD$
    - $L_2(u) = Q + uE$
- They intersect if for a t and u.
    - $P + tD =  Q + tE$
    - $tD - uE = Q - P$
- Eliminating U
    - $(tD - uE) \times{} E = (Q - P) \times{} E$
    - $t(D \times{} E) - u(E \times{} E) = (Q - P) \times{} E$
        - $E \times{} E = 0$
    - $t(D \times{} E) = (Q - P) \times{} E$
    - $t = \frac{(Q - P) \times E}{D \times{} E}$
    - Then intersection point can be found by
        - $Intersection = P + tD$
    - If Lines are parellel the denominator will be zero.
        - $D \times{} E = 0$
            - Parallel or Collinear

- **Note: Cross product in 2D is scaler. As the Vector pointing outwards points in z axis. Which for 2D we don't care about.**

### Segment

- Finite portion of line between two points.\
  $S(t) = A + t(B - A) \in [0,1]$
    - Where A and B are two points. $B - A$ gives a vector starting from A ending at B.

### Segment vs Segment Intersection

- $S_1(t) = A + t(B -A) , t \in [0,1]$
- $S_2(t) = C + u(D - C), u \in [0,1]$

Find values of t and u where these two intersect. If t and u are in range $[0,1]$. Then these segments intersect.

### Closest Point from a Line

Say a point $P$ and a line with point $A$ and direction $D$.

- The line can be defined by\
   $Q = A + tD$
- If drawing a vector from point $P$ to line\
   $P - Q = P - (A + tD)$
- At closest point\
   $(P -Q) \cdot{} D = 0$
    - As the line is perpendicular dot prod is 0.
- Substituting\
   $(P - (A + tD)) \cdot{} D = 0$\
- Then
  $t = \frac{(P - A) \cdot{} D}{D \cdot{} D}$
- The closes point will can be now found by placing the value of t in below formula.\
   $Q = A + tD$

### Closes point on a Segment

- Say a Segment and point\
   $S(t) = A + t(B - A), t \in [0,1]$\
  $P$

- let D = (B -A)
- Then
  $t = \frac{(P - A) \cdot{} D}{D \cdot{} D}$\
- Now the value of t should be clamped. Basically if greater than 1 then it is 1 and less than 0 then it is zero.
  $t = max(0, min(1,t))$
- The closes point will can be now found by placing the value of t in below formula.\
   $Q = A + tD$

### Distance Point -> [Line , Segment]

- Use the previous formula to find shortest vector. And use its magnitude. Its the minimum distance.

## Rays

- Ray represents a half line. Originates from a point and extends to infinity in one direction.
- $R(t) = O + tD, t \in [0, \infty]$
    - Or say $t \ge 0$
    - D is normalized
- Usage: ray casting, ray tracing etc.

### Rays vs Point

- Same as line vs point. Just $t \ge 0$

### Rays vs Segment

- Same as line vs segment. Just $t \ge 0$

### Rays vs Circle

- Same as line vs circle. Just $t \ge 0$

### Rays vs AABB

- Given a ray
    - $R(t) = O + tD, t \in [0, \infty]$
- And a AABB
    - $Bmin = (xmin, ymin)$
    - $Bmax = (xmax, ymax)$

- Think of box as x slab and y slab. Rays enters each box at some t and exits at some t.
    - If there is some t where t is inside both slabs at same time then there is intersection.
- Say $R(t) = O + tD, t \in [0, +\infty]$
- In x-dimenstion
    - $x(t) = Ox + t \cdot{} Dx, t \in [0, +\infty]$
    - To find intersection at xmin
    - $xmin = Ox + t \cdot{} Dx$
    - $t = (xmin - Ox) / Dx$
- Similarly t should be found for other coordinate and other point.
    - Say the output is tx1, tx2, ty1, ty2
- $txmin = min(tx1, tx2)$
- $txmax = max(tx1, tx2)$
- $tymin = min(ty1, ty2)$
- $tymax = max(ty1, ty2)$

- Overlapping condition
    - $tEnter = max(txmin, tymin)$
    - $tExit = min(txmax, tymax)$
- Overlapping interval
    - tEnter = max(txmin,tymin)
    - tExit = min(txmax, tymax)
- Intersection Condition
    - $tExit \ge tEnter$
    - and
    - $tExit \ge 0$
        - $tExix \ge 0$ because intersection should be in fron of ray.

## Circles

- Represented by two things. A point $C = (x, y)$ and a radius $r$.

- All points on the circle satisfy and **Standard Equation**.
    - $(x - c_x)^2 + (y - c_y)^2 = r^2$
- Vector form
    - $|| P - C|| = r$\
      or
    - $||P - C|| ^ 2 = r ^2 $
        - Easier to compute
- Parametric form
    - $x = c_x + r \cos \theta$
    - $y = c_y + r \sin \theta$
        - Where $\theta \in [0, 2\pi)$

### Point inside a circle

- If distance between point P and circle center C is smaller than r. Then its inside.

### Circle vs Circle collision

- Say two circles with center $C_1 \& C_2$ and radius $r_1 \& r_2$.
- For collision
    - $||C_1 - C_2|| < r_1 + r_2$\
      or
    - $||C_1 - C_2||^2 < (r_1 + r_2)^2$

    - If equal on both sides then touching. If distance is smaller then sum of both radius then overlapping.

#### Collision normal

- $n = \frac{C_2 - C_1}{||C_2 - C_1||}$
    - Points from C_1 to C_2.

#### Penetration Depth

- $penetration = (r_1 + r_2) - ||C_2 - C_1||$

### Circle vs Line

- A circle collides with a line if shortest distance beween line a circles center is smaller than radius.
    - $distance(C, line) \le r$
- To find the distance. Find perpendicular from line to circles center.

### Circle vs Segment

- Find distane between center point and segment. If distance is smaller than radius. Then its collision.

### Circle vs AABB

- Find distance from center. Then check if distane is smaller than radius.

### Closet Point on Circle from a Point

- Find direction from center to point.
    - $\vec{d} = P - C$
- Normalize it
    - $\hat{d} = \frac{\vec{d}}{||\vec{d}||}$
- Move from center by radius.
    - $Q = C + r \cdot{} \hat{d}$
        - Where r is radius.

## AABB Representation

- AABB -> Axis algined bounding box.
    - The alignement to axis basically no rotation. Makes the calculation very easy.

- Represented using min max form.
    - $min = (x_{min}, y_{min})$
    - $max = (x_{max}, y_{max})$
    - min = Bottom left corner
    - max = Top right corner

### Point inside AABB

- A point ($P = (x,y)$) lies inside if below condition is true.
    - $x_{min} \le x \le x_{max} \land y_{min} \le y \le y_{max}$
    - The above also mean if it is on boundy it is inside.

### AABB vs AABB

- Let there be two bounding box A and B.

- If
    - $A.xmin \le B.xmax \land A.xmax \ge B.xmin$
    - and
    - $A.ymin \le B.ymax \land A.ymax \ge B.ymin$

### Swept AABB

- Good for fast moving object.
- We take two AABB one is static and one is moving. The moving one is reduced to a point. And the static ones size is increased by half the size of moving one.
    - A ray is cast from moving AABB in the direction of movement. And checked if collision will happen.

### Circle vs AABB

- Find the point on AABB closest to the circle center.
    - $closest.x = clamp(cx, box.min.x, box.max.x)$
    - $closest.y = clamp(cy, box.min.y, box.max.y)$
        - The clamp method sets the value to second or third param. If the first value is not in the range of those values.
        - Basically say cx is greater than max.x then output would be max.x. The corner of box on x axis.
- Now mesure distance from circle center to that point.
    - $dx = cx - closest.x$
    - $dy = cy - closest.y$
    - $dx \times dx + dy \times dy \le r \times r$
        - Comparing squared distance.

## Polygons

### Convex Polygons

- If all angles are smaller than or equal to 180 deg.
- Every line segment connecting two points is inside the polygon.

### Shoelace formula

- Calculates signed area of polygon.
    - If going CCW it is positive else negative.

- $$Area = \frac{1}{2} \sum_{i=0}^{n-1}(x_i y_{i + 1} - x_{i+1}y_i)$$

### CCW (Counter Clock Wise)

- Ordering of polygon matters. For example while calculating area with solace CCW produces positive result else negative.
- Defines normal direction. CCW defines front of polygon. Basically normal direction or face orientation.
- CCW should not be mixed with CW(Clockwise)

### Edge normals

- Edges connects two vertices. These have two normals in 2d. Outwards and inwards.

### Point inside convex polygon

- Each edge creates a half plane. We just need to check. If the point is inside all those half planes.

Can be tested with below equation.

- $d = (p - v_i) \cdot n_i$
    - p is the point.
    - $v_i$ is a point on the edge.
    - $n_i$ is the normal direction.
    - We are substracting the point from a point on edge.
        - This gives a vector pointing from edge to point.
    - If its in opposite direction of normal d will be smaller than zero. meaning its inside.
    - $d > 0 $ outside
    - $d < 0 $ inside
    - $d = 0$ On edge
- This is need to be check with every edge.

### Centeroid

- Centeroid is polygons center of mass.
    - If uniformely dense.

## Perpendiculars

- Two vectors are perpendicular if they are at 90deg's.
    - $\vec{a} \cdot \vec{b} = 0$

- Given a 2D vector $\vec{v} = (x, y)$
    - $\vec{a}_\perp = (-y , x)$
    - $\vec{a}_\perp = (y , -x)$

- If using CCW use $(-y , x)$ for outward normals.
- **Normals should be normalized.**

- **Contact Normal** - The normal which is used/chosen during collision.
    - Say a square. On which a square collided on right edge. The right edges normal is contact normal.
    - Or say contact normal is edge normal with minimum penetration.

## Projecting Point onto Axis

- Basically means take a point P and an axis a. Tell how much of point p lies on axis.
    - **Drop a perpendicular from point to axis and mesure wheer it lands**.
- The axis should be normalized.

- $proj = P \cdot \hat{a}$

- To get the actual projected point on axis.
    - $P_{proj} = (P \cdot \hat{a})\hat{a}$

### Projecting Shape on Axis

- Take every point of a shape and project it on the axis. Then take the minimum and maxium on the axis.
- For circle
    - $c = C \cdot \hat{a}$
    - $[min, max] = [c - r, c + r]$

- If say
    - $A = [a_{min}, a_{max}]$
    - $B = [b_{min}, b_{max}]$

    - The overlap condition\
      $a_{max} \ge b_{min} \land b_{max} \ge a_{min}$

- If overlapp happens
    - $penetration = min(a_{max}, b_{max}) - max(a_{min}, b_{min})$

## SAT (Seprating Axis Theorem)

- If two objects are not colliding. Then there is at least one axis where their projections do not overlap.
    - This axis is called seprating axis.

- Otherwise the objects are colliding.

- For convex polygons, Only edge normals can seprate them.
    - If two polygons. A and B with edges M and N. Having normals m and n. Then total axis to check is m + n.

- Polygon vs Circle: In case of circle, take a normal from circle center to the closest vertex on polygon. Use it as a axis for SAT.

- AABB vs AABB for these only two axis exists. x and y axis.

- The axis with minimum penetration is the collision normal.

### MTV (Minimum Translation Vector)

- $MTV = {collision \space normal} \times {penetration \space depth} $

- Tells the minimum amount to move both the objects to seprate both objects.

- Also when applying MTV. Move objects by the amount on how movable they are. Usually inverse mass.
    - If static object. Mass is infinity and inverse mass is 0.

### Collision Geometry Outputs

## Contact Point Calculation

- The point where shapes touch.
- Single contact point can cause roatational jitter so multiple contact points can be used.

### AABB vs AABB

- Compute overlap on x and y.
- Collision normal = axis with minimum overlap.
- Contact point = Center of overlapping region on axis.
    - Find the point on AABB A that is closest to center of B.
    - $contact.x = clamp(centerB.x, minA.x, maxA.x)$
    - $contact.y = clamp(centerB.y, minA.y, maxA.y)$

### Circle vs Circle

- $normal = normalize(centerB - centerA)$
- $contact = centerA + normal * radiusA$
    - Basically start of center of circle A and move along collsion normal radiusA times.

### Polygon vs Polygon

- Use reference edge. The edge which gave normal for MTV.
- Project incident polygon vertices on that edge.
- Pick deepest penetrating vertex.

## Collision Normal Selection

- Tells which direction to push objects apart.
- Which direction impulse act.
- Normal comes from from the axis with minimum penetration depth.

## Parallel Near Parallel Cases

- Check with cross product.
- $prod = cross(e_1, e_2)$
- $if(abs(prod) < epsilon)$
    - epsilon = 1e+9

## Degenerate Edges

- Edges with very small length.
- Skip them. Say $len(edge) < epsilon$ skip them
