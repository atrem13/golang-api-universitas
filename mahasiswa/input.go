package mahasiswa

type GetMahasiswaDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateMahasiswaInput struct {
	Nim     string `json:"nim" `
	Nama    string `json:"nama" `
	ProdiID int    `json:"prodi_id" `
}
