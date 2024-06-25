package pembelian

type CreatePembelianInput struct {
	IdProyek   int    `json:"id_proyek" binding:"required"`
	IdMaterial int    `json:"id_material" binding:"required"`
	Jumlah     int    `json:"jumlah" binding:"required"`
	Status     string `json:"status"`
}

type UpdatePembelianInput struct {
	Jumlah int    `json:"jumlah" binding:"required"`
	Status string `json:"status"`
}
type GetPembelianDetailInput struct {
	ID int `uri:"id" binding:"required"`
}