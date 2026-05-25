package notes

import (
	"errors"
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
	if !errors.Is(err, ErrEmptyTitle) {
		t.Fatalf("ожидалась ошибка ErrEmptyTitle, но получена: %v", err)
	}
}

func TestGetNotesByTag(t *testing.T) {
	testFile := "test_tags.json"
	defer os.Remove(testFile)

	store := NewStore(testFile)
	service := NewService(store)

	var err error

	_, err = service.AddNote("заметка 1", "текст", []string{"работа", "важное"})
	if err != nil {
		t.Fatalf("Не удалось подготовить окружение теста (Заметка 1): %v", err)
	}
	_, err = service.AddNote("заметка 2", "текст", []string{"учёба"})
	if err != nil {
		t.Fatalf("Не удалось подготовить окружение теста (Заметка 2): %v", err)
	}
	_, err = service.AddNote("заметка 3", "текст", []string{"работа"})
	if err != nil {
		t.Fatalf("Не удалось подготовить окружение теста (Заметка 3): %v", err)
	}

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

	_, err := service.AddNote("купить хлеб", "в ярче", []string{"быт"})
	if err != nil {
		t.Fatalf("Не удалось создать тестовую заметку для удаления: %v", err)
	}

	err = service.DeleteNote(1)
	if err != nil {
		t.Fatalf("не удалось удалить существующую заметку: %v", err)
	}
	allNotes := service.GetNotesByTag("")
	if len(allNotes) != 0 {
		t.Errorf("ожидалось, что база будет пустой после удаления, но там осталось заметок: %d", len(allNotes))
	}

	err = service.DeleteNote(999)
	if !errors.Is(err, ErrNoteNotFound) {
		t.Errorf("ожидалась ошибка ErrNoteNotFound при удалении несуществующей заметки, но получена: %v", err)
	}
}

func TestUpdateNote_EmptyTitleError(t *testing.T) {
	testFile := "test_update_err.json"
	defer os.Remove(testFile)

	store := NewStore(testFile)
	service := NewService(store)

	_, err := service.AddNote("оригинальный заголовок", "текст", []string{})
	if err != nil {
		t.Fatalf("не удалось создать заметку для теста обновления: %v", err)
	}
	err = service.UpdateNote(1, "", "новый текст", []string{})
	if !errors.Is(err, ErrEmptyTitle) {
		t.Errorf("ожидалась ошибка ErrEmptyTitle при обновлении с пустым заголовком, но получена: %v", err)
	}
}

func TestUpdateNote_NotFound(t *testing.T) {
	testFile := "test_update_notfound.json"
	defer os.Remove(testFile)

	store := NewStore(testFile)
	service := NewService(store)

	err := service.UpdateNote(999, "заголовок", "текст", []string{})
	if !errors.Is(err, ErrNoteNotFound) {
		t.Errorf("ожидалась ошибка ErrNoteNotFound при обновлении несуществующей заметки, но получена: %v", err)
	}
}

func TestGetNoteByID_NotFound(t *testing.T) {
	testFile := "test_getbyid.json"
	defer os.Remove(testFile)

	store := NewStore(testFile)
	service := NewService(store)

	_, err := service.GetNoteByID(42)
	if !errors.Is(err, ErrNoteNotFound) {
		t.Errorf("ожидалась ошибка ErrNoteNotFound, но получена: %v", err)
	}
}
