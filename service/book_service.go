package service

import (
	"errors"

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
	ValidateBookByDay() error
}

func (s *BookService) CreateBook(book models.SlotBooking) (data models.SlotBooking, err error) {
	date := utils.FormatGetDate()
	hour := utils.FormatGetHour()

	book.TanggalPelayanan = date
	book.JamPelayanan = hour

	result, err := s.bookRepo.CreateBook(book, date, hour)
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

func (s *BookService) ValidateBookByDay() error {
	date := utils.FormatGetDate()
	count, err := s.bookRepo.GetBookByDate(date)
	if err != nil {
		return err
	}

	if count > 40 {
		return errors.New("slot booking di hari ini sudah penuh")
	}

	return nil

}
