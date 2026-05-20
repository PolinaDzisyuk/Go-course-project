package notes

import (
	"encoding/json"
	"errors"
	"notes-app/models"
	"os"
	"slices"
	"sync"
)

type Store struct {
	filename string
	notes    []models.Note
	nextID   int
	mutex    sync.Mutex
}

func NewStore(filename string) *Store {
	s := &Store{
		filename: filename,
		notes:    make([]models.Note, 0),
		nextID:   1,
	}
	s.loadFromFile()
	return s
}

func (s *Store) loadFromFile() {
	file, err := os.ReadFile(s.filename)
	if err != nil {
		return
	}
	var loadedNotes []models.Note
	if err := json.Unmarshal(file, &loadedNotes); err == nil {
		s.notes = loadedNotes
		for _, note := range s.notes {
			if note.ID >= s.nextID {
				s.nextID = note.ID + 1
			}
		}
	}
}

func (s *Store) saveToFile() error {
	data, err := json.MarshalIndent(s.notes, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filename, data, 0644)
}

func (s *Store) Save(note models.Note) models.Note {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	note.ID = s.nextID
	s.nextID++
	s.notes = append(s.notes, note)
	_ = s.saveToFile()
	return note
}

func (s *Store) FindAll() []models.Note {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.notes
}

func (s *Store) DeleteByID(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for i, note := range s.notes {
		if note.ID == id {
			s.notes = slices.Delete(s.notes, i, i+1)
			_ = s.saveToFile()
			return nil
		}
	}
	return errors.New("Err: Заметка не найдена")
}
