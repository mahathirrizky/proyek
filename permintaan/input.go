package permintaan

type CreatePermintaanInput struct {
	IdProyek int                           `json:"id_proyek" binding:"required"`
	Status   string                        `json:"status"`
	Details  []CreatePermintaanDetailInput `json:"details" binding:"required,dive"`
}

type CreatePermintaanDetailInput struct {
	IdMaterial int `json:"id_material" binding:"required"`
	Jumlah     int `json:"jumlah" binding:"required"`
}

type UpdatePermintaanInput struct {
	Status  string                        `json:"status"`
	Details []UpdatePermintaanDetailInput `json:"details" binding:"required,dive"`
}

type UpdatePermintaanDetailInput struct {
	IdMaterial int `json:"id_material" binding:"required"`
	Jumlah     int `json:"jumlah" binding:"required"`
}
