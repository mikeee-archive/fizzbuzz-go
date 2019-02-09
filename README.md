# fizzbuzz-go

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)[![Maintainability](https://api.codeclimate.com/v1/badges/4151f350ea03f3b30824/maintainability)](https://codeclimate.com/github/mikeee/fizzbuzz-go/maintainability)[![Test Coverage](https://api.codeclimate.com/v1/badges/4151f350ea03f3b30824/test_coverage)](https://codeclimate.com/github/mikeee/fizzbuzz-go/test_coverage)

buzz... fizz?

## So erm... what is this 🤔

Fizz buzz is a simple game that replaces multiples of 3 and 5 with Fizz or Buzz, if the number is a multiple of both then it's replaced by FizzBuzz! This code can be reused to check a number using the fizz buzz logic. There are test harnesses in place to help ensure the code works and is valid.

## Usage

Assuming that you have Golang installed, you can import the package using the following command.

```bash
go get github.com/mikeee/fizzbuzz-go
```

You can reuse the main function if you'd like. This relies on the 'fmt' package which should come with your distribution of Go. Alternatively it's possible to use the `strconv` library if that takes your fancy.

```go
import "fmt"
// Check takes an integer and applies the fizzbuzz logic to it
func Check(value int) (result string) {
    if value%3 == 0 {
        result = "Fizz"
    }
    if value%5 == 0 {
        result += "Buzz"
    }
    if result == "" {
        result = fmt.Sprintf("%d", value)
    }
    return result
}
```

## Example

```go
package main

import (
    "fmt"
    "github.com/mikeee/fizzbuzz-go"
)

func main() {
    for i := 1; i < 100; i++ {
        fmt.Println(fizzbuzz.Check(i))
    }
}
```

## Licence

Released under the MIT License - For full details check out [LICENSE](https://github.com/mikeee/fizzbuzz-go/blob/master/LICENSE)