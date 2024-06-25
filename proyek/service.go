package proyek

type Service interface {
    GetAllProyeks() ([]ProyekTable, error)
    GetProyekByID(input GetProyekDetailInput) (ProyekTable, error)
    CreateProyek(input CreateProyekInput) (ProyekTable, error)
    UpdateProyek(inputID GetProyekDetailInput, inputData UpdateProyekInput) (ProyekTable, error)
    DeleteProyek(inputID GetProyekDetailInput) error
}

type service struct {
    repository Repository
}

func NewService(repository Repository) *service {
    return &service{repository}
}

func (s *service) GetAllProyeks() ([]ProyekTable, error) {
    return s.repository.FindAll()
}

func (s *service) GetProyekByID(input GetProyekDetailInput) (ProyekTable, error) {
    return s.repository.FindByID(input.ID)
}

func (s *service) CreateProyek(input CreateProyekInput) (ProyekTable, error) {
    proyek := ProyekTable{
        NamaProyek: input.NamaProyek,
    }
    return s.repository.Save(proyek)
}

func (s *service) UpdateProyek(inputID GetProyekDetailInput, inputData UpdateProyekInput) (ProyekTable, error) {
    proyek, err := s.repository.FindByID(inputID.ID)
    if err != nil {
        return proyek, err
    }

    proyek.NamaProyek = inputData.NamaProyek
		

    return s.repository.Update(proyek)
}

func (s *service) DeleteProyek(inputID GetProyekDetailInput) error {
    proyek, err := s.repository.FindByID(inputID.ID)
    if err != nil {
        return err
    }

    return s.repository.Delete(proyek)
}
