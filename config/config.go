package config

type Config struct {
	AppName     string
	DefaultUser string
}

func LoadConfig() *Config {
	return &Config{
		AppName:     "--- МЕНЕДЖЕР ЗАМЕТОК ---",
		DefaultUser: "Пользователь",
	}
}
