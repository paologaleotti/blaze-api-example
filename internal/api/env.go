package api

import "blaze/pkg/util"

type EnvConfig struct {
	DatabaseUrl   string
	EnableMetrics bool
}

func InitEnv() EnvConfig {
	return EnvConfig{
		DatabaseUrl:   util.GetEnvOrPanic("DATABASE_URL"),
		EnableMetrics: util.GetEnvOrDefault("ENABLE_METRICS", "") == "true",
	}
}
