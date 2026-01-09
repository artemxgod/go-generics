# Generics in go

## Introduction 
- First seen in go v1.18 generics create balance between type safety and code reusability by parametrizing types

### [Example](./01_intro.go)

```go
func FunctionName[T Constraint](param T) T {
    // ... function body
}
```

- check