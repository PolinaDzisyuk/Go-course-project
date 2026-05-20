package notes

import "notes-app/models"

type Repository interface {
	Save(note models.Note) models.Note
	FindAll() []models.Note
	DeleteByID(id int) error
}
