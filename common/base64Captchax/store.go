package base64Captchax

type Store interface {
	SetCaptcha(prefix string, id string, value string) error
	GetCaptcha(prefix string, id string, clear bool) string
	VerifyCaptcha(prefix string, id, answer string, clear bool) bool

	Set(id string, value string) error
	Get(id string, clear bool) string
	Verify(id, answer string, clear bool) bool
}
