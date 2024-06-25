package material

type MaterialFormatter struct {
	ID           int    `json:"id"`
	NamaMaterial string `json:"nama_material"`
	Spesifikasi  string `json:"spesifikasi"`
	JumlahStok   int    `json:"jumlah_stok"`
}

func FormatMaterial(material MaterialTable, jumlahStok int) MaterialFormatter {
	return MaterialFormatter{
		ID:           material.IdMaterial,
		NamaMaterial: material.NamaMaterial,
		Spesifikasi:  material.Spesifikasi,
		JumlahStok:   jumlahStok,
	}
}

type StokFormatter struct {
	ID         int `json:"id"`
	IDProyek   int `json:"id_proyek"`
	IDMaterial int `json:"id_material"`
	Jumlah     int `json:"jumlah"`
}

func FormatStok(stok StokTable) StokFormatter {
	return StokFormatter{
		ID:         stok.IdStok,
		IDProyek:   stok.IdProyek,
		IDMaterial: stok.IdMaterial,
		Jumlah:     stok.Jumlah,
	}
}

func FormatMaterials(materials []MaterialTable, stokData map[int]int) []MaterialFormatter {
	var formattedMaterials []MaterialFormatter
	for _, material := range materials {
		jumlahStok := stokData[material.IdMaterial]
		formattedMaterial := FormatMaterial(material, jumlahStok)
		formattedMaterials = append(formattedMaterials, formattedMaterial)
	}
	return formattedMaterials
}
