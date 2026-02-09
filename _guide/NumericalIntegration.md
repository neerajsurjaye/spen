# Numerical Integration

It basically says based on what I know right now. What will happen tiny bit later.

In math

- Integration = Adding up infintely small changes.

In computers

- We approximate that by adding up small finite chunks.

We Do

- position += velocity \* $\Delta t$
- The repeated += over time is numerical integration.

Integration = Adding up small chunks to get a total.

## Explicit Euler(Don't use it)

- In explicit euler basically position updated before updating velocity. And same with velocity its updated before acceleration.
- The system gains energy in this.

## SemiImplicit Euler

- Basically update velocity first then use that velocity to updte the position.
- The update order is only changed.
- Energy is slightly lost in this.
