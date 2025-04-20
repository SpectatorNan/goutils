package tools

import (
	"fmt"
	"strings"
	"testing"
)

func TestKrand(t *testing.T) {
	type args struct {
		size int
		kind int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test_num", args: args{size: 10, kind: KC_RAND_KIND_NUM}},
		{name: "test_lower", args: args{size: 10, kind: KC_RAND_KIND_LOWER}},
		{name: "test_upper", args: args{size: 10, kind: KC_RAND_KIND_UPPER}},
		{name: "test_symbol", args: args{size: 10, kind: KC_RAND_KIND_SYMBOL}},
		{name: "test_num_lower", args: args{size: 10, kind: KC_RAND_KIND_NUM_LOWER}},
		{name: "test_num_upper", args: args{size: 10, kind: KC_RAND_KIND_NUM_UPPER}},
		{name: "test_num_symbol", args: args{size: 10, kind: KC_RAND_KIND_NUM_SYMBOL}},
		{name: "test_lower_upper", args: args{size: 10, kind: KC_RAND_KIND_LOWER_UPPER}},
		{name: "test_lower_symbol", args: args{size: 10, kind: KC_RAND_KIND_LOWER_SYMBOL}},
		{name: "test_upper_symbol", args: args{size: 10, kind: KC_RAND_KIND_UPPER_SYMBOL}},
		{name: "test_num_lower_upper", args: args{size: 10, kind: KC_RAND_KIND_NUM_LOWER_UPPER}},
		{name: "test_num_lower_symbol", args: args{size: 10, kind: KC_RAND_KIND_NUM_LOWER_SYMBOL}},
		{name: "test_num_upper_symbol", args: args{size: 10, kind: KC_RAND_KIND_NUM_UPPER_SYMBOL}},
		{name: "test_lower_upper_symbol", args: args{size: 10, kind: KC_RAND_KIND_LOWER_UPPER_SYMBOL}},
		{name: "test_all", args: args{size: 10, kind: KC_RAND_KIND_ALL}},
	}

	for _, tt := range tests {
		t.Logf(strings.Repeat("=", 50))
		t.Run(tt.name, func(t *testing.T) {
			got := Krand(tt.args.size, tt.args.kind)
			t.Logf("%s Krand(%v, %v) = %v", tt.name, tt.args.size, tt.args.kind, got)
		})
		t.Logf(strings.Repeat("=", 50))
		fmt.Println()
	}
}
