package proyek

import "time"

type ProyekTable struct {
	IdProyek         int `gorm:"primaryKey;autoIncrement"`
	NamaProyek string
	CreatedAt  time.Time 
	UpdatedAt  time.Time 
}
