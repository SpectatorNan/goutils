package tools

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Int32 []int32

func (f Int32) Len() int {
	return len(f)
}

func (f Int32) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f Int32) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func GetWeekday(day int32) string {
	switch day {
	case 1:
		return "周一"
	case 2:
		return "周二"
	case 3:
		return "周三"
	case 4:
		return "周四"
	case 5:
		return "周五"
	case 6:
		return "周六"
	case 7:
		return "周日"
	default:
		return ""
	}
}

func WeekDayToString(day time.Weekday) string {
	switch day {
	case 1:
		return "一"
	case 2:
		return "二"
	case 3:
		return "三"
	case 4:
		return "四"
	case 5:
		return "五"
	case 6:
		return "六"
	case 0:
		return "日"
	default:
		return ""
	}
}

func WeekDayToInt(day time.Weekday) int {
	switch day {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return 4
	case 5:
		return 5
	case 6:
		return 6
	case 0:
		return 7
	default:
		return 0
	}
}
func WeekWorkTime(workDayArr []int32) string {
	if len(workDayArr) < 1 {
		return "休息中"
	}
	dayoff := []int32{1, 2, 3, 4, 5, 6, 7}
	if len(workDayArr) > 7 {
		return ""
	}
	daySlice := Int32(workDayArr)
	sort.Sort(daySlice)
	workTime := ""
	var cdate = int32(0)
	var start = int32(-1)
	var end = int32(-1)
	var isBreak = false
	for _, day := range daySlice {
		if day < 1 || day > 7 {
			continue
		}
		dayoff = RemoveInt32Slice(dayoff, day)
		if cdate == 0 {
			start = day
			cdate = day
		} else {
			if cdate != day-1 {
				isBreak = true
				end = cdate
				cdate = 0
			} else {
				cdate = day
				end = cdate
			}
		}
	}
	if start > end {
		temp := start
		start = end
		end = temp
	}
	if start == 1 && end == 7 {
		workTime = "周一至周日"
	} else if start < end {
		if start == 1 && end == 5 && !isBreak {
			workTime = "工作日营业"
		} else if !isBreak {
			workTime = fmt.Sprintf("%s至%s", GetWeekday(start), GetWeekday(end))
		} else
		//if (start + 2) == end {
		//	wday := GetWeekday(start + 1)
		//	workTime = fmt.Sprintf("%s休息", wday)
		//} else
		{
			if len(dayoff) < 4 {
				wday := ""
				dayoffStrs := make([]string, 0)
				for _, day := range dayoff {
					dayoffStrs = append(dayoffStrs, GetWeekday(day))
				}
				wday = fmt.Sprintf("%s休息", strings.Join(dayoffStrs, ","))
				workTime = wday
			} else {
				wday := ""
				for _, r := range daySlice {
					rday := GetWeekday(r)
					wday += rday
					wday += ","
				}
				if strings.HasSuffix(wday, ",") {
					wday = wday[:len(wday)-1]
				}
				workTime = wday
			}
		}
	}
	return workTime
}
