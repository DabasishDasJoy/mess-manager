package services

import (
	"errors"
	"mess-manager/pkg/domain"
	"mess-manager/pkg/models"
	"mess-manager/pkg/types"
)

type BookService struct {
	repo domain.IBookRepo
}

func BookServiceInstance(bookRepo domain.IBookRepo) domain.IBookService {
	return &BookService{
		repo: bookRepo,
	}
}

// func (service *BookService) GetBooks(bookId uint) ([]types.BookRequest, error)
func (service *BookService) CreateBook(book *models.Book) error {
	if err := service.repo.CreateBook(book); err != nil {
		return errors.New("book was not created")
	}
	return nil
}

func (service *BookService) GetBooks(bookID uint) ([]types.BookRequest, error) {
	var allBooks []types.BookRequest
	book := service.repo.GetBooks(bookID)
	if len(book) == 0 {
		return nil, errors.New("no book found")
	}
	for _, val := range book {
		allBooks = append(allBooks, types.BookRequest{
			ID:          val.ID,
			BookName:    val.BookName,
			Author:      val.Author,
			Publication: val.Publication,
		})
	}
	return allBooks, nil
}
