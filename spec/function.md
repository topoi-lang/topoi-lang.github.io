## Function construction
The type of every function in adhere to the following constraint:
```
Record -> Record
```
In other words, every function can only receive argument with type of record and return value with type of record.

## Function application

All function application in Topoi must follow the following syntax rule:

```
A . F
```
or
```
A . F E
```
Note that function application is left-associative (like number addition):
```
A . F E . F E = (A . F E) . F E

similar to

1 + 2 + 3 = (1 + 2) + 3
```

Where 
```
A: (any type)
F: Function
E: Record
```

Suppose `F`'s input type is `Record` with 1 field:
```
F: (K1: T1) -> R
```
And:
```
A: T1
```
Then:
```
(A . F) : R
```



Suppose `F`'s input type is `Record` with more than 1 field:
```
F: (K1: T1, K2: T2, K3: T3) -> R
```
and
```
A: T1
```
then
```
E: (K2: T2, K3: T3)
```
and
```
(A . F E) : R
```
Similarly, if 
```
A: T2
```
then
```
E: (K1: T1, K3: T3)
```
and
```
(A . F E) : R
```
And so on and so forth.


Suppose `F`'s input type is `Record` with 0 field:
```
F: () -> R
```
Then due to the semantical rule aformentioned, there's no way to invoke this function. However, if one wants to create a function that does not require an input, one can use the following idiom:
```
F: (_: ()) -> R
().F
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

