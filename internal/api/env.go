package api

import "blaze/pkg/util"

type EnvConfig struct {
	DatabaseUrl string
}

var envVarMappings = util.EnvMapping{
	"DATABASE_URL": &env.DatabaseUrl,
}

var env = &EnvConfig{}

func InitEnv() *EnvConfig {
	for key, goVar := range envVarMappings {
		*goVar = util.GetEnvOrPanic(key)
	}

	return env
}
