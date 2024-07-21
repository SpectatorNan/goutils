package imagex

const (
	WebpMax = 16383
	AvifMax = 65536
)

type ExtraParams struct {
	Width     int // in px
	Height    int // in px
	MaxWidth  int // in px
	MaxHeight int // in px
}
