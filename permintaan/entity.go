package permintaan

import (
	"time"
)

type PermintaanTable struct {
	IdPermintaan int       `gorm:"primaryKey;autoIncrement"`
	IdProyek     int       
	Status       string    
	CreatedAt    time.Time 
	Details      []PermintaanDetailTable `gorm:"foreignKey:IdPermintaan"`
}

type PermintaanDetailTable struct {
	IdDetail     int      `gorm:"primaryKey;autoIncrement"`
	IdPermintaan int       
	IdMaterial   int       
	Jumlah       int       
}
