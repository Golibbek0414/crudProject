package config

type Config struct {
	HTTPPort string
	TimeOut  int
}

func Load() Config {
	config := Config{":8000", 7}

	return config
}
