package config

type Captcha struct {
	KeyLong   int `json:",default=6"`
	ImgWidth  int `json:",default=240"`
	ImgHeight int `json:",default=80"`
}
