## Matrices

A rectangular grid of numbers

$ A = \begin{bmatrix}
1 & 2\\
3 & 4
\end{bmatrix}$

- Usually represents a Vector, Linear Transform(How vectors change) etc.
- A vector tranformation is also known linear transformation.
    - It can also be said matrix transform space. Like rotate, scale, flip sear etc.

### Matrix vs Vectors

- Vector represents some Position, direction, force, velocity etc.
- Matrix represents how Vectors change after some tranformation.

## Vectors as matrices

Vectors can be represented using matrices.

### 2D

$A = \begin{bmatrix}
X \\
Y
\end{bmatrix}$

### 3D

$A = \begin{bmatrix}
X \\
Y \\
Z 
\end{bmatrix}$

## Identity Matrix

Represents a matrix which when multiplied does nothing. It's similar to multiplying by 1.

Represented by $I$

$I = \begin{bmatrix}
1 & 0 & 0 \\
0 & 1 & 0\\
0 & 0 & 1
\end{bmatrix}$

- It basicallty mean space is not transformed after applying this matrix.
- This can be initial value for a matrix.

## Matrix Multiplication

Matrix multiplication mean composition of transformation.

$C = A \cdot{} B$

- This means apply transform B then apply transform A.

Condition for multiplication

- If $A = m \times n$ and $B = n \times o$
- Then multiplication is possible.
- Number of columns of A = number of rows of B for multiplication to be possible

## Transforming Vector using a transform matrix

Say we have a transform

- $ A = \begin{bmatrix}
  1 & 2\\
  3 & 4
  \end{bmatrix}$
- This states that after transform $\hat{i}$ will be at 1,3 and $\hat{j}$ will be at 2,4
- Ths first column is transform for $\hat{i}$ and the second is for $\hat{j}$

We have a vector and want to know its value after transform

- $\vec{V} = \begin{bmatrix}
  2 \\
  5
  \end{bmatrix}$
    - This can also be written as $\vec{V} = 2 \times \hat{i} + 5 \times \hat{j}$

- By multiplying we mean
    - $\vec{V} = 2 \times transformd(\hat{i}) + 5 \times transformd(\hat{j})$
    - $\vec{V} = 2 \times \begin{bmatrix}
                                    1 \\
                                    3
                            \end{bmatrix} 
                + 5 \times \begin{bmatrix}
                                    2 \\
                                    4
                            \end{bmatrix}$
    - $\vec{V} = \begin{bmatrix}
                        1 \\
                        3
                \end{bmatrix} 
                + \begin{bmatrix}
                                    4 \\
                                    8
                            \end{bmatrix}$

    - $\vec{V} = \begin{bmatrix}
      5 \\
      11
      \end{bmatrix}$

**Order of multiplication Matters**

- $ A \cdot{} B \ne B \cdot{} A$

## 2 x 2 Matrix multiplication

If We have two matrices

- $ X = \begin{bmatrix}
  A & B\\
  C & D
  \end{bmatrix}$

- $ Y = \begin{bmatrix}
  M & N\\
  O & P
  \end{bmatrix}$

Then Y transforming from X would be .

- $ X \cdot Y = \begin{bmatrix}
  A \times M + B \times O & A \times N + B \times P\\
  C \times M + D \times O & C \times N + D \times P
  \end{bmatrix}$

## 2D Matrix

- Represents how vectors transforms over a 2D space.
- Stores rules for how vectors change over 2D space.
    - This cannot do translation.

## Scaling Matrix

- Scaling matrix changes the size of vectors.
    - It can stretch, shrink, flip(negative scale)

- $S = \begin{bmatrix}
        S_x & 0\\
        0 & S_y
        \end{bmatrix}$
    - If $S_x == S_y$. Then it uniformely scales.
- Non uniform scaling physics can break it. Thus it is avoided.

## Rotation Matrix

- The below matrix rotates counter-clockwise.
- $S = \begin{bmatrix}
        Cos \theta & -Sin \theta\\
        Sin \theta & Cos \theta
        \end{bmatrix}$
- Clockwise rotation
    - $S = \begin{bmatrix}
        Cos \theta & Sin \theta\\
        -Sin \theta & Cos \theta
        \end{bmatrix}$
- Note: Always use radians
    - $radians = degress \times \frac{\pi}{180}$

### Undo rotation

- For rotation transpose = inverse of a matrix.
    - Inverse = Used to reverse a operation performed by a matrix.
    - Transpose = Swapping rows and columns
- This works because rotation matrix are orthogonal.
  j

## Applying Transformations (TRS)

Objects do not store the transformed state. But these transformation parameters.

- T : Transform as Vector
- R : Rotation as angle
- S : Scale as a vector

When any transformation update comes it is applied on these values.\
When we want objects transformed state we use these values to create matrices and combine them.

It should be done in this order only $T \cdot{} R \cdot{} S$. Where Scale is applied first then rotation then translation.

**Note: This applies to column vectors**

## Inverse

A inverse matrices undoes the transformatoin done by a matrix.

$A \cdot{} A^{-1} = A^{-1} \cdot{} A = I$

In graphics pipeline just inverse TRS instead of full matrix inversion.

**Note: Inverse doesn't exist if a matrix transforms a vectors to lower dimesnion. Like a 3d space to a plane, line or point.**

- Inverse translation: If translation is $t_x , t_y$ then inverse is $-t_x , -t_y$
- Inverse rotation: If rotation is $\theta{}$ then inverse is $-\theta{}$.
- Inverse Scale: If scale is $s_x, s_y$ then inverse is $\frac{1}{s_x} , \frac{1}{s_y}$
- Inverse TRS: $M^{-1} = S^{-1} \cdot{} R^{-1} \cdot{} T^{-1}$
    - **In TRS. First scale then rotate then transform. So in inverse first reverse transform then reverse rotate and then reverse scale.**
