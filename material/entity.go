package material

import "time"

type MaterialTable struct {
	IdMaterial           int `gorm:"primaryKey;autoIncrement"`
	NamaMaterial string
	Spesifikasi  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type StokTable struct {
	IdStok         int `gorm:"primaryKey;autoIncrement"`
	IdProyek   int
	IdMaterial int
	Jumlah     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
