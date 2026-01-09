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

- check