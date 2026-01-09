# Generics in go

## Introduction 
- First seen in go v1.18 generics create balance between type safety and code reusability by parametrizing types
- Syntax:
```go
func FunctionName[T Constraint](param T) T {
    // ... function body
}
```

### [Code example](./01_intro.go)

## The problem
- before generics we had to write nearly identical logic for different types
- second option was `interface{}` which can be unsafe and requires type asserting 

### [Code example](./02_problem.go)

## Constraints
- `any` allows any type, if we use it in our Max() function we might get error because `>` operator is not defined in our type `T`. This case wont compile. We should use `ordered` constraint

### [Code example](./03_constraint.go)

## Reusable data structures
- before generics, creating a reusable data structure like Stack, Queue, List or Tree meant using `interface{}` and as we already know it requires type assertion and can lead to panic

### [Stack using generics](./03_constraint.go)

## Flexibility
- allows to operate on different types value while preserving safety. Useful for utility and transformation functions

### [Transform and filter functions](./05_flexibility.go)

## Under the hood
- go implements generics using **monomorphization** - type-specific versions of generic functions are being generated.

```go
// source code
func Identity[T any](value T) T {
    return value
}

func main() {
    a := Identity(42)       // int version
    b := Identity("hello")  // string version
}
```

```go
// generated
func Identity_int(value int) int {
    return value
}

func Identity_string(value string) string {
    return value
}
```

- this technique optimizes runtime performance but costs code size
 