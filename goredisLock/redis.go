package goredisLock

import gred "github.com/redis/go-redis/v9"

type (
	Script = gred.Script
)

func NewScript(script string) *Script {
	return gred.NewScript(script)
}
