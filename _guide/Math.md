## Math

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

- Checking if velocity is zero. Say is velocity is nearly zero.
- If two surface are in contact. Otherwise the surfaces will never reach each other.
