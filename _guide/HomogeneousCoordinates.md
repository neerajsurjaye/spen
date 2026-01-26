## Homogeneous Coordinates

It is used beacause we want all transfomrations to happen in a uniform way.\
Translation cannot happen on a 2d space with a 2d matrix.

Thus we embed the 2d point in a 3d space and use shear transformation to move the points.
[Explanation](https://www.youtube.com/watch?v=AheaTd_l5Is)

Say we have a point $A = (x, y)$

We want to translate it by $(t_x, t_y)$.

We would first embedd the point to 3d space $A = (x, y, 1)$.

For translation the following matrix would be used.\
$ T = \begin{bmatrix}
1 & 0 & t_x \\
0 & 1 & t_y \\
0 & 0 & 1
\end{bmatrix}$

Then we would perform $A \cdot{} T$

$ T \cdot{} A = \
\begin{bmatrix}
1 & 0 & t_x \\
0 & 1 & t_y \\
0 & 0 & 1
\end{bmatrix} \
\begin{bmatrix}
x \\
y \\
1
\end{bmatrix} =
\begin{bmatrix}
x + t_x \\
y + t_y \\
1
\end{bmatrix} $

To go back to 2d space. We do ((x + tx) / w, (y + ty) / w). Where w is the third coordinate in our case itis one.

So the output would be.

$A' = \begin{bmatrix}
x + t_x \\
y + t_y
\end{bmatrix} $

**This should be only done for points not direction...**
