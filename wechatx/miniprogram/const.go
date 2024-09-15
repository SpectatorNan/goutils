package miniprogram

const (
	apiServer = "https://api.weixin.qq.com"
)

const (
	path_token  = "/cgi-bin/token"
	path_ticket = "/cgi-bin/ticket/getticket"
)

const (
	grantType_ClientCredential = "client_credential"
)

var (
	paramsAccessToken = func(appId, secret string) map[string]string {
		m := make(map[string]string)
		m["grant_type"] = grantType_ClientCredential
		m["appid"] = appId
		m["secret"] = secret
		return m
	}
	paramsGetTicket = func(token string) map[string]string {
		m := make(map[string]string)
		m["access_token"] = token
		m["type"] = "jsapi"
		return m
	}
)

type CacheType int

const (
	CacheTypeMemory = iota
	CacheTypeRedis
)
const (
	cacheKeyWithAccessToken = "wechatx:accessToken"
	cacheKeyWithJSApiTicket = "wechatx:jsapi:ticket"
)

type CacheItem struct {
	ExpireUnix int64
	Data       string
}
