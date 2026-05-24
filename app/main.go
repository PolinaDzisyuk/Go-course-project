package main

import (
	"flag"
	"fmt"
	"notes-app/config"
	"notes-app/notes"
	"os"
)

func main() {
	helpFlag := flag.Bool("info", false, "Показать подробное пояснение к запуску проекта и архитектуре")
	StorageFile := flag.String("file", "notes.json", "Файл для хранения заметок")
	flag.Usage = func() {
		fmt.Println("<МЕНЕДЖЕР ЗАМЕТОК>\nинструкция по использованию")
		fmt.Println("\nДоступные флаги:")
		flag.PrintDefaults()
		fmt.Println("\nПример запуска:")
		fmt.Println("  go run ./app                  - Запуск приложения в интерактивном режиме")
		fmt.Println("  go run ./app --info           - Вывод пояснений по архитектуре проекта")
		fmt.Println("  go run ./app --help  	        - Показать это справочное окно")
		fmt.Println("  go run ./app --file=file.json - Запуск с пользовательским файлом для хранения заметок")
	}
	flag.Parse()
	if *helpFlag {
		printProjectInfo()
		os.Exit(0)
	}

	cfg := config.LoadConfig()
	store := notes.NewStore(*StorageFile)
	service := notes.NewService(store)
	ui := notes.NewCLI(service)

	fmt.Println("Приложение запущено.")
	fmt.Printf("Файл для сохранения заметок: %s\n\n", *StorageFile)
	fmt.Printf("Приветствуем вас, %s!\n", cfg.DefaultUser)
	ui.Run()
}

func printProjectInfo() {
	fmt.Println("Информация о проекте")
	fmt.Println("\nОкружение:")
	fmt.Println("   + для работы требуется установленный Go (версии 1.21+).")
	fmt.Println("   + данные сохраняются локально в корень проекта: notes.json.")
	fmt.Println("\nСлои архитектуры:")
	fmt.Println("   + models (types.go):    структура данных заметки.")
	fmt.Println("   + storage.go (Store):   сохранение JSON, потокобезопасность (sync.Mutex).")
	fmt.Println("   + service.go (Service): бизнес-логика и валидация.")
	fmt.Println("   + interface.go (CLI):   интерактивное консольное меню.")
	fmt.Println("\nКоманды запуска из корня проекта:")
	fmt.Println("   + go run ./app          - запуск приложения.")
}
