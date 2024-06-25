package adminproyek

import "time"

type AdminProyekTable struct {
    IdAdminProyek int `gorm:"primaryKey;autoIncrement"`
    IdUser        int
    IdProyek      int
    CreatedAt     time.Time
    UpdatedAt     time.Time
}
