# Variables and Types

- [Introduction](#introduction)
- [Variable Declaration and Initialization](#variable-declaration-and-initialization)
- [Visibility](#visibility)
- [Scope](#scope)
- [Naming Conventions](#naming-conventions)
- [Data Types](#data-types)
- [Type Inference](#type-inference)
- [Type Conversion](#type-conversion)
- [Mistakes](#mistakes)
- [Best Practices](#best-practices)

## Introduction

Katika Go, variables zinatumika kwa ajili ya kuhifadhi value za data.Kutokana na uwezo wa kuhifadhi hizo data zinaweza zikatumika kwa ajili ya kufanya operesheni mbalimbali.Kuweza kujua variables na aina zake ni muhimu ili kuweza kuandika programms za go ambazo ziko effective.Mfano wa namna tunaweza tukatumia variable kwenye upande wa mpira!

> Kwenye upande wa mpira, huwa tunakusanya taarifa za  vitu vifuatavyo matokeo(score), mda uliobaki(time left). Kwa hiyo fikiria variable kama contena ambalo unaweza ukahifadhi taarifa(information) za timu ili ziweze kutumiwa baadae. Mfano: Unaweza ukawa na variable inaitwa `score`, ambayo inahifadhi matokeo ya mchezo. Variable nyingine ni `timeLeft` yaani muda uliobaki.

Hivyo ndio variable itatumika katika ulimwengu wa programming pia,kwa ajili ya kuhifadhi taarifa(information).

Golang ni `statically typed langauge` kwamba variable ikishakuwa [declared](/lessons/beginner/variables_and_types.md#variable-declaration-and-initialization) katika Go haiwezi kubadilishwa kuwa ya type nyingine, Mfano ukitengeneza variable ya type number katika Go haiwezi kubadilishwa kuwa variable ya type string katika go ila ungekuwa kwa upande wa `dynamic typed langauge` kama Python unaweza ukabadilisha.

Mfano Katika Golang: Kwamba variable ya data moja haiwezi badilishwa kuwa ya data nyingine.

```go
package main

import "fmt"

func main() {
	var miaka int = 20
	var miaka string = "Ishirini"

	fmt.Println(miaka)
//     # command-line-arguments
// .\main.go:7:6: miaka redeclared in this block
//         .\main.go:6:6: other declaration of miaka

}

```

Mfano katika Python:

```py
miaka = 20
miaka = "Ishirini"
print(miaka) // "Ishirini
```


## Variable Declaration and Initialization

- Kufanya declaration ya variable katika Go utaanza na `var` keyword ikifuatiwa na jina la variable yako na kisha `type` aina ya data.

Mfano-1

```go
var jina string
```

- Variables zinakuwa declared kwa kutumia `var` keyword ikifuatiwa na jina la variable na aina(type). 

Mfano-2

```golang
var miaka int
```

Mfano hapo juu unafanya declaration ya variable inayoitwa `miaka` ya aina `integer`.

**Initialization** ni kitendo cha kuipa value variable yako.Mfano variable yetu ya miaka baada ya kuwa declared pale juu tumeipa na value(initialization) yake ambayo ni 30.

```golang
var miaka int = 30;
```

Katika Go,kuna shorthand syntax ya variable declaration na initialization

```golang
miaka := 30
```

Hio inafanya automatic infers ya type ya variable declaration yako kutokana na na value inayakokuwa assigned kwa mfano hapo value yetu ni 30 hivyo type ni int kwa hiyo compiler ili kuweza kujua kwamba data type yake ni ipi? inaangalia kwanza value inayohofadhiwa katika variable yako.Mifano zaidi ya shorthand syntax.

```golang
package main

import "fmt"

func main() {
	jina := "Neicore"
	kozi := "Computer Science"
	miaka := 20

	fmt.Println("Jina:", jina)
	fmt.Println("Miaka:", miaka)
	fmt.Println("Kozi:", kozi)
}

```

## Visibility

Visibility , ina determine kama variable, funnction constant, au type yoyote ile inawezaa kuwa accesed kutoka nje ya package ambayo imekuwa declared. Katika Go, visibility inakuwa controlled kutokana na naming conventions, majina ambayo yanaanza na herufi kubwa ya mwanzo yanaweza kuwa exported yaani kuwa accesed nje ya package ila kwa herufi ambazi zinaanza na herufi ndogo haziwezi kuwa accesed nje ya package.Mfano mzuri wa hili ni pale unatakapokuwa unajifunza jinsi ya kutengeneza `API` Kwa kutumia Golang kuna baadhi ya function hautaweza kupata acces yake kwenye `main.go` file kama jina la function limeanza na herufi ndogo.

Chini ni mfano unaoeleza variable declaration, initialization na matumizi.

```go
package main

import "fmt"

func main() {
    var miaka int
    miaka = 30

    var jina string = "Nei"

    ameolewa := false

    fmt.Println("jina:", jina)
    fmt.Println("miaka:", miaka)
    fmt.Println("Ameolewa:", ameolewa)
}

```

## Scope

Scope au naweza nikasema wigo ni eno la code ambapo variable, constant au function inaweza kuwa accesed. Ina determine wapi katika program yako unaweza ukarefer identifier fulani unaweza ukaita identifier yako. Katika go, variables zina block scope, maana kwamba unaweza ukapata access ya variable katika block ile tu ambapo variable yako imekuwa declared. Mfano:

```go
package main

import "fmt"

func main() {
    // Variable 'x' imekuwa declared ndani ya scope ya 'main' function
    var x int = 10
    fmt.Println(x) // 'x' inakuwa accessible hapa
}

```

Katika mfano hapo juu x imekuwa accesed ndani ya `main` function tu ile kama ungejaribu kupata accesed ya x nje ya main function haitawezekana kutokana na scope.

> Mfano mzuri wa variable ni ule wa FIFA na TFF, kwamba FIFA anauwezo wa kutoa onyo kwa TFF ila TFF hawezi kutoa onyo kwa FIFA, hivyo FIFA ni kama Global variable na TFF ni kama local Variable.

Katika Go, variables zinakuwa na block scoped level, maana kwamba variable zinakuwa accessed ndani ya block fulani ambazo zinakuwa declared ila kitendo cha kutaka kuaccess variable nje na block italeta error katika program yako pia hizo variable zinazokuwa ndani ya block zinaitwa **local variables**.

Variable zinazokuwa declared katika package level(nje ya main function) zinakuwa accesed throughout the package na ndio zinaitwa **global variable**.

## Naming Conventions

Katika Go,naming conventions ni muhimu ili uweze kuandika code ambazo zipo clean: Hapo chini ni mfano wa naming conventions ambazo zipo common

- **Tumia camelCase:** Go inatumia camelCase kwa ajili ya kufanya naming ya variables, functions, na package levels identifiers. Hii ina maana kwamba inaanza na herufi ndogo(lowercase letter) na kutumia mixed case kuweza kutofautisha maneno, mfano `myVariableName`, `calculateScore`, `packageName`.

## Data Types

Aina za data katika Go zipo za aina nne Tu Ambazo ni: integer, float, string na booleans.

 *1. Integers*: 
 Integers ziantumika kuhifadhi namba njima katika Go. `int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64`.Kwenye upande integers kuna signed integers na unsigned integers.

 Signed integers zinawakilisha integers ambazo ni hasi(negative) na chanya(positive), Mfano:

 ```go
 package main

import "fmt"

func main() {
    var x int = -42
    var y int8 = -8
    var z int16 = -16384
    var w int32 = -2147483648
    var v int64 = -9223372036854775808

    fmt.Println("int:", x)
    fmt.Println("int8:", y)
    fmt.Println("int16:", z)
    fmt.Println("int32:", w)
    fmt.Println("int64:", v)
}

 ```

Unsigned integer data types zinawakilisha integers ambazo ni chanya(positive) pamoja na sifuri.

Mfano:

```go
package main

import "fmt"

func main() {
    var x uint8 = 255
    var y uint16 = 65535
    var z uint32 = 4294967295
    var w uint64 = 18446744073709551615

    fmt.Println("uint8:", x)
    fmt.Println("uint16:", y)
    fmt.Println("uint32:", z)
    fmt.Println("uint64:", w)
}

```

 *2. Floating-point numbers*:
 Float ni namba ambazo zinawakilisha namba njima ila zikiwa na sehemu. Mfano wa floats ni `float, float32, float64`

 ```go
 package main

import "fmt"

func main() {
    // Declaration and initialization of float32
    var float32Num float32 = 3.14
    fmt.Println("float32:", float32Num)

    // Declaration and initialization of float64 (by default, floating point literals are considered float64)
    var float64Num float64 = 3.14
    fmt.Println("float64:", float64Num)

    // If you don't specify the type, Go will infer it based on the context
    inferredFloat := 3.14
    fmt.Printf("Inferred float type: %T\n", inferredFloat)
}

 ```
 *3. Strings*: `string` 
 Katika Go, mifuatano ni mifuatano ya baiti zinazowakilisha maandishi yanayoweza kusomeka na binadamu. Strings katika Go hazibadiliki, kumaanisha mara zinapoundwa, maudhui yake hayawezi kubadilishwa. Haya hapa ni maelezo ya mifuatano katika Go na mifano

 ```go
 package main

import "fmt"

func main() {
    // Declaration and initialization of a string
    var str1 string = "Hello, World!"
    fmt.Println(str1)

    // Short declaration syntax
    str2 := "This is another string."
    fmt.Println(str2)

    // Empty string declaration
    var emptyString string
    fmt.Println("Empty string:", emptyString)
}

 ```
 *4. Booleans*: `bool`

 Value za bulliani katika Go zinawakilisha truth values na zina value mbili tu ambazo zi kweli au sikweli(true au false)

 ```go
 package main

import "fmt"

func main() {
    // Declaration and initialization of boolean variables
    var b1 bool = true
    var b2 bool = false

    fmt.Println("b1:", b1)
    fmt.Println("b2:", b2)

    // Short declaration syntax
    b3 := true
    fmt.Println("b3:", b3)
}

 ```

## Type Inference

- Type inference inaipa Go uwezo wa kujua jina la variable automatiki kutokana na value inayokuwa assigned hio tunaita `impilicit`. Explicit type declaration ,hii inahusika na kuspecify aina ya variable yako wakati wa kufanya declaration ya variable.Mfano wa impilicit na explicit

```golang
package main

import "fmt"

func main() {
    
 // explicit
 var miaka int
 miaka = 30
 var jina string = "Nei"

 // implicit
 ameolewa := false

 fmt.Println("jina:", jina)
 fmt.Println("miaka:", miaka)
 fmt.Println("Ameoewa:", ameolewa)
}

```

## Type Conversion

- Type Conversion inakuwezesha kuweza kubadili value kutoka type moja hadi kuenda type nyingine.Mfano:

```golang
var x int = 10
var y float64 = float64(x)
```

Hapa tunfanya conversion ya integer x kwenda float64 na kufanya assigning kwa y.

## Mistakes

- Yafuatayo ni baadhi ya makosa ambayo watu wanafanya wakati wanajifunza Golang kwenye upande wa variables na data types.

1. Kufanya declaration ya variable katika Go na kutoitumia, Mfano:

```go
package main

import "fmt"

func main() {
	jina := "Neicore"
	kozi := "Computer Science"
	miaka := 20

	fmt.Println("Jina:", jina)
	fmt.Println("Miaka:", miaka)
    // output: # command-line-arguments
// .\main.go:7:2: kozi declared and not used

}

```

Hivyo ni muhimu kama endapo utatumia variable katika go hakikisha umeifanyia kazi na kama hutumii itoe kabisa au unaweza ukaweka kama comment.

2. Kuzipa variable majina ambayo hayaeleweki:
Mfano:

```golang
package main

import "fmt"

func main() {
	j := "Neicore"
	k := "Computer Science"
	x := 20

	fmt.Println("J:", j)
	fmt.Println("K:", k)
	fmt.Println("X:", x)
}
```

Hapo juu nimefanya declaration na intialization ya value zetu ila majina ya variable yana utata hiyo hakikisha unachagua meaningfull na descriptive name za variable zako.

## Best Practices