// file: adminproyek/repository.go

package adminproyek

import (
    "gorm.io/gorm"
)

// AdminProyekRepository adalah interface untuk operasi-operasi terhadap entitas AdminProyekTable.
type AdminProyekRepository interface {
    CreateAdminProyek(adminProyek AdminProyekTable) (*AdminProyekTable, error)
    UpdateAdminProyek(adminProyek AdminProyekTable) (*AdminProyekTable, error)
    DeleteAdminProyek(adminProyek AdminProyekTable) error
    GetAdminProyekByID(id int) (*AdminProyekTable, error)
    GetAllAdminProyeks() ([]AdminProyekTable, error)
}

// adminProyekRepository adalah implementasi dari AdminProyekRepository menggunakan GORM.
type adminProyekRepository struct {
    db *gorm.DB
}

// NewAdminProyekRepository membuat instance baru dari adminProyekRepository.
func NewAdminProyekRepository(db *gorm.DB) AdminProyekRepository {
    return &adminProyekRepository{db}
}

func (repo *adminProyekRepository) CreateAdminProyek(adminProyek AdminProyekTable) (*AdminProyekTable, error) {
    if err := repo.db.Create(&adminProyek).Error; err != nil {
        return nil, err
    }
    return &adminProyek, nil
}

func (repo *adminProyekRepository) UpdateAdminProyek(adminProyek AdminProyekTable) (*AdminProyekTable, error) {
    if err := repo.db.Save(&adminProyek).Error; err != nil {
        return nil, err
    }
    return &adminProyek, nil
}

func (repo *adminProyekRepository) DeleteAdminProyek(adminProyek AdminProyekTable) error {
    if err := repo.db.Delete(&adminProyek).Error; err != nil {
        return err
    }
    return nil
}

func (repo *adminProyekRepository) GetAdminProyekByID(id int) (*AdminProyekTable, error) {
    var adminProyek AdminProyekTable
    if err := repo.db.First(&adminProyek, id).Error; err != nil {
        return nil, err
    }
    return &adminProyek, nil
}

func (repo *adminProyekRepository) GetAllAdminProyeks() ([]AdminProyekTable, error) {
    var adminProyeks []AdminProyekTable
    if err := repo.db.Find(&adminProyeks).Error; err != nil {
        return nil, err
    }
    return adminProyeks, nil
}
