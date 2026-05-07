package privacy

type MaskLevel int

const (
	MaskPublic MaskLevel = iota
	MaskInternal
	MaskSuperUser
)

type ViewerContext struct {
	Role  []string
	Level MaskLevel
}

type Desensitize interface {
	MakeDesensitize(ctx ViewerContext) any
	DesensitizeType() DesensitizeType
}

type DesensitizeType int

const (
	DesTypeObject DesensitizeType = iota
	DesTypeArray
	DesTypeMap
)
