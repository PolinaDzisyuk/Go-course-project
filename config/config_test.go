package config

import "testing"

func TestLoadConfig(t *testing.T) {
	cfg := LoadConfig()
	if cfg.AppName == "" {
		t.Error("имя приложения не должно быть пустым")
	}
	if cfg.DefaultUser != "Пользователь" {
		t.Errorf("ожидался пользователь 'Пользователь', получен: %s", cfg.DefaultUser)
	}
}
