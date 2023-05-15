package tools

import "strconv"

type ID int64

func IDParse(str string) ID {
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return ID(i64)
}

func IDFrom(i64 int64) ID {
	return ID(i64)
}

func (id ID) ToString() string {
	return strconv.FormatInt(int64(id), 10)
}

func (id ID) GetInt64() int64 {
	return int64(id)
}

func StringToInt64(str string) int64 {
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i64
}

func Int64ToString(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}
