package pembelian

import "time"

type PembelianFormatter struct {
    IdPembelian int       `json:"id_pembelian"`
    IdMaterial  int       `json:"id_material"`
    IdProyek  int       `json:"id_proyek"`
    Jumlah      int       `json:"jumlah"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

func FormatPembelian(pembelian PembelianTable) PembelianFormatter {
    return PembelianFormatter{
        IdPembelian: pembelian.IdPembelian,
        IdMaterial:  pembelian.IdMaterial,
        IdProyek:    pembelian.IdProyek,
        Jumlah:      pembelian.Jumlah,
        Status:      pembelian.Status,
        CreatedAt:   pembelian.CreatedAt,
        UpdatedAt:   pembelian.UpdatedAt,
    }
}

func FormatPembelianList(pembelianlist []PembelianTable) []PembelianFormatter {
    pembelianFormatters := make([]PembelianFormatter, len(pembelianlist))
    for i, p := range pembelianlist {
        pembelianFormatters[i] = FormatPembelian(p)
    }
    return pembelianFormatters
}
