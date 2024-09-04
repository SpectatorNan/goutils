package sensitiveFilter

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stringx"
	"testing"
)

func TestOne(t *testing.T) {
	filter := stringx.NewTrie([]string{
		"AV演员",
		"苍井空",
		"AV",
		"日本AV女优",
		"AV演员色情",
	})
	keywords := filter.FindKeywords("日本AV演员兼电视、电影演员。苍井空AV女优是xx出道, 日本AV女优们最精彩的表演是AV演员色情表演")
	fmt.Println(keywords)
}
