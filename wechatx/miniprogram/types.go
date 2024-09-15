package miniprogram

type ResponseErrCode interface {
	GetErrCode() int
	GetErrMsg() string
}

type ErrResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e ErrResponse) GetErrMsg() string {
	return e.ErrMsg
}
func (e ErrResponse) GetErrCode() int {
	return e.ErrCode
}

type AccessTokenResponse struct {
	ErrResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type JSApiTicketResponse struct {
	ErrResponse
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}
