// file: adminproyek/input.go

package adminproyek

type CreateAdminProyekInput struct {
    IDUser   int `json:"id_user" binding:"required"`
    IDProyek int `json:"id_proyek" binding:"required"`
}

type UpdateAdminProyekInput struct {
    IDUser   int `json:"id_user"`
    IDProyek int `json:"id_proyek"`
}
