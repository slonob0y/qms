package repository

import (
	"github.com/slonob0y/qms/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

type BookRepoInterface interface {
	CreateBook(book models.SlotBooking, day, hour string) (models.SlotBooking, error)
	FindAllBank() ([]models.Bank, error)
	DeleteBook(status string) error
	GetBankById(id string) (models.Bank, error)
}

func (r *BookRepository) CreateBook(book models.SlotBooking, day, hour string) (models.SlotBooking, error) {
	result := r.db.Create(&book)
	if result.Error != nil {
		return book, result.Error
	}

	return book, nil
}

// func(r *BookRepository) CreateBook(book *models.SlotBooking) error {
// 	query := `INSERT INTO slot_bookings (id_booking, tanggal_pelayanan, jam_pelayanan, keperluan_layanan, status, id_bank_tujuan, id_user)
// 				VALUES ($1, $2, $3, $4, $5, $6, $7)`
// 	_, err := r.db.Exec(query, book.ID, book.TanggalPelayanan, book.JamPelayanan, book.KeperluanLayanan, book.Status, book.BankID, book.UserID)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func(r *BookRepository) CreateFullBook(book models.SlotBooking) (models.SlotBooking, error) {
// 	trx := r.db.Create(&book)

// 	return book, trx.Error
// }

func (r *BookRepository) FindAllBank() ([]models.Bank, error) {
	var banks []models.Bank
	findResult := r.db.Find(&banks)

	return banks, findResult.Error
}

func (r *BookRepository) DeleteBook(id string) error {
	var book models.SlotBooking
	findResult := r.db.Unscoped().Where("id_booking = ?", id).Delete(&book)
	if findResult.Error != nil {
		return findResult.Error
	}

	if findResult.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *BookRepository) GetBankById(id string) (models.Bank, error) {
	var bank models.Bank
	findResult := r.db.Where("id = ?", id).First(&bank)
	return bank, findResult.Error
}
