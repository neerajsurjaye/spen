## Vectors

Represent direction and magnitude. Can be used to calcuate.

-   Direction and magnitude of froces
-   Loction of point in space
-   Distance between poins in space.

Vector is composed of components which define the magnitude in specific coordinate axis.

-   For 2d vector say $\vec{A}$ = [x, y]. x and y are its components.

### Vector Addition

Adding two vectors produces a third vector. Which when both vectors point in same direction helps increase the magnitude and when in opposite direction reduces the magnitude.
The magnitude and direction of input vectors also define the direction of output vector.

### Scaler Multiplicatoin

Changes the magnitude of the vector. Doesn't change direction if not multiplied by a negative.

### Vector Substraction

Non commutative. Its just vector addition with the second vector multiplied with -1.

The output of substraction of two vectors. Give a direction vector. If I have two vectors $\vec{A}$ and $\vec{B}$ then $\vec{B}$ - $\vec{A}$ gives direction from A to B.

The magnitude of this is distance between these two points.

### Unit Vector

Vector of magnitude(Length) 1.

### Vector Normalizatoin

Converting a vector to a vector of magnitude 1. So the vector only contains info about the direction.

if $\vec{A}$ is a vector its unit vector can be calculated by. $\vec{U} = \frac {\vec{A}} {\|{\vec{A}}\|}$

### Dot Product(Scaler Product)

If

-   $\vec{A}$ = [X, Y]

-   $\vec{B}$ = [M, N]

Then dot product is

-   dot = $\vec{A} * \vec{B} = (X * M) + (Y * N)$

    or

-   $\vec{A} * \vec{B} = |\vec{A}| |\vec{B}| cos{\theta{}}$

Helps in

1. Mesures how much a vector points to the direction of other.

    - Positive value = same direction
    - Zero = Perpendicular
    - Negative Value = Opposite direction

2. Can also help find angle between two vectors.

    - $cos{\theta{}} = \frac{\vec{A} * \vec{B}}{|\vec{A}| |\vec{B}|}$

3. Checking perpendicularity
    - $\vec{A} * \vec{B} = 0$

Some other examples would be light intensity(Angle between surface normal and light drection), collision reponse etc.

### Vector projection

Finds projection of vector A onto Vector B.

-   $proj_BA = (\vec{A} * \hat{B})\hat{B}$
-   $proj_BA = \frac{(\vec{A} * \vec{B})}{|B|^2}\vec{B}$

Basically it says. Give dot product of A on normalized B so magnitude of B doesn't effect the shadow/projections magnitude. So we get the magnitude of shadow of A on B. Then multiply by normalized B. So, we get direction of shadow.

-   It is done in this way cause for A . B dot product gives. Shadow magnitude of A \* magnitude of B. But if B is normalized magnitude of shadow of A is only taken as magnitude of B is 1.

### Vector rejection(Tangent Component)

Say we have vector A and B.\
Then vector rejection would be the vector that starts from proj of A onto B and ends at A.

-   $\vec{rej} = \vec{A} - proj_B(\vec{A})$

Used for collision response, sliding on slopes etc

### Reflection vector

Say we have a suface with a normal vector $\hat{n}$. And a incident vector $\vec{A}$. Then to calculate reflection vector.

-   Take projection along the normal $proj_nA$. This will point downwards.
-   Take rejection by substracting $\vec{A}$ from $proj_nA$.
    -   This will give a vector from projection to vector A. Basically vector containg direction of ray without downwards vector.
-   Now reverse projection vector so it points upwards.
-   Add rejection and inverted projection.

So

-   $\vec{ref} = rej_nA - proj_nA$
-   $\vec{ref} = (\vec{A} - proj_nA) - proj_nA$
-   $\vec{ref} = \vec{A} - 2proj_nA$
-   $\vec{ref} = \vec{A} - 2proj_nA$
-   $\vec{ref} = \vec{A} - 2(\vec{A} * \hat{n})\hat{n}$

### Suface normals

Unit vector perpendicular to a suface.\
Helps calculate reflection, collision, brightness of object etc.

### Decomposing Velocity

Means splliting veclotiy into two vectors along a unit vector. The two vectors are perpendicular and tangent along the unit vector.\
Ex: The unit vector can be normal of a surface and velocity of object can be taken. This can calculate how fast a object is sliding along a surface and how fast the object is moving into the surface.

### Linear Interpolation

Tells what value lies between two points at a specific distance between them.\
$LERP(A, B, t) = A + t(B - A)$

or for vectors

$\vec{v}(t)=\vec{A} + t(\vec{B} - \vec{A})$

Can help in

-   Moving object from one position to another.\
-   For velocity dampning due to say friction or air resistence(How far the object will moving with velocity $\vec{V}$ will reach 0 when lerped with time.)
    -   $LERP(\vec{V}, 0, t)$

There are other functions which also interpolate in different ways. Say the interpolation is fast at start then slows down at middle.

### Epsilon Comparision

Floating point numbers can seem similar but can be off by a very small number.\
Epsilon comparision takes small tolerance value. Which basically tells are these numbers close enough.

$|a - b| < \epsilon$

Ex:
Say $\epsilon$ is 0.000001 and we have two numbers $a = 3.00000000$ and $b = 3.000000001$.

Then a == b will give false.\
But taking epsion comparision.

$|a - b| < \epsilon$. If nearly equal True else false.

Can be used for

-   Checking if velocity is zero. Say is velocity is nearly zero.
-   If two surface are in contact. Otherwise the surfaces will never reach each other.
