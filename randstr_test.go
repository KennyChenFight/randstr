package randstr

import "testing"

const n = 6

func BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabet(b *testing.B) {
	generator := NewFastGenerator(CharSetEnglishAlphabet)
	for i := 0; i < b.N; i++ {
		generator.GenerateRandomStr(n)
	}
}

func BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabetLowercase(b *testing.B) {
	generator := NewFastGenerator(CharSetEnglishAlphabetLowercase)
	for i := 0; i < b.N; i++ {
		generator.GenerateRandomStr(n)
	}
}

func BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabetUppercase(b *testing.B) {
	generator := NewFastGenerator(CharSetEnglishAlphabetUppercase)
	for i := 0; i < b.N; i++ {
		generator.GenerateRandomStr(n)
	}
}

func BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabetBase62(b *testing.B) {
	generator := NewFastGenerator(CharSetBase62)
	for i := 0; i < b.N; i++ {
		generator.GenerateRandomStr(n)
	}
}

func BenchmarkFastGenerator_GenerateRandomStr_EnglishAlphabetBase64(b *testing.B) {
	generator := NewFastGenerator(CharSetBase64)
	for i := 0; i < b.N; i++ {
		generator.GenerateRandomStr(n)
	}
}