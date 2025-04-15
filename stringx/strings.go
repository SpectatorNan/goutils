package stringx

import "strings"

func Split(s, sep string) []string {
	if s == "" {
		return make([]string, 0)
	}
	return strings.Split(s, sep)
}

func SplitIgnoreEmpty(s, sep string) []string {
	if s == "" {
		return make([]string, 0)
	}
	parts := strings.Split(s, sep)
	var result []string
	for _, part := range parts {
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}

func GetAlphabeticSequenceByNumber(n int) string {
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

func NextAlphabeticSequence(s string) string {
	if s == "" {
		return "A"
	}

	n := len(s)
	for i := n - 1; i >= 0; i-- {
		if s[i] < 'Z' {
			return s[:i] + string(s[i]+1) + strings.Repeat("A", n-i-1)
		}
	}
	return strings.Repeat("A", n+1)
}
