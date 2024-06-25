package pembelian

import (
	"errors"
	"time"
)

type PembelianService interface {
	CreatePembelian(idMaterial int, jumlah int, status string) (*PembelianTable, error)
	UpdatePembelian(id int, jumlah int, status string) (*PembelianTable, error)
	DeletePembelian(id int) error
	GetPembelianByID(id int) (*PembelianTable, error)
	GetAllPembelian() ([]PembelianTable, error)
}

type pembelianService struct {
	repository PembelianRepository
}

func NewPembelianService(repository PembelianRepository) *pembelianService {
	return &pembelianService{repository}
}

func (s *pembelianService) CreatePembelian(idMaterial int, jumlah int, status string) (*PembelianTable, error) {
	// Validasi input
	if idMaterial == 0 {
		return nil, errors.New("ID Material cannot be empty")
	}
	if jumlah <= 0 {
		return nil, errors.New("jumlah must be greater than zero")
	}

	// Buat objek PembelianTable baru
	pembelian := &PembelianTable{
		IdMaterial: idMaterial,
		Jumlah:     jumlah,
		Status:     status,
		CreatedAt:  time.Now(),
	}

	// Simpan ke database menggunakan repository
	if err := s.repository.CreatePembelian(pembelian); err != nil {
		return nil, err
	}

	return pembelian, nil
}

func (s *pembelianService) UpdatePembelian(id int, jumlah int, status string) (*PembelianTable, error) {
	// Validasi input
	if jumlah <= 0 {
		return nil, errors.New("jumlah must be greater than zero")
	}

	// Cari PembelianTable berdasarkan ID
	existingPembelian, err := s.repository.GetPembelianByID(id)
	if err != nil {
		return nil, err
	}

	// Update data PembelianTable
	existingPembelian.Jumlah = jumlah
	existingPembelian.Status = status

	// Simpan perubahan ke database menggunakan repository
	if err := s.repository.UpdatePembelian(id, existingPembelian); err != nil {
		return nil, err
	}

	// If status is "selesai", update the stock
	if status == "selesai" {
		if err := s.repository.UpdateStock(existingPembelian.IdMaterial, existingPembelian.Jumlah); err != nil {
			return nil, err
		}
	}

	return existingPembelian, nil
}

func (s *pembelianService) DeletePembelian(id int) error {
	// Hapus PembelianTable berdasarkan ID menggunakan repository
	if err := s.repository.DeletePembelian(id); err != nil {
		return err
	}
	return nil
}

func (s *pembelianService) GetPembelianByID(id int) (*PembelianTable, error) {
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
