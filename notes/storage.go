package notes

import (
	"errors"
	"notes-app/models"
	"slices"
)

type Store struct {
	notes  []models.Note
	nextID int
}

func NewStore() *Store {
	return &Store{
		notes:  make([]models.Note, 0),
		nextID: 1,
	}
}

func (s *Store) Save(note models.Note) models.Note {
	note.ID = s.nextID
	s.nextID++
	s.notes = append(s.notes, note)
	return note
}

func (s *Store) FindAll() []models.Note {
	return s.notes
}

func (s *Store) DeleteByID(id int) error {
	for i, note := range s.notes {
		if note.ID == id {
			s.notes = slices.Delete(s.notes, i, i+1)
			return nil
		}
	}
	return errors.New("Err: Заметка не найдена")
}
