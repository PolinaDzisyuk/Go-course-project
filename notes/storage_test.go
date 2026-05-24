package notes

import (
	"os"
	"testing"
)

func Test_FullScenario(t *testing.T) {
	testFile := "test_interface_full.json"
	_ = os.Remove(testFile)
	defer os.Remove(testFile)

	store := NewStore(testFile)
	service := NewService(store)
	ui := NewCLI(service)

	userInput := "2\nКупить хлеб\nВ магазине у дома\nбыт, продукты\n" +
		"1\n" +
		"3\nбыт\n" +
		"4\n1\n" +
		"5\n"

	inputFile := "test_interaction.txt"
	_ = os.WriteFile(inputFile, []byte(userInput), 0644)
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

func Test_InvalidSubmenus(t *testing.T) {
	testFile := "test_interface_invalid.json"
	defer os.Remove(testFile)

	store := NewStore(testFile)
	service := NewService(store)
	ui := NewCLI(service)

	userInput := "9\n4\nabc\n5\n"

	inputFile := "test_invalid.txt"
	_ = os.WriteFile(inputFile, []byte(userInput), 0644)
	defer os.Remove(inputFile)

	f, _ := os.Open(inputFile)
	defer f.Close()

	oldStdin := os.Stdin
	os.Stdin = f
	defer func() { os.Stdin = oldStdin }()

	ui.Run()
}
