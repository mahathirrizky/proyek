package permintaan



type PermintaanFormatter struct {
	ID        int                         `json:"id"`
	IdProyek  int                         `json:"id_proyek"`
	Status    string                      `json:"status"`
	CreatedAt string                      `json:"created_at"`
	Details   []PermintaanDetailFormatter `json:"details"`
}

type PermintaanDetailFormatter struct {
	ID         int `json:"id"`
	IdMaterial int `json:"id_material"`
	Jumlah     int `json:"jumlah"`
}

func FormatPermintaan(permintaan *PermintaanTable) PermintaanFormatter {
	details := []PermintaanDetailFormatter{}
	for _, detail := range permintaan.Details {
		detailFormatter := PermintaanDetailFormatter{
			ID:         detail.IdDetail,
			IdMaterial: detail.IdMaterial,
			Jumlah:     detail.Jumlah,
		}
		details = append(details, detailFormatter)
	}

	return PermintaanFormatter{
		ID:        permintaan.IdPermintaan,
		IdProyek:  permintaan.IdProyek,
		Status:    permintaan.Status,
		CreatedAt: permintaan.CreatedAt.Format("2006-01-02 15:04:05"),
		Details:   details,
	}
}

func FormatPermintaanList(permintaanList []PermintaanTable) []PermintaanFormatter {
	formattedList := []PermintaanFormatter{}
	for _, permintaan := range permintaanList {
		formattedPermintaan := FormatPermintaan(&permintaan)
		formattedList = append(formattedList, formattedPermintaan)
	}
	return formattedList
}
