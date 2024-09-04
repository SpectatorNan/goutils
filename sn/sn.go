package sn

import (
	"fmt"
	"github.com/SpectatorNan/goutils/tools"
	"time"
)

type SnPrefix string

const (
	SN_PREFIX_VIP_ORDER  SnPrefix = "OV" //会员订单前缀 order vip
	SN_PREFIX_VIP_REFUND          = "VRF"
)

func GenSn(snPrefix SnPrefix) string {
	return fmt.Sprintf("%s%s%s", snPrefix, time.Now().Format("20060102150405"), tools.Krand(8, tools.KC_RAND_KIND_NUM))
}
