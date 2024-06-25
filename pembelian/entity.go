package pembelian

import "time"

type PembelianTable struct {
    IdPembelian       int `gorm:"primaryKey;autoIncrement"`
    IdMaterial int
    Jumlah     int   
    Status     string 
    CreatedAt  time.Time 
    UpdatedAt  time.Time
}
