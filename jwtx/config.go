package jwtx

type Config struct {
	//AccessSecret string
	AccessExpire int64
	Issuer       string
	BufferTime   int64 `json:",default=86400"`
}
type Login struct {
	Multipoint bool `json:",default=true"`
}
