# Pointers

- [Introduction](#introduction)
- [The * And & Operators](#the--and--operators)
- [The New Function](#the-new-function)
- [Working With Nil](#working-with-nil)
- [Types With Internal Pointers](#types-with-internal-pointers)

## Introduction

- Pointer katika Go ni ile variable ambayo inachukua same location na variable nyingine katika programu yako.Kama unatoka kwenye background ya C++ au hata C, utakuwa unajua vizuri nini maana ya pointers na kwanini zinatumika!

- Je pointers zipoje na zinafanya kazi vipi? Ili kuweza kufanya declaration ya pointer variable katika program yako utatumia * (asterisk),hapo utakuwa umefanya declaration ya pointer variable. na kuweka dereference tutatumia &.Angalia mfano kwanza hapo chini!, kwanjia ya kawaida declaration ya variable katika golang inakuwa hivi

```go
 var a int = 20
```

Tunaweza kutengeneza pointer variable katika ambayo itakuwa inapoint kwenye variable yako a, Mfano

```go
 var a int = 20
 var b *int = &a
```

Tumetengeneza pointer variable b ambayo inapoint kwenye memory location ambayo ni sawa na ya variable a.Ili kupata value ya variable zako utatumia referencing operator ambayo ni * na kupata memory location ya variable yako utatumia ampresand operator &,angalia mfano hapo chini!

```go
 var a int = 20
 var b *int = &a
 fmt.Println(*b, b) // Output: 20 0xc00000a0a8
```

- Kutoka kwenye mfano hapo juu,kupata value ya pointer variable b tutatumia * referencing operator na kupata memory location ya b itakuwa plain tu.Kwenye upande wa memory location usiwe na wasiwasi napo kwa sababau value memory location inaweza ikawa inatofautiana kutokana na computer.

- Derefrencing inatumika kwa ajili ya pointer variable tu na si vingine.Kwa kutumia dereference operator unaweza kubadili value ya variable integer type kwa kutumia referencing operator.Mfano:

```go
  *b = 30
 fmt.Println(*b) // 30
```

- Na ukijaribu kupata value ya variable yako jibu litakuwa 30 kwa kutumia referencing operator.

```go
package main

func main() {
 var a int = 20
 var b *int = &a
}
```

## The * And & Operators

- Katika Go,pointer inawakilishwa kwa kutumia `* (asterisk)` ikifuatiwa na aina ya value ambaye imehifadhiwa.

- `*` Inatumika kufanya `dereference` ya pointer variable.Dereferencing pointer inatupa access ya value ambapo pointer imepoint.

## The New Function

## Working With Nil

## Types With Internal Pointers
