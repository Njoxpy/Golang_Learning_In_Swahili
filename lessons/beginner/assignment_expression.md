- [Assignment Expression](#assignment-expression)
- [Implicit](#implicit)
- [Explicit](#explicit)
- [Printing to Console](#printing-to-console)

## Assignment Expression

## Implicit

- Implicit katika Programming ni kama neno automatic kwamna sio wewe unayefanya hicho kitendo ila kinafanywa na compiler au interpreter ila kwenye upnade wa Golang,implicit ni hivyo hivyo kwamba implicit inafanywa Interpreter yenyewe kwamba unafanya declaration ya variable yako kwa kutumia shorthand form kisha unaacha majukumu kwa interpreter yenyewe iweze kutambua kwamba hii ni data type gani ila interpreter inaangalia je value ambayo ipo ni ipi kama ni number basi data type ni number na kama ipo ndani ya " " quotes basi hiyo ni string na kuendelea.
  Mfano:

```go
	miaka := 20;
```

Katika mfano hapo juu temefanya miaka variable kufanywa implicit type conversion.Implicit ni kama tunaambia Go iweze kuguess(kukadiria) kuwa je data type yetu ni ipi!

**_Implicit Behaviour_**

- Go inatabia ya implicit behaviour pia kwenye programming,mfano zero value kwa ajili ya variables ambazo hazitakuwa initialized.

```go
var z int // zero value of int is 0 implicitly assigned to z

```

- Namna nyingine ya kufanya implicit kwenye Go ni pale ambapo unafanya declaration ya variable yako ila bila kuipa data type ya variable yako kwenye program yako kwa Mfano:

```go
	var age = 20;
```

Kwenye mfano wetu hapo Go ndio itafanya guessing ya kuweza kutambua kwamba variable type ya program yetu ni ipi.

## Explicit

- Kinyume cha implicit ni explicit, kwenye upande wa explicit kwenye kufanya declaration na initialization ya variable tutazingatia jina la variable data type yake pia value na var kwa ajili ya kufanya intialization, kutoka kwenye mfano ambao tumetumia pale juu kwenye upande wa explicit itakuwa hivi.

```go
	var miaka int = 20;
```

- Unapofanya declaration kwa kutumia explcit kwenye upande wa Go ina maana kwamba ili uweze kufanya decalaration ya inabidi tumie `var`, `const`,`type` `func` na `package` keyword ili uweze kufanya declaration ya variables zako mfano

```go
var x int = 10
const PI = 3.14
type Person struct {
    Name string
    Age  int
}
func Add(a, b int) int {
    return a + b
}

```

Kwamba tumefanya declaration na initialization ya variable yetu ambayo ni miaka.

## Printing to Console

```go
package main

import "fmt"

func main() {
    // Print a string
    fmt.Println("Hello, World!")

    // Print formatted output
    name := "John"
    age := 30
    fmt.Printf("Name: %s, Age: %d\n", name, age)

    // Print without newline
    fmt.Print("This ")
    fmt.Print("is ")
    fmt.Print("on ")
    fmt.Print("the same line.\n")

    // Print using Sprintf
    message := fmt.Sprintf("Hello, %s!", name)
    fmt.Println(message)
}

```

`fmt.Println()` inatumika kuweza kuprint lini kwenda kwenye console ikiwa na newline mwishoni, kama ambavyo kwenye C++ tukitumia cout haitatengeneza newline ila hadi uweke \n ila kwenye Go ukitumia `fmt.Println()` new line inatengenezwa muda wowote.

`fmt.Printf()` inatumika kuweza kuformat ni sawa na `printf` kwenye Lugha ya C,hivyo kama kuna variable yako inaitwa miaka inaweza kuwa formated, na bulliani value inaweza pia kuwa formatted kwa kutumia %T

```go
    name := "John"
    age := 30
    fmt.Printf("Name: %s, Age: %d\n", name, age)
```

Ni muhimu kuzingatia formatting kwenye program yako kama umetumia `Printf()` kinyume na hapo utapata hii warning.

```go
fmt.Printf call has arguments but no formatting directives
```

Pia ni muhimu kujua format specifier ambazo zinatumika kwenye program yako kwa kila data type kama ni integer tumia, angalia hapo chini

1. %s: Kwaajili ya kufanya formating ya string.
2. %d: Kwaajili formatting ya decimal integer
3. %f: Floating point numbe
4. %t : Boolean values(true or false)
5. %v : Format into it's default format
6. %T: Format data ya value yako
7. %p: Format pointer address

Kabla ya kufanya formatting ni muhimu kujua je data type yako ni ipi na format specifier yake ni ipi.

`fmt.Print()` inatumika kuprint bila kuwa na newline mwishoni.Mfano wa program hapo Chini

```go
	fmt.Print("Welcome")
	fmt.Print("Hello, Go is awesome!")

    // output: WelcomeHello, Go is awesome!

```

Kama ambavyo umeona kwenye program yako newline haijawa created.

`fmt.Sprintf()` inatumika kufanya format ya string bila kuweza kuprint immediately na matokeo yake yanakuwa stored kwenye variable `message`.
