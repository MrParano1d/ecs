package core

const (
	EnvDebug   = "debug"
	EnvRelease = "release"
)

type ConfigRes struct {
	Environment string
}

func NewConfig() *ConfigRes {
	return &ConfigRes{
		Environment: EnvRelease,
	}
}

func NewDebugConfig() *ConfigRes {
	c := NewConfig()
	c.Environment = EnvDebug
	return c
}
