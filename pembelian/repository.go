package pembelian

import (
	"gorm.io/gorm"
	"proyek/material"
)

type PembelianRepository interface {
	CreatePembelian(data *PembelianTable) error
	UpdatePembelian(id GetPembelianDetailInput, data *PembelianTable) error
	DeletePembelian(id GetPembelianDetailInput) error
	GetPembelianByID(id GetPembelianDetailInput) (*PembelianTable, error)
	GetPembelianByProyekID(proyekID int) ([]PembelianTable, error)
	GetAllPembelian() ([]PembelianTable, error)
	UpdateStock(materialID int, jumlah int) error
}

type pembelianRepository struct {
	db *gorm.DB
}

func NewPembelianRepository(db *gorm.DB) PembelianRepository {
	return &pembelianRepository{db}
}

func (repo *pembelianRepository) CreatePembelian(data *PembelianTable) error {
	if err := repo.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (repo *pembelianRepository) UpdatePembelian(id GetPembelianDetailInput, data *PembelianTable) error {
	if err := repo.db.Model(&PembelianTable{}).Where("id_pembelian = ?", id.ID).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (repo *pembelianRepository) DeletePembelian(id GetPembelianDetailInput) error {
	if err := repo.db.Where("id_pembelian = ?", id.ID).Delete(&PembelianTable{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *pembelianRepository) GetPembelianByID(id GetPembelianDetailInput) (*PembelianTable, error) {
	var pembelian PembelianTable
	if err := repo.db.Where("id_pembelian = ?", id.ID).First(&pembelian).Error; err != nil {
		return nil, err
	}
	return &pembelian, nil
}

func (repo *pembelianRepository) GetPembelianByProyekID(proyekID int) ([]PembelianTable, error) {
	var pembelians []PembelianTable
	if err := repo.db.Where("id_proyek = ?", proyekID).Find(&pembelians).Error; err != nil {
		return nil, err
	}
	return pembelians, nil
}

func (repo *pembelianRepository) GetAllPembelian() ([]PembelianTable, error) {
	var pembelian []PembelianTable
	if err := repo.db.Find(&pembelian).Error; err != nil {
		return nil, err
	}
	return pembelian, nil
}

func (repo *pembelianRepository) UpdateStock(materialID int, jumlah int) error {
	var stok material.StokTable
	if err := repo.db.Where("id_material = ?", materialID).First(&stok).Error; err != nil {
		return err
	}

	stok.Jumlah += jumlah
	if err := repo.db.Save(&stok).Error; err != nil {
		return err
	}

	return nil
}
