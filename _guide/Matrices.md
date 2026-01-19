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
