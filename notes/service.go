package notes

import (
	"errors"
	"notes-app/models"
	"time"
)

var ErrEmptyTitle = errors.New("заголовок не может быть пустым")
var ErrNoteNotFound = errors.New("заметка не найдена")

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{store: store}
}

func (s *Service) AddNote(title, content string, tags []string) (models.Note, error) {
	if title == "" {
		return models.Note{}, ErrEmptyTitle
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
	err := s.store.DeleteByID(id)
	if err != nil {
		return errors.Join(ErrNoteNotFound, err)
	}
	return nil
}

func (s *Service) UpdateNote(id int, title, content string, tags []string) error {
	if title == "" {
		return ErrEmptyTitle
	}
	updatedNote := models.Note{
		Title:   title,
		Content: content,
		Tags:    tags,
	}
	err := s.store.Update(id, updatedNote)
	if err != nil {
		return errors.Join(ErrNoteNotFound, err)
	}
	return nil
}

func (s *Service) GetNoteByID(id int) (models.Note, error) {
	note, err := s.store.FindByID(id)
	if err != nil {
		return models.Note{}, errors.Join(ErrNoteNotFound, err)
	}
	return note, nil

}
