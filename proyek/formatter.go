package proyek

import "time"

type ProyekFormatter struct {
	IdProyek         int       `json:"id"`
	NamaProyek string    `json:"nama_proyek"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FormatProyek(proyek ProyekTable) ProyekFormatter {
	return ProyekFormatter{
		IdProyek:   proyek.IdProyek,
		NamaProyek: proyek.NamaProyek,
		CreatedAt: proyek.CreatedAt,
		UpdatedAt: proyek.UpdatedAt,
	}

}

func FormatProyeks(proyeks []ProyekTable) []ProyekFormatter {
	var proyeksFormatter []ProyekFormatter

	for _, proyek := range proyeks {
		proyeksFormatter = append(proyeksFormatter, FormatProyek(proyek))
	}

	return proyeksFormatter
}
