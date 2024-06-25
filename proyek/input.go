package proyek

type CreateProyekInput struct {
    NamaProyek string `json:"nama_proyek" binding:"required"`
}

type UpdateProyekInput struct {
    NamaProyek string `json:"nama_proyek" binding:"required"`
}

type GetProyekDetailInput struct {
	ID int `uri:"id" binding:"required"`
}