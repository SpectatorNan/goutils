package tools

import (
	"math/rand"
	"time"
)

/*
const (
	KC_RAND_KIND_NUM    = 0 // 纯数字
	KC_RAND_KIND_LOWER  = 1 // 小写字母
	KC_RAND_KIND_UPPER  = 2 // 大写字母
	KC_RAND_KIND_ALL    = 3 // 数字、大小写字母
	KC_RAND_KIND_SYMBOL = 4 // 数字、大小写字母、常规符号 !"#$%&'()*+,-./
)

// 随机字符串
func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}, {15, 33}}, make([]byte, size)
	is_all := kind > 3 || kind < 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = r.Intn(4)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + r.Intn(scope))
	}
	return string(result)
}
*/

const (
	KC_RAND_KIND_NUM                = 0  // 纯数字
	KC_RAND_KIND_LOWER              = 1  // 小写字母
	KC_RAND_KIND_UPPER              = 2  // 大写字母
	KC_RAND_KIND_SYMBOL             = 3  // 符号
	KC_RAND_KIND_NUM_LOWER          = 4  // 数字+小写字母
	KC_RAND_KIND_NUM_UPPER          = 5  // 数字+大写字母
	KC_RAND_KIND_NUM_SYMBOL         = 6  // 数字+符号
	KC_RAND_KIND_LOWER_UPPER        = 7  // 小写+大写字母
	KC_RAND_KIND_LOWER_SYMBOL       = 8  // 小写字母+符号
	KC_RAND_KIND_UPPER_SYMBOL       = 9  // 大写字母+符号
	KC_RAND_KIND_NUM_LOWER_UPPER    = 10 // 数字+小写+大写字母
	KC_RAND_KIND_NUM_LOWER_SYMBOL   = 11 // 数字+小写字母+符号
	KC_RAND_KIND_NUM_UPPER_SYMBOL   = 12 // 数字+大写字母+符号
	KC_RAND_KIND_LOWER_UPPER_SYMBOL = 13 // 小写+大写字母+符号
	KC_RAND_KIND_ALL                = 14 // 数字、大小写字母、符号
)

// 随机字符串
func Krand(size int, kind int) string {
	kinds := map[int][][]int{
		KC_RAND_KIND_NUM:                {{10, 48}},
		KC_RAND_KIND_LOWER:              {{26, 97}},
		KC_RAND_KIND_UPPER:              {{26, 65}},
		KC_RAND_KIND_SYMBOL:             {{15, 33}},
		KC_RAND_KIND_NUM_LOWER:          {{10, 48}, {26, 97}},
		KC_RAND_KIND_NUM_UPPER:          {{10, 48}, {26, 65}},
		KC_RAND_KIND_NUM_SYMBOL:         {{10, 48}, {15, 33}},
		KC_RAND_KIND_LOWER_UPPER:        {{26, 97}, {26, 65}},
		KC_RAND_KIND_LOWER_SYMBOL:       {{26, 97}, {15, 33}},
		KC_RAND_KIND_UPPER_SYMBOL:       {{26, 65}, {15, 33}},
		KC_RAND_KIND_NUM_LOWER_UPPER:    {{10, 48}, {26, 97}, {26, 65}},
		KC_RAND_KIND_NUM_LOWER_SYMBOL:   {{10, 48}, {26, 97}, {15, 33}},
		KC_RAND_KIND_NUM_UPPER_SYMBOL:   {{10, 48}, {26, 65}, {15, 33}},
		KC_RAND_KIND_LOWER_UPPER_SYMBOL: {{26, 97}, {26, 65}, {15, 33}},
		KC_RAND_KIND_ALL:                {{10, 48}, {26, 97}, {26, 65}, {15, 33}},
	}

	result := make([]byte, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		kindSet := kinds[kind]
		ikind := r.Intn(len(kindSet))
		scope, base := kindSet[ikind][0], kindSet[ikind][1]
		result[i] = uint8(base + r.Intn(scope))
	}
	return string(result)
}
