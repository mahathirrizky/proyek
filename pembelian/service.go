package pembelian

import (
	"errors"
	"time"
)

type PembelianService interface {
	CreatePembelian(input CreatePembelianInput) (*PembelianTable, error)
	UpdatePembelian(id GetPembelianDetailInput, input UpdatePembelianInput) (*PembelianTable, error)
	DeletePembelian(id GetPembelianDetailInput) error
	GetPembelianByID(id GetPembelianDetailInput) (*PembelianTable, error)
	GetPembelianByProyekID(proyekID int) ([]PembelianTable, error)
	GetAllPembelian() ([]PembelianTable, error)
}

type pembelianService struct {
	repository PembelianRepository
}

func NewPembelianService(repository PembelianRepository) *pembelianService {
	return &pembelianService{repository}
}

func (s *pembelianService) CreatePembelian(input CreatePembelianInput) (*PembelianTable, error) {
	// Validasi input
	if input.IdMaterial == 0 {
		return nil, errors.New("ID Material cannot be empty")
	}
	if input.Jumlah <= 0 {
		return nil, errors.New("jumlah must be greater than zero")
	}

	// Buat objek PembelianTable baru
	pembelian := &PembelianTable{
		IdProyek:   input.IdProyek,
		IdMaterial: input.IdMaterial,
		Jumlah:     input.Jumlah,
		Status:     input.Status,
		CreatedAt:  time.Now(),
	}

	// Simpan ke database menggunakan repository
	if err := s.repository.CreatePembelian(pembelian); err != nil {
		return nil, err
	}

	return pembelian, nil
}

func (s *pembelianService) UpdatePembelian(id GetPembelianDetailInput, input UpdatePembelianInput) (*PembelianTable, error) {
	// Validasi input
	if input.Jumlah <= 0 {
		return nil, errors.New("jumlah must be greater than zero")
	}

	// Cari PembelianTable berdasarkan ID
	existingPembelian, err := s.repository.GetPembelianByID(id)
	if err != nil {
		return nil, err
	}

	// Update data PembelianTable
	existingPembelian.Jumlah = input.Jumlah
	existingPembelian.Status = input.Status

	// Simpan perubahan ke database menggunakan repository
	if err := s.repository.UpdatePembelian(id, existingPembelian); err != nil {
		return nil, err
	}

	// If status is "selesai", update the stock
	if input.Status == "selesai" {
		if err := s.repository.UpdateStock(existingPembelian.IdMaterial, existingPembelian.Jumlah); err != nil {
			return nil, err
		}
	}

	return existingPembelian, nil
}

func (s *pembelianService) DeletePembelian(id GetPembelianDetailInput) error {
	// Hapus PembelianTable berdasarkan ID menggunakan repository
	if err := s.repository.DeletePembelian(id); err != nil {
		return err
	}
	return nil
}

func (s *pembelianService) GetPembelianByID(id GetPembelianDetailInput) (*PembelianTable, error) {
	// Ambil PembelianTable berdasarkan ID menggunakan repository
	pembelian, err := s.repository.GetPembelianByID(id)
	if err != nil {
		return nil, err
	}
	return pembelian, nil
}

func (s *pembelianService) GetAllPembelian() ([]PembelianTable, error) {
	// Ambil semua PembelianTable menggunakan repository
	pembelian, err := s.repository.GetAllPembelian()
	if err != nil {
		return nil, err
	}
	return pembelian, nil
}

func (s *pembelianService) GetPembelianByProyekID(proyekID int) ([]PembelianTable, error) {
	return s.repository.GetPembelianByProyekID(proyekID)
}
