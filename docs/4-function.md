# Function

## Constructing function
We can construct an unnamed function using the `{arguments -> body}` notation.

However, note that the function arguments must be of record type, but the return type can be any type.

For example:
```hs
add := {
  (value: Int, by: Int) -> 0
}
```

## Function type
Each function has a type, for example, the `add` function above has the type of:
```
(value: Int, by: Int) -> Int
```

Because of that, we can also define the `add` function as such:
```hs
add 
  : (value: Int, by: Int) -> Int
  = {
    (value, by) -> 0
  }
```

## Using function
In Topoi, we can use (a.k.a. apply) a function using the following notation (note that this is the only allowed notation):
```
arg1.func(arg2 := value1, arg3 := value2, ...)
```
Note that the first argument can ignore the property name, and the order is not important.

For example, the following function usage are equivalent:
```
1.add(by := 2)
2.add(value := 1)
(value:=1, by:=2).add
```

We can also chain function call:
```
1.add(by:= 2).add(by:= 3) 
```

Because of the dot notation, we can also directly apply an unnamed function (a.k.a. lambda) to an expression.

```
(x:=1, y:=2).{
  (x, y) -> x.add(by:=y)
}
```

## Pattern Matching (not finalized)
We can pattern match a value using the `equal` operator.

For example:
```hs
divide : (value: Int, by: Int) -> (result: Result.where(ok:= Int, fail:= String)) := {
  (by = 0) -> Result.fail("Cannot divide by 0"),
  (result, by) -> (
    result := -- something
  )
}
```

Another example:
```hs
type Shape := 
  | Circle (radius: Float)
  | Rectangle (height: Float, width: Float)
  | None

area: (shape: Shape) -> (result: Float) := {
  (shape = Circle(radius)) -> 
    radius.square.multiple(by:= pi),

  (shape = Rectangle(height, width)) ->
    height.multiple(by:= width),

  (shape = None) ->
    0
}
```
