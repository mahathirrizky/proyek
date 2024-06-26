package permintaan

import (
	"errors"
	"proyek/material"
	"time"
)

type PermintaanService interface {
	CreatePermintaan(input CreatePermintaanInput) (*PermintaanTable, error)
	UpdatePermintaan(id int, input UpdatePermintaanInput) (*PermintaanTable, error)
	DeletePermintaan(id int) error
	GetPermintaanByID(id int) (*PermintaanTable, error)
	GetPermintaanByProyekID(proyekID int) ([]PermintaanTable, error) // Perubahan di sini
	GetAllPermintaan() ([]PermintaanTable, error)
}

type permintaanService struct {
	repository         PermintaanRepository
	materialRepository material.Repository
}

func NewPermintaanService(repository PermintaanRepository, materialRepository material.Repository) PermintaanService {
	return &permintaanService{repository, materialRepository}
}

func (s *permintaanService) CreatePermintaan(input CreatePermintaanInput) (*PermintaanTable, error) {
	permintaan := &PermintaanTable{
		IdProyek:  input.IdProyek,
		Status:    input.Status,
		CreatedAt: time.Now(),
	}

	if err := s.repository.CreatePermintaan(permintaan); err != nil {
		return nil, err
	}

	for _, detail := range input.Details {
		permintaanDetail := &PermintaanDetailTable{
			IdPermintaan: permintaan.IdPermintaan,
			IdMaterial:   detail.IdMaterial,
			Jumlah:       detail.Jumlah,
		}

		if err := s.repository.CreatePermintaanDetail(permintaanDetail); err != nil {
			return nil, err
		}

		// Reduce material stock
		stokUpdate := material.StokTable{
			IdMaterial: detail.IdMaterial,
			Jumlah:     -detail.Jumlah, // Decrease stock
		}
		if _, err := s.materialRepository.UpdateStok(stokUpdate); err != nil {
			return nil, err
		}
		permintaan.Details = append(permintaan.Details, *permintaanDetail)
	}

	return permintaan, nil
}

func (s *permintaanService) UpdatePermintaan(id int, input UpdatePermintaanInput) (*PermintaanTable, error) {
	// Retrieve permintaan by ID
	permintaan, err := s.repository.GetPermintaanByID(id)
	if err != nil {
		return nil, err
	}

	// Check if permintaan status is "batal"
	if permintaan.Status == "batal" {
		return nil, errors.New("cannot update permintaan with status 'batal'")
	}

	// Backup original details for rollback if input status is "batal"
	var originalDetails []PermintaanDetailTable
	if input.Status == "batal" {
		originalDetails = make([]PermintaanDetailTable, len(permintaan.Details))
		copy(originalDetails, permintaan.Details)
	}

	// Update permintaan details based on input
	for i, detail := range input.Details {
		permintaan.Details[i].Jumlah = detail.Jumlah
	}

	// Check if status is changing to "batal"
	if input.Status == "batal" {
		// Rollback stock to original state before update
		for _, detail := range originalDetails {
			stokUpdate := material.StokTable{
				IdMaterial: detail.IdMaterial,
				Jumlah:     detail.Jumlah, // Return stock
			}
			if _, err := s.materialRepository.UpdateStok(stokUpdate); err != nil {
				return nil, err
			}
		}

		// Update permintaan status
		permintaan.Status = input.Status
	}

	// Persist changes to permintaan including details and status
	if err := s.repository.UpdatePermintaan(id, permintaan); err != nil {
		return nil, err
	}

	return permintaan, nil
}

func (s *permintaanService) DeletePermintaan(id int) error {
	return s.repository.DeletePermintaan(id)
}

func (s *permintaanService) GetPermintaanByID(id int) (*PermintaanTable, error) {
	return s.repository.GetPermintaanByID(id)
}

func (s *permintaanService) GetPermintaanByProyekID(proyekID int) ([]PermintaanTable, error) {
	permintaans, err := s.repository.GetPermintaanByProyekID(proyekID)
	if err != nil {
		return nil, err
	}

	return permintaans, nil // Mengembalikan slice []PermintaanTable langsung
}

func (s *permintaanService) GetAllPermintaan() ([]PermintaanTable, error) {
	return s.repository.GetAllPermintaan()
}
