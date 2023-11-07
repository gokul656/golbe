package config

var EnvConf = LoadEnvConfig()

type EnvConfig struct {
	Port int
}

func LoadEnvConfig() *EnvConfig {
	return &EnvConfig{
		Port: 5150,
	}
}
