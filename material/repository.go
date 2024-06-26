package material

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindByID(id int) (MaterialTable, error)
	CreateMaterial(material MaterialTable) (MaterialTable, error)
	UpdateMaterial(material MaterialTable) (MaterialTable, error)
	FindAllMaterials() ([]MaterialTable, error)
	FindAllStok() ([]StokTable, error)
	CreateStok(stok StokTable) (StokTable, error)
	FindStokByMaterialID(materialID int) (StokTable, error)
	UpdateStok(stok StokTable) (StokTable, error)
}

type repository struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(id int) (MaterialTable, error) {
	var material MaterialTable
	err := r.db.Where("id_material = ?", id).First(&material).Error
	return material, err
}

func (r *repository) CreateMaterial(material MaterialTable) (MaterialTable, error) {
	err := r.db.Create(&material).Error
	return material, err
}

func (r *repository) UpdateMaterial(material MaterialTable) (MaterialTable, error) {
	err := r.db.Save(&material).Error
	return material, err
}

func (r *repository) CreateStok(stok StokTable) (StokTable, error) {
	err := r.db.Create(&stok).Error
	return stok, err
}

func (r *repository) FindStokByMaterialID(materialID int) (StokTable, error) {
	var stok StokTable
	err := r.db.Where("id_material = ?", materialID).First(&stok).Error
	return stok, err
}

func (r *repository) UpdateStok(stok StokTable) (StokTable, error) {
	var existingStok StokTable
	if err := r.db.Where("id_material = ?", stok.IdMaterial).First(&existingStok).Error; err != nil {
		return existingStok, err
	}

	existingStok.Jumlah += stok.Jumlah
	err := r.db.Save(&existingStok).Error
	return existingStok, err
}

func (r *repository) FindAllMaterials() ([]MaterialTable, error) {
	var materials []MaterialTable
	err := r.db.Find(&materials).Error
	return materials, err
}

func (r *repository) FindAllStok() ([]StokTable, error) {
	var stok []StokTable
	err := r.db.Find(&stok).Error
	return stok, err
}
