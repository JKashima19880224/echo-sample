package config

type Config struct {
	Port string
}

func (c *Config) Build() *Config {
	return &Config{Port: "8080"}
}
