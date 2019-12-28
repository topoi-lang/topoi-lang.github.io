# Expressions

## Record 
Record expression are key-value pairs.
Each record are enclosed by round brackets.

### Record Construction
Empty record is denoted as `()`.

Examples of non-empty records:

```go
(name := "topoi", age := "12")
```

Each property of a record can so be optionally typed, to ensure that the property will be initiated with the correct type.

```go
(name : String = "topoi", age := "0")
```

We can also construct nested record:
```go
(
  name := "topoi", 
  friend := (
    name := "haskell", 
    age := "30"
  )
)
```

### Record Type

The type of each record can be inferred by the Topoi type system, for example  `(name := "topoi", age := "12")`
has the type of `(name: String, age: String)`

### Record destructuring
The property of each record can be destructured, and be populated in the current environment.

```hs
(name, age) := (name := "topoi", age := "0")
-- `name` and `age` will be exposed to the current environment as named variables
```
Note that not all of the properties need to be destructured.
```hs
(name) := (name := "topoi", age := "0")
```
Also, the destructuring order does not matter:
```hs
(age, name) := (name := "topoi", age := "0")
```

Besides that, we can also re-alias the property name using they keyword `as`.

```hs
(name as myName, age) = (name := "topoi", age := 1)
-- `myName` and `age` is exposed, but `name` is not
```

## Accessing properties
We can do it by using the dot notation.

```
record.propertyName
```

For example:
```hs
me := (name := "john", age := 2)

myName = me.name
```

## Mutating properties
In Topoi, properties mutation is not allowed, however we can use the spread notation with the triple dot operator.

For example,
```hs
me := (name := "john", age := 2)

newMe := (...me, age := 3)
```
Basically it's the same as JavaScript spread operation.

We can also combine the properties of multiple records into one as such:
```sh
me := (name := "john", age := 2)
otherDetail := （address := "sea")

fullMe = (...me, ...otherDetail)
```

