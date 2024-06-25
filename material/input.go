package material

type CreateMaterialInput struct {
	IDProyek   int `json:"id_proyek" binding:"required"`
	NamaMaterial string `json:"nama_material" binding:"required"`
	Spesifikasi  string `json:"spesifikasi" binding:"required"`
	Jumlah     int `json:"jumlah" binding:"required"`
}

type CreateStokInput struct {
	IDProyek   int 
	IDMaterial int 
	Jumlah     int 
}

type UpdateMaterialInput struct {
	IDProyek     int    `json:"id_proyek" binding:"required"`
	NamaMaterial string `json:"nama_material" binding:"required"`
	Spesifikasi  string `json:"spesifikasi" binding:"required"`
	Jumlah       int    `json:"jumlah" binding:"required"`
}
