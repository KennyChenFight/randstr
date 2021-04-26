// Package randstr provides generating random string
// As a quick start for generator:
//	const randomStrLength = 6
//	generator := randstr.NewFastGenerator(randstr.CharSetEnglishAlphabet)
//  generator.GenerateRandomStr(randomStrLength)
package randstr

import (
	"math"
	"math/rand"
	"time"
	"unsafe"
)

type RandomStrGenerator interface {
	GenerateRandomStr(n int) string
}

type CharSetType string

const (
	CharSetEnglishAlphabet CharSetType  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharSetEnglishAlphabetLowercase CharSetType = "abcdefghijklmnopqrstuvwxyz"
	CharSetEnglishAlphabetUppercase CharSetType = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharSetBase62                  CharSetType = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CharSetBase64                  CharSetType = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/"
)

var randSrc = rand.NewSource(time.Now().UnixNano())

type FastGenerator struct {
	charSetType CharSetType
	letterIdxBits int
	letterIdxMask int64
	letterIdxMax int
}

func NewFastGenerator(setType CharSetType) *FastGenerator {
	letterIdxBits := int(math.Trunc(math.Log2(float64(len(setType)))+1))
	return &FastGenerator{
		charSetType:   setType,
		letterIdxBits: letterIdxBits,
		letterIdxMask: 1<<letterIdxBits - 1,
		letterIdxMax:  63 / letterIdxBits,
	}
}

func (g *FastGenerator) GenerateRandomStr(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, randSrc.Int63(), g.letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), g.letterIdxMax
		}
		if idx := int(cache & g.letterIdxMask); idx < len(g.charSetType) {
			b[i] = g.charSetType[idx]
			i--
		}
		cache >>= g.letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}