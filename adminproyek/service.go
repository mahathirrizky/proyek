// file: adminproyek/service.go

package adminproyek



// AdminProyekService adalah interface untuk operasi-operasi bisnis terhadap entitas AdminProyekTable.
type AdminProyekService interface {
    CreateAdminProyek(input CreateAdminProyekInput) (*AdminProyekTable, error)
    UpdateAdminProyek(id int, input UpdateAdminProyekInput) (*AdminProyekTable, error)
    DeleteAdminProyek(id int) error
    GetAdminProyekByID(id int) (*AdminProyekTable, error)
    GetAllAdminProyeks() ([]AdminProyekTable, error)
}

type adminProyekService struct {
    repo AdminProyekRepository
}

// NewAdminProyekService membuat instance baru dari adminProyekService.
func NewAdminProyekService(repo AdminProyekRepository) AdminProyekService {
    return &adminProyekService{repo}
}

func (s *adminProyekService) CreateAdminProyek(input CreateAdminProyekInput) (*AdminProyekTable, error) {
    adminProyek := AdminProyekTable{
        IdUser:   input.IDUser,
        IdProyek: input.IDProyek,
    }
    return s.repo.CreateAdminProyek(adminProyek)
}

func (s *adminProyekService) UpdateAdminProyek(id int, input UpdateAdminProyekInput) (*AdminProyekTable, error) {
    adminProyek, err := s.repo.GetAdminProyekByID(id)
    if err != nil {
        return nil, err
    }
    
    adminProyek.IdUser = input.IDUser
    adminProyek.IdProyek = input.IDProyek
    
    return s.repo.UpdateAdminProyek(*adminProyek)
}

func (s *adminProyekService) DeleteAdminProyek(id int) error {
    adminProyek, err := s.repo.GetAdminProyekByID(id)
    if err != nil {
        return err
    }
    
    return s.repo.DeleteAdminProyek(*adminProyek)
}

func (s *adminProyekService) GetAdminProyekByID(id int) (*AdminProyekTable, error) {
    return s.repo.GetAdminProyekByID(id)
}

func (s *adminProyekService) GetAllAdminProyeks() ([]AdminProyekTable, error) {
    return s.repo.GetAllAdminProyeks()
}
