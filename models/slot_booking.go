package models

type SlotBooking struct {
	ID       			uint   		`gorm:"primaryKey" json:"id_booking"`
	TanggalPelayanan 	int 		`json:"tanggal_pelayanan"`
	JamPelayanan 		int 		`json:"jam_pelayanan"`
	KeperluanLayanan	int 		`gorm:"primaryKey" json:"keperluan_layanan"`
	Status 				string 		`gorm:"primaryKey" json:"status"`
	BankID 				uint		`json:"id_bank_tujuan"`
	UserID 				uint 		`json:"id_user"`
}