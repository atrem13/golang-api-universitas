package mahasiswa

type MahasiswaFormatter struct {
	ID      int    `json:"id`
	Nim     string `json:"nim"`
	Nama    string `json:"nama"`
	ProdiID int    `json:"prodi_id"`
}

func FormatMahasiswa(mahasiswa Mahasiswa) MahasiswaFormatter {
	mahasiswaFormatter := MahasiswaFormatter{}
	mahasiswaFormatter.ID = mahasiswa.ID
	mahasiswaFormatter.Nim = mahasiswa.Nim
	mahasiswaFormatter.Nama = mahasiswa.Nama
	mahasiswaFormatter.ProdiID = mahasiswa.ProdiID

	return mahasiswaFormatter
}

func FormatMahasiswas(mahasiswa []Mahasiswa) []MahasiswaFormatter {
	mahasiswasFormatter := []MahasiswaFormatter{}
	for _, mahasiswa := range mahasiswa {
		mahasiswaFormatter := FormatMahasiswa(mahasiswa)
		mahasiswasFormatter = append(mahasiswasFormatter, mahasiswaFormatter)
	}
	return mahasiswasFormatter
}
