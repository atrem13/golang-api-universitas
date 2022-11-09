package mahasiswa

type GetMahasiswaDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateMahasiswaInput struct {
	Nim     string `json:"nim" binding:"required"`
	Nama    string `json:"nama" binding:"required"`
	ProdiID int    `json:"prodi_id" binding:"required"`
}
