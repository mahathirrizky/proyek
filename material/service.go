package material

type Service interface {
	CreateMaterial(input CreateMaterialInput) (MaterialTable, error)
	CreateStok(input CreateStokInput) (StokTable, error)
	GetAllMaterials() ([]MaterialFormatter, error)
	GetAllStok() ([]StokTable, error)
	UpdateMaterial(id int, input UpdateMaterialInput) (MaterialTable, error)
	GetMaterialByID(id int) (MaterialFormatter, error)
}

type service struct {
	repository Repository
}

func NewMaterialService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateMaterial(input CreateMaterialInput) (MaterialTable, error) {
	material := MaterialTable{
		NamaMaterial: input.NamaMaterial,
		Spesifikasi:  input.Spesifikasi,
	}

	newMaterial, err := s.repository.CreateMaterial(material)
	return newMaterial, err
}

func (s *service) CreateStok(input CreateStokInput) (StokTable, error) {
	stok := StokTable{
		IdProyek:   input.IDProyek,
		IdMaterial: input.IDMaterial,
		Jumlah:     input.Jumlah,
	}

	newStok, err := s.repository.CreateStok(stok)
	return newStok, err
}

func (s *service) GetAllMaterials() ([]MaterialFormatter, error) {
	materials, err := s.repository.FindAllMaterials()
	if err != nil {
		return nil, err
	}

	stokData, err := s.calculateMaterialStock(materials)
	if err != nil {
		return nil, err
	}

	formattedMaterials := FormatMaterials(materials, stokData)
	return formattedMaterials, nil
}

func (s *service) GetAllStok() ([]StokTable, error) {
	stok, err := s.repository.FindAllStok()
	return stok, err
}

func (s *service) calculateMaterialStock([]MaterialTable) (map[int]int, error) {
	stok, err := s.repository.FindAllStok()
	if err != nil {
		return nil, err
	}

	stokData := make(map[int]int)
	for _, s := range stok {
		stokData[s.IdMaterial] += s.Jumlah
	}

	return stokData, nil
}

func (s *service) UpdateMaterial(id int, input UpdateMaterialInput) (MaterialTable, error) {
	material, err := s.repository.FindByID(id)
	if err != nil {
		return MaterialTable{}, err
	}

	material.NamaMaterial = input.NamaMaterial
	material.Spesifikasi = input.Spesifikasi

	updatedMaterial, err := s.repository.UpdateMaterial(material)
	if err != nil {
		return updatedMaterial, err
	}

	// Update the stock based on the new quantity
	stok, err := s.repository.FindStokByMaterialID(material.IdMaterial)
	if err != nil {
		return updatedMaterial, err
	}

	stok.Jumlah = input.Jumlah
	_, err = s.repository.UpdateStok(stok)
	if err != nil {
		return updatedMaterial, err
	}

	return updatedMaterial, nil
}
func (s *service) GetMaterialByID(id int) (MaterialFormatter, error) {
	material, err := s.repository.FindByID(id)
	if err != nil {
		return MaterialFormatter{}, err
	}

	stok, err := s.repository.FindStokByMaterialID(material.IdMaterial)
	if err != nil {
		return MaterialFormatter{}, err
	}

	return FormatMaterial(material, stok.Jumlah), nil
}