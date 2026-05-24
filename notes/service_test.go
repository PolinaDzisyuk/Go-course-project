package notes

import (
	"os"
	"testing"
)

func TestAddNote_Success(t *testing.T) {
	testFile := "test_notes.json"
	defer os.Remove(testFile)

	store := NewStore(testFile)
	service := NewService(store)

	title := "тестовый заголовок"
	content := "тестовое содержимое"
	tags := []string{"тест", "go"}

	note, err := service.AddNote(title, content, tags)

	if err != nil {
		t.Fatalf("ожидался успешный запуск, но получена ошибка: %v", err)
	}

	if note.ID != 1 {
		t.Errorf("ожидался ID равный 1, но получен: %d", note.ID)
	}

	if note.Title != title {
		t.Errorf("ожидался заголовок '%s', но получен: '%s'", title, note.Title)
	}
}

func TestAddNote_EmptyTitleError(t *testing.T) {
	testFile := "test_notes_err.json"
	defer os.Remove(testFile)

	store := NewStore(testFile)
	service := NewService(store)

	_, err := service.AddNote("", "текст без заголовка", []string{})

	if err == nil {
		t.Fatal("ожидалась ошибка обработки пустого заголовка, но метод вернул err = nil")
	}

	expectedError := "Заголовок не может быть пустым"
	if err.Error() != expectedError {
		t.Errorf("ожидался текст ошибки '%s', но получен: '%s'", expectedError, err.Error())
	}
}

func TestGetNotesByTag(t *testing.T) {
	testFile := "test_tags.json"
	defer os.Remove(testFile)

	store := NewStore(testFile)
	service := NewService(store)

	_, _ = service.AddNote("заметка 1", "текст", []string{"работа", "важное"})
	_, _ = service.AddNote("заметка 2", "текст", []string{"учёба"})
	_, _ = service.AddNote("заметка 3", "текст", []string{"работа"})
	workNotes := service.GetNotesByTag("работа")
	if len(workNotes) != 2 {
		t.Errorf("ожидалось 2 заметки с тегом 'работа', найдено: %d", len(workNotes))
	}

	emptyNotes := service.GetNotesByTag("отпуск")
	if len(emptyNotes) != 0 {
		t.Errorf("ожидалось 0 заметок с тегом 'отпуск', найдено: %d", len(emptyNotes))
	}
	allNotes := service.GetNotesByTag("")
	if len(allNotes) != 3 {
		t.Errorf("ожидалось получить все 3 заметки при пустом теге, получено: %d", len(allNotes))
	}
}

func TestDeleteNote(t *testing.T) {
	testFile := "test_delete.json"
	defer os.Remove(testFile)

	store := NewStore(testFile)
	service := NewService(store)
	_, _ = service.AddNote("купить хлеб", "в ярче", []string{"быт"})

	err := service.DeleteNote(1)
	if err != nil {
		t.Fatalf("не удалось удалить существующую заметку: %v", err)
	}
	allNotes := service.GetNotesByTag("")
	if len(allNotes) != 0 {
		t.Errorf("ожидалось, что база будет пустой после удаления, но там осталось заметок: %d", len(allNotes))
	}
	err = service.DeleteNote(999)
	if err == nil {
		t.Error("ожидалась ошибка при удалении несуществующей заметки с ID 999, но err = nil")
	}
}
