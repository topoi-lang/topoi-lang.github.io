# Record

## Invoking properties that has type of function
Suppose we have a variable `x` which have the record type of
```
(
  f: (k: T,) -> (),
)
```
Suppose:
```
e: T
```
Then, one might be tempted to invoke `f` using the following code:
```
x.f(k := e,)
```
However, due to semantics of function application, which states that the argument should appear before the function expression (in this case `f`), the code above is invalid.

The correct way to invoke `f` is:
```
e.(x.f)
```

## Record type spreading

Note that `Record` type can be spreaded as well.

Suppose
```
type alias A = (K1: T1,)
type alias B = (K2: T2, K3: T3)
type alias C = (...A, ...B)
```
Then
```
C : (K1: T1, K2: T2, K3: T3)
```