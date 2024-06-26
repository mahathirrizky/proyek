package permintaan

import (
	"gorm.io/gorm"
)

type PermintaanRepository interface {
	CreatePermintaan(data *PermintaanTable) error
	CreatePermintaanDetail(detail *PermintaanDetailTable) error
	UpdatePermintaan(id int, data *PermintaanTable) error
	DeletePermintaan(id int) error
	GetPermintaanByID(id int) (*PermintaanTable, error)
	GetPermintaanByProyekID(proyekID int) ([]PermintaanTable, error)
	GetAllPermintaan() ([]PermintaanTable, error)
}

type permintaanRepository struct {
	db *gorm.DB
}

func NewPermintaanRepository(db *gorm.DB) PermintaanRepository {
	return &permintaanRepository{db}
}

func (repo *permintaanRepository) CreatePermintaan(data *PermintaanTable) error {
	return repo.db.Create(data).Error
}

func (r *permintaanRepository) CreatePermintaanDetail(detail *PermintaanDetailTable) error {
	return r.db.Create(detail).Error
}

func (repo *permintaanRepository) UpdatePermintaan(id int, data *PermintaanTable) error {
	// Begin transaction to ensure atomic updates
	tx := repo.db.Begin()

	// Update permintaan
	if err := tx.Model(&PermintaanTable{}).Where("id_permintaan = ?", id).Updates(data).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update permintaan details if available
	if len(data.Details) > 0 {
		for _, detail := range data.Details {
			if err := tx.Model(&PermintaanDetailTable{}).Where("id_permintaan = ?", detail.IdPermintaan).Updates(&detail).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	// Commit transaction if all updates succeed
	return tx.Commit().Error
}

func (repo *permintaanRepository) DeletePermintaan(id int) error {
	return repo.db.Where("id_permintaan = ?", id).Delete(&PermintaanTable{}).Error
}

func (repo *permintaanRepository) GetPermintaanByID(id int) (*PermintaanTable, error) {
	var permintaan PermintaanTable
	err := repo.db.Where("id_permintaan = ?", id).Preload("Details").First(&permintaan).Error
	return &permintaan, err
}

func (repo *permintaanRepository) GetPermintaanByProyekID(proyekID int) ([]PermintaanTable, error) {
	var permintaans []PermintaanTable

	// Fetch permintaans with details using Preload
	err := repo.db.Where("id_proyek = ?", proyekID).Preload("Details").Find(&permintaans).Error
	if err != nil {
		return nil, err
	}

	return permintaans, nil
}

func (repo *permintaanRepository) GetAllPermintaan() ([]PermintaanTable, error) {
	var permintaans []PermintaanTable
	err := repo.db.Preload("Details").Find(&permintaans).Error
	return permintaans, err
}
