# Beginner Questions

- [Arrays Slices](#arrays-slices)
- [Assignment Expression](#assignment-expression)
- [Control Flow](#control-flow)
- [Error Handling](#error-handling)
- [Functions](#functions)
- [Introduction](#introduction)
- [Maps](#maps)
- [Packages](#packages)
- [Structs](#structs)
- [Variable](#variable)

## Arrays Slices

1. How do you access the 4th element of an array or
slice?
2. What is the length of a slice created using:
`make([]int, 3, 9)`?

3. Given the array: `x := [6]string{"a","b","c","d","e","f"}` what would `x[2:5]` give you?
4. Write a program that finds the smallest number
in this list:

```go
x := []int{
 48,96,86,68,
 57,82,63,70,
 37,34,83,27,
 19,97, 9,17,
}
```

## Assignment Expression

## Control Flow

1. What does the following program print:

```go
i := 10
if i > 10 {
 fmt.Println("Big")
} else {
 fmt.Println("Small")
}
```

2. Write a program that prints out all the numbers
evenly divisible by 3 between 1 and 100. (3, 6, 9,
etc.)

3. Write a program that prints the numbers from 1
to 100. But for multiples of three print "Fizz" instead of the number and for the multiples of five
print "Buzz". For numbers which are multiples
of both three and five print "FizzBuzz".

4. When to use `for loop` in Go,When to use `if statement` over `switch statement` in Go.

## Error Handling

## Functions

1. Sum is a function which takes a slice of numbers
and adds them together. What would its function signature look like in Go?

2. Write a function which takes an integer and
halves it and returns true if it was even or false
if it was odd. For example half(1) should return
(0, false) and half(2) should return (1,
true).

3. Write a function with one variadic parameter
that finds the greatest number in a list of numbers.

4. Using makeEvenGenerator as an example, write a
makeOddGenerator function that generates odd
numbers.
5. The Fibonacci sequence is defined as: `fib(0) =
0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).` Write a recursive function which can find fib(n).

6. What are defer, panic and recover? How do you
recover from a run-time panic?

## Introduction

## Maps

## Packages

1. Why do we use packages?

2. What is the difference between an identifier that starts with a capital letter and one which doesn’t? (Average vs average)

3. What is a package alias? How do you make one?

5. How would you document the functions you created in #3?

## Structs

1. What's the difference between a method and a
function?

2. Why would you use an embedded anonymous
field instead of a normal named field?

3. Add a new method to the Shape interface called
perimeter which calculates the perimeter of a
shape. Implement the method for Circle and
Rectangle.

## Variable

1. What are two ways to create a new variable?
2. What is the value of x after running:
`x := 5; x += 1?`
3. What is scope and how do you determine the scope of a variable in Go?
4. What is the difference between var and const?
5. Using the example program as a starting point,
write a program that converts from Fahrenheit into Celsius. `(C = (F - 32) * 5/9)`
6. Write another program that converts from feet
`into meters. (1 ft = 0.3048 m)`
