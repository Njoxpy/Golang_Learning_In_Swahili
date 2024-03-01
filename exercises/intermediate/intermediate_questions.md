# Intermediate Questions

- [Channels](#channels)
- [Concurrency](#concurrency)
- [File IO](#file-io)
- [Interfaces](#interfaces)
- [Pointers](#pointers)
- [Testing](#testing)

## Pointers

1. How do you get the memory address of a variable?

2. How do you assign a value to a pointer?

3. How do you create a new pointer?

4. What is the value of x after running this program:

```go
func square(x *float64) {
 *x = *x * *x
}
func main() {
 x := 1.5
 square(&x)
}
```

5. Write a program that can swap two integers

```go
(x := 1; y := 2; swap(&x, &y) should give you
x=2 and y=1
```

## Channels

## Concurrency

## Interfaces

## File IO

## Testing

1. How do you get the memory address of a variable?
2. How do you assign a value to a pointer?
3. How do you create a new pointer?
4. What is the value of x after running this program:

    ```go
    func square(x *float64) {
    *x = *x* *x
    }

    func main() {
    x := 1.5
    square(&x)
    }
    ```

5.Write a program that can swap two integers
`(x := 1; y := 2; swap(&x, &y)` should give you `x=2 and y=1)`
