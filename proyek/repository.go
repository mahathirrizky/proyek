package proyek

import "gorm.io/gorm"

type Repository interface {
    FindAll() ([]ProyekTable, error)
    FindByID(id int) (ProyekTable, error)
    Save(proyek ProyekTable) (ProyekTable, error)
    Update(proyek ProyekTable) (ProyekTable, error)
    Delete(proyek ProyekTable) error
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
    return &repository{db}
}

func (r *repository) FindAll() ([]ProyekTable, error) {
    var proyeks []ProyekTable
    err := r.db.Find(&proyeks).Error
    return proyeks, err
}

func (r *repository) FindByID(id int) (ProyekTable, error) {
    var proyek ProyekTable
    err := r.db.First(&proyek, id).Error
    return proyek, err
}

func (r *repository) Save(proyek ProyekTable) (ProyekTable, error) {
    err := r.db.Create(&proyek).Error
    return proyek, err
}

func (r *repository) Update(proyek ProyekTable) (ProyekTable, error) {
    err := r.db.Save(&proyek).Error
    return proyek, err
}

func (r *repository) Delete(proyek ProyekTable) error {
    err := r.db.Delete(&proyek).Error
    return err
}
