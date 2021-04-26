# randstr
generate random string very faster.

The code is based on this awesome StackOverflow answer by [icza](https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go).

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
```
You can reference example folder.

## Benchmark
```bash
goos: windows
goarch: amd64
pkg: github.com/KennyChenFight/randstr
cpu: Intel(R) Core(TM) i7-9700 CPU @ 3.00GHz
BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabet
BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabet-8              27867314                42.19 ns/op            1 allocs/op
BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabetLowercase
BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabetLowercase-8     29626336                41.15 ns/op            1 allocs/op
BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabetUppercase
BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabetUppercase-8     29051328                42.44 ns/op            1 allocs/op
BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabetBase62
BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabetBase62-8        41708664                29.78 ns/op            1 allocs/op
BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabetBase64
BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabetBase64-8        14275873                86.11 ns/op            1 allocs/op
```