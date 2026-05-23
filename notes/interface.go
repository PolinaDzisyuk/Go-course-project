package notes

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CLI struct {
	service *Service
}

func NewCLI(service *Service) *CLI {
	return &CLI{service: service}
}

func (c *CLI) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- МЕНЕДЖЕР ЗАМЕТОК ---")
		fmt.Println("1. Показать все заметки")
		fmt.Println("2. Добавить заметку")
		fmt.Println("3. Найти заметки по тегу")
		fmt.Println("4. Удалить заметку по ID")
		fmt.Println("5. Редактировать заметку")
		fmt.Println("6. Выход")
		fmt.Print("Выберите действие (1-6): ")

		if !scanner.Scan() {
			break
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			c.showNotes("")
		case "2":
			c.addNote(scanner)
		case "3":
			fmt.Print("Введите тег для поиска: ")
			scanner.Scan()
			tag := strings.TrimSpace(scanner.Text())
			c.showNotes(tag)
		case "4":
			c.deleteNote(scanner)
		case "5":
			c.editNote(scanner)
		case "6":
			fmt.Println("До свидания!!")
			return

		default:
			fmt.Println("Ошибка: Неверный пункт меню. Попробуйте снова.")
		}
	}
}

func (c *CLI) showNotes(tag string) {
	notesList := c.service.GetNotesByTag(tag)
	if len(notesList) == 0 {
		fmt.Println("\nЗаметок пока нет или ничего не найдено.")
		return
	}

	fmt.Println("\n=== СПИСОК ЗАМЕТОК ===")
	for _, note := range notesList {
		fmt.Printf("[%d] %s\n", note.ID, note.Title)
		fmt.Printf("    Текст: %s\n", note.Content)
		if len(note.Tags) > 0 {
			fmt.Printf("    Теги: %s\n", strings.Join(note.Tags, ", "))
		}
		fmt.Printf("    Дата: %s\n", note.Date.Format("02.01.2006 15:04"))
		fmt.Println("----------------------")
	}
}

func (c *CLI) addNote(scanner *bufio.Scanner) {
	fmt.Print("Введите заголовок заметки: ")
	scanner.Scan()
	title := strings.TrimSpace(scanner.Text())

	fmt.Print("Введите текст заметки: ")
	scanner.Scan()
	content := strings.TrimSpace(scanner.Text())

	fmt.Print("Введите теги через запятую (например: учёба,go): ")
	scanner.Scan()
	tagsInput := strings.TrimSpace(scanner.Text())

	var tags []string
	if tagsInput != "" {
		parts := strings.Split(tagsInput, ",")
		for _, p := range parts {
			tags = append(tags, strings.TrimSpace(p))
		}
	}

	note, err := c.service.AddNote(title, content, tags)
	if err != nil {
		fmt.Printf("Ошибка при добавлении: %v\n", err)
		return
	}
	fmt.Println("Новая заметка!")
	fmt.Printf("Создана заметка с ID: %d\n", note.ID)
}

func (c *CLI) deleteNote(scanner *bufio.Scanner) {
	fmt.Print("Введите ID заметки для удаления: ")
	scanner.Scan()
	idInput := strings.TrimSpace(scanner.Text())

	id, err := strconv.Atoi(idInput)
	if err != nil {
		fmt.Println("Ошибка: ID должен быть числом.")
		return
	}

	err = c.service.DeleteNote(id)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Println("Заметка успешно удалена!")
}

func (c *CLI) editNote(scanner *bufio.Scanner) {
	fmt.Print("Введите ID заметки для редактирования: ")
	scanner.Scan()
	idInput := strings.TrimSpace(scanner.Text())

	id, err := strconv.Atoi(idInput)
	if err != nil {
		fmt.Println("Ошибка: ID должен быть числом.")
		return
	}

	fmt.Print("Введите новый заголовок заметки: ")
	scanner.Scan()
	title := strings.TrimSpace(scanner.Text())

	fmt.Print("Введите новый текст заметки: ")
	scanner.Scan()
	content := strings.TrimSpace(scanner.Text())

	fmt.Print("Введите новые теги через запятую: ")
	scanner.Scan()
	tagsInput := strings.TrimSpace(scanner.Text())

	var tags []string
	if tagsInput != "" {
		parts := strings.Split(tagsInput, ",")
		for _, p := range parts {
			tags = append(tags, strings.TrimSpace(p))
		}
	}

	err = c.service.UpdateNote(id, title, content, tags)
	if err != nil {
		fmt.Printf("Ошибка при редактировании: %v\n", err)
		return
	}
	fmt.Println("Заметка успешно обновлена!!")
}
