package common
const (
	CurrEnv = EnvTypeLocal
)

type EnvType int

const (
	EnvTypeLocal = iota
	EnvTypeProd
)

type ContextKey string