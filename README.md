# words

```go
import "github.com/typerat/words"
```

Package words generates word combinations consisting of an adjective describing
an animal. It is useful for creating human-readable strings to identify versions
of a program (as in the Ubuntu naming scheme).

## Usage

#### func  Get

```go
func Get(input interface{}) string
```
Get returns pseudorandom words for a given input.

#### func  Random

```go
func Random() string
```
Random returns almost true random words.

#### func  RandomAdjective

```go
func RandomAdjective() string
```
RandomAdjective returns a random adjective. Use the adjective to describe any
noun devoid of a description.
