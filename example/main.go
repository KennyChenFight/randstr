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
