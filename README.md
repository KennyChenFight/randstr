# randstr
generate random string very faster.

The code is based on this awesome StackOverflow answer by [icza](https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go/22892986#22892986).

## How to use
### install
`go get github.com/KennyChenFight/randstr`
### Example
```go
package main

import (
	"fmt"
	"github.com/KennyChenFight/randstr"
)

func main() {
	const randomStrLength = 6
	generator := randstr.NewFastGenerator(randstr.CharSetEnglishAlphabet)
	fmt.Println(generator.GenerateRandomStr(randomStrLength))
}
You can reference example folder.
```