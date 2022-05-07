# go

## My golang scripts repo

### hello_world

The first script is called `hello_world` as is customary in the world of programming.

## Key observations and notes with go

`package main` is a special package that will make the code compile to an executable.

`fmt` is for format

then we define a function with `func`

then we call the Println from the `fmt` package

Most programs start with the boilerplate:

```go
// comments after double slash
package main

import (
	"fmt"
)

func main() {
	
}
```

### data types

Go is strong typed, we need to declare the type before variable assignment. The following types are common, others exist:

* int8/16/23/64 - int's of various sizes
* uint8/16/23/64 - unsigned int's
* float32/64
* string

### Program control

You can get it to predict the type using the assignment `:=`

There is only one type of loop - for loop

If and if else statements are possible, example in the fizzbuzz program.

### slices

The list/array data structure is a slice in golang.


### maps

the mapping/dictionary type is called map

### functions

define functions with `func`