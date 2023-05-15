package copierx

import (
	"github.com/SpectatorNan/goutils/common/errorx"
	"github.com/jinzhu/copier"
	"strconv"
)

var I64ToString = copier.Option{
	IgnoreEmpty: false,
	DeepCopy:    true,
	Converters: []copier.TypeConverter{
		{
			SrcType: int64(0),
			DstType: copier.String,
			Fn: func(src interface{}) (interface{}, error) {
				i, ok := src.(int64)
				if !ok {
					return nil, errorx.TypeMismatchForConvertErr
				}
				return strconv.FormatInt(i, 10), nil
			},
		},
	},
}
