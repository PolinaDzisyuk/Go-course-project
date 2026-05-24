package notes

import (
	"os"
	"testing"
)

func Test_Run_ExitScenario(t *testing.T) {
	testFile := "test_interface.json"
	defer os.Remove(testFile)

	store := NewStore(testFile)
	service := NewService(store)
	ui := NewCLI(service)
	inputFile := "test_input.txt"
	_ = os.WriteFile(inputFile, []byte("5\n"), 0644)
	defer os.Remove(inputFile)

	f, err := os.Open(inputFile)
	if err != nil {
		t.Fatalf("Не удалось создать тестовый ввод: %v", err)
	}
	defer f.Close()
	oldStdin := os.Stdin
	os.Stdin = f
	defer func() {
		os.Stdin = oldStdin
	}()
	ui.Run()
}

func Test_EditNoteScenario(t *testing.T) {
	testFile := "test_interface_edit.json"
	_ = os.Remove(testFile)
	defer os.Remove(testFile)

	store := NewStore(testFile)
	service := NewService(store)
	ui := NewCLI(service)

	_, _ = service.AddNote("Старый заголовок", "Старый текст", []string{"тест"})

	userInput := "5\n1\n\nНовый текст заметки\n\n6\n"

	inputFile := "test_edit_input.txt"
	_ = os.WriteFile(inputFile, []byte(userInput), 0644)
	defer os.Remove(inputFile)

	f, _ := os.Open(inputFile)
	defer f.Close()

	oldStdin := os.Stdin
	os.Stdin = f
	defer func() { os.Stdin = oldStdin }()

	ui.Run()
}
