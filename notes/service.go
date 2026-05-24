package notes

import (
	"errors"
	"notes-app/models"
	"time"
)

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{store: store}
}

func (s *Service) AddNote(title, content string, tags []string) (models.Note, error) {
	if title == "" {
		return models.Note{}, errors.New("Заголовок не может быть пустым")
	}

	note := models.Note{
		Title:   title,
		Content: content,
		Tags:    tags,
		Date:    time.Now(),
	}

	savedNote := s.store.Save(note)
	return savedNote, nil
}

func (s *Service) GetNotesByTag(tag string) []models.Note {
	allNotes := s.store.FindAll()
	if tag == "" {
		return allNotes
	}

	var filtered []models.Note
	for _, note := range allNotes {
		for _, t := range note.Tags {
			if t == tag {
				filtered = append(filtered, note)
				break
			}
		}
	}
	return filtered
}

func (s *Service) DeleteNote(id int) error {
	return s.store.DeleteByID(id)
}

func (s *Service) UpdateNote(id int, title, content string, tags []string) error {
	if title == "" {
		return errors.New("Заголовок не может быть пустым")
	}
	updatedNote := models.Note{
		Title:   title,
		Content: content,
		Tags:    tags,
	}
	return s.store.Update(id, updatedNote)
}

func (s *Service) GetNoteByID(id int) (models.Note, error) {
	return s.store.FindByID(id)
}
