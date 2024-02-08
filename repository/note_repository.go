package repository

import "github.com/ice-009/go-fiber-crud/model"

type NoteRepository interface {
	Save(note model.Note)
	Update(note model.Note)
	Delete(note model.Note)
	FindById(id int) model.Note
	FindAll() []model.Note
}
