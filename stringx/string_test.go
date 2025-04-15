package stringx

import (
	"strings"
	"testing"
)

func TestNextAlphabeticSequence(t *testing.T) {
	str := "A"
	t.Logf("str 0: %v", str)
	for i := 1; i < 100; i++ {
		str = NextAlphabeticSequence(str)
		t.Logf("str %d: %v", i, str)
	}
}

func TestGetAlphabeticSequenceByNumber(t *testing.T) {

	for i := 0; i < 100; i++ {
		t.Logf("str %d: %v", i, GetAlphabeticSequenceByNumber(i))
	}
}

func BenchmarkNextAlphabeticSequence(b *testing.B) {
	str := "A"
	for i := 0; i < b.N; i++ {
		str = NextAlphabeticSequence(str)
	}
}

func BenchmarkGetAlphabeticSequenceByNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getAlphabeticSequenceByNumber(i)
	}
}

func BenchmarkGetAlphabeticSequenceByNumber1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getAlphabeticSequenceByNumber1(i)
	}
}

func BenchmarkGetAlphabeticSequenceByNumber2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getAlphabeticSequenceByNumber2(i)
	}
}

func getAlphabeticSequenceByNumber(n int) string {
	seq := strings.Builder{}
	for n >= 0 {
		seq.WriteByte(byte('A' + n%26))
		n = n/26 - 1
	}
	result := []rune(seq.String())
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

func getAlphabeticSequenceByNumber1(n int) string {
	seq := strings.Builder{}
	for n >= 0 {
		seq.WriteByte(byte('A' + n%26))
		n = n/26 - 1
	}

	// 反转字符串的过程中直接构建
	result := make([]byte, seq.Len())
	for i := 0; i < seq.Len(); i++ {
		result[i] = seq.String()[seq.Len()-1-i]
	}

	return string(result)
}

func getAlphabeticSequenceByNumber2(n int) string {
	source := make([]byte, 0)
	for n >= 0 {
		source = append(source, byte('A'+n%26))
		n = n/26 - 1
	}

	// 反转字符串的过程中直接构建
	result := make([]byte, len(source))
	for i := 0; i < len(source); i++ {
		result[i] = source[len(source)-1-i]
	}

	return string(result)
}
