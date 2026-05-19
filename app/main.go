package main

import (
	"fmt"
	"notes-app/config"
	"notes-app/notes"
)

func main() {
	cfg := config.LoadConfig()
	store := notes.NewStore()
	service := notes.NewService(store)
	ui := notes.NewCLI(service)
	fmt.Println("Приложение запущено.")
	fmt.Printf("Приветствуем вас, %s!\n", cfg.DefaultUser)
	ui.Run()
}
