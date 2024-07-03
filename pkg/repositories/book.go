package repositories

import (
	"mess-manager/pkg/domain"
	"mess-manager/pkg/models"

	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

func BookDBInstance(d *gorm.DB) domain.IBookRepo {
	return &bookRepo{
		db: d,
	}
}

func (repo *bookRepo) CreateBook(book *models.Book) error {
	if err := repo.db.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (repo *bookRepo) GetBooks(bookID uint) []models.Book {
	var Book []models.Book
	var err error

	if bookID != 0 {
		err = repo.db.Where("id = ?", bookID).Find(&Book).Error
	} else {
		err = repo.db.Find(&Book).Error
	}

	if err != nil {
		return []models.Book{}
	}
	return Book
}
