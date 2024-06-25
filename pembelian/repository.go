package pembelian

import (
	"gorm.io/gorm"
	"proyek/material"
)

type PembelianRepository interface {
	CreatePembelian(data *PembelianTable) error
	UpdatePembelian(id int, data *PembelianTable) error
	DeletePembelian(id int) error
	GetPembelianByID(id int) (*PembelianTable, error)
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

func (repo *pembelianRepository) UpdatePembelian(id int, data *PembelianTable) error {
	if err := repo.db.Model(&PembelianTable{}).Where("id_pembelian = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (repo *pembelianRepository) DeletePembelian(id int) error {
	if err := repo.db.Where("id_pembelian = ?", id).Delete(&PembelianTable{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo *pembelianRepository) GetPembelianByID(id int) (*PembelianTable, error) {
	var pembelian PembelianTable
	if err := repo.db.Where("id_pembelian = ?", id).First(&pembelian).Error; err != nil {
		return nil, err
	}
	return &pembelian, nil
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
