package main

import "gorm.io/gorm"

// Proyek
type Proyek struct {
    gorm.Model
    NamaProyek string `json:"nama_proyek"`
}

// Material
type Material struct {
    gorm.Model
    NamaMaterial string `json:"nama_material"`
    Spesifikasi  string `json:"spesifikasi"`
}

// Stok
type Stok struct {
    gorm.Model
    IDProyek   uint `json:"id_proyek"`
    IDMaterial uint `json:"id_material"`
    Jumlah     int  `json:"jumlah"`
}

// User
type User struct {
    gorm.Model
    Nama     string `json:"nama"`
    Username string `json:"username"`
    Password string `json:"password"`
    Role     string `json:"role"`
}

// AdminProyek
type AdminProyek struct {
    gorm.Model
    IDUser   uint `json:"id_user"`
    IDProyek uint `json:"id_proyek"`
}

// Pembelian
type Pembelian struct {
    gorm.Model
    IDMaterial uint   `json:"id_material"`
    Jumlah     int    `json:"jumlah"`
    Status     string `json:"status"`
}
