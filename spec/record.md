# Record
## Construction
Record is a type of data structure that consists of 0 or more key-value pairs(KVP) .  
Record is enclosed within parenthesis.  
Each KVP must be terminated with comma.  
Each KVP can be statically typed.

Record with zero KVP is denoted as
`()`, and has the type of `()`.

Record with one or more KVP:
```
(x := 1,)
(y := 2, z := 3,)
```
Record with statically typed KVP.
```
(x: Int := 1,)
(y: Int := 2, z := 3)
```

Note that the value of each KVP can be any expression.

Also, each record cannot contain duplicated keys, which makes the following example invalid:
```
(x := 1, x := 1,)
```

## Destructuring
Destructuring is the process where the values of a record can be assigned to new variables in the current environment without using record property accessing.

Say we have a record `r`:
```
r : (x: T, y: U)
```
Then we can destructure the value of `x` and `y` into the current environment:
```
(x, y) := r
```
Doing this will expose two new variables to the current environment, which are `x: T` and `y: U`.

We can also re-alias the destructured value to other names that differs from their original key.

```
(x as j, y as k) := r
```
Doing this will expose `j: T` and `k: U` to the current environment.

## Shorthand
Suppose we have two variables:
```
x: T
y: U
```
To construct a new record out of `x` and `y`, we can write:
```
r := (x, y)
```
which is equivalent to:
```
r := (x := x, y := y)
```

## Spreading
We can combine the KVPs from multiple records into one record by using the spreading operator `...`.


Suppose we have two records:
```
r : (x: T, y: U)
s : (y: V, z: W)
```
and
```
a := (...r, ...s)
b := (...s, ...r)
```
then
```
a : (x: T, y: V, z: W)
b : (x: T, y: U, z: W)
```
This implies that the sequence of spreading is significant.

## Property access
Property access means to access the value of a given key (a.k.a property). This action can be achieved via the dot notation.

Suppose:
```
r: (x: T, y: (a: U))
```
Then:
```
r.x   : T
r.y   : (a: U)
r.y.a : U
```

Note that the dot operator is left associative, thus:
```
a . b . c = (a . b) . c
```


## Equality
Two record type, X and Y are consider equal if:
- they have the same number of KVPs
- the set of all keys of X is equal to the set of all keys of Y
- the type of value of each matching KVP are equal

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

## Row type polymorphism
Say:
```
a : (j: T, k: U, l: V)
f : (x: (j: T)) -> R
```
then:
```
a . f
```
is invalid, because the type of `x` does not require `(k: U, l: V)`. However to allow row type polymorphism, we just need to tweak `f` to use the record type spreading:
```
f : (infer type: Type, x: (j: T, ...Type)) -> R
```
