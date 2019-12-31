## Function construction
The type of every function in adhere to the following constraint:
```
Record -> Record
```
In other words, every function can only receive argument with type of record and return value with type of record.

## Function application

All function application in Topoi must follow the following syntax rule:

### Input with 0 field
Say:
```
f : () -> R
```
then invocation of `f` will be:
```
. f
```
and:
```
(. f) : R
```
Note that this syntax is not associative, `..f` does not mean `.(.f)`, it is simply invalid.

### Input with 1 field
Say:
```
f: (k: T) -> R
```
and
```
e: T
```
then the invocation of `f` will be:
```
e . f
```
and
```
(e . f) : R
```
Note that this syntax is left-associative, meaning that:
```
e . f . g = (e . f) . g
```

### Input with more than 1 fields
Say:
```
f : (j: T, k: U, l: V) -> R
```
and:
```
a: T
b: U
c: V
```
then the invocation of `f` can be any of the following:
```
a.f(k:=b, l:=c)
b.f(j:=a, l:=c)
c.f(j:=a, k:=b)
```
and all of the expression above will have the type of:
```
R
```
This syntax is also left-associative, meaning that:
```
a . f b . g c = (a . f b) . g c
```

## Relation to similar syntax
By reading the specification of `Record` one may notice that the syntax of record property access and that of function application are the same. Therefore, this section aims to provide a way to compile both of this feature without ambiguity.

Firstly, the parser only knows how to parse function application, so everything that resembles `A.F` or `A.F E` will be treated as function application.

Suppose we have the following code:
```
a := (x:=1, y:=2).x
```
After parsing, we get the syntax tree (assume Haskell-like syntax):
```
Assignment {
  left: Variable {name: "a"},
  right: FunctionApplication {
    function: Variable {name: "x"},
    leftArgument: Record {
      pairs: [
        {key: "x", value: Number {value: 1}},
        {key: "y", value: Number {value: 2}},
      ]
    },
    rightArgument: None
  }
}
```
During typechecking the application of `x`, the typechecker should first check if the `leftArgument` has type of `Record`, if yes, then check if the `function` is `Variable`, and if the `name` of the `Variable` matches one of the `key` of the `leftArgument`, then the `FunctionApplication` will be transformed into `RecordPropertyAccess`, else the typechecker should look for `x` in the function table.

A Haskell-Javascript-like pseudocode to perform such checking is presented below:
```hs
case expression of
  FunctionApplication {
    function: Variable {name},
    leftArgument: Record {pairs},
    rightArgment: None
  } -> 
    if pairs.some(pair -> pair.key == name) then
      RecordPropertyAccess {
        record: leftArgument,
        property: name
      }
    else
      -- treat as normal function application 
```

## Generic function
Generic functions are functions which contain types that can only be inferred during invocation.

However, unlike languages like Java and C#, where type variables has special syntax (e.g. in angular brackets `<T, U>`), type variables in Topoi are just like term-level expression.

TBC.