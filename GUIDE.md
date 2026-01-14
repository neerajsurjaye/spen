# Guide

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
