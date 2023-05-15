package tools

import (
	"testing"
)

func TestWeekTimeStr(t *testing.T) {
	workTime := WeekWorkTime([]int32{1, 2, 3})
	t.Log(workTime)
}
