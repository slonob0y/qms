package service

import (
	"github.com/slonob0y/qms/models"
	"github.com/slonob0y/qms/repository"
	"github.com/slonob0y/qms/utils"
)

type BookService struct {
	bookRepo repository.BookRepoInterface
}

func NewBookService(bookRepo repository.BookRepoInterface) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

type BookServiceInterface interface {
	CreateBook(book models.SlotBooking) (data models.SlotBooking, err error)
	GetBank() ([]models.Bank, error)
	DeleteBook(status string) error
	GetBankById(id string) (models.Bank, error)
}

func (s *BookService) CreateBook(book models.SlotBooking) (data models.SlotBooking, err error) {
	day, hour := utils.GetDateAndHour()
	book.TanggalPelayanan = day
	book.JamPelayanan = hour

	result, err := s.bookRepo.CreateBook(book, day, hour)
	if err != nil {
		return result, err
	}

	return result, err
}

func (s *BookService) GetBank() ([]models.Bank, error) {
	banks, err := s.bookRepo.FindAllBank()

	return banks, err
}

func (s *BookService) DeleteBook(id string) error {
	err := s.bookRepo.DeleteBook(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *BookService) GetBankById(id string) (models.Bank, error) {
	movies, err := s.bookRepo.GetBankById(id)

	return movies, err
}
