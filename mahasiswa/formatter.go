package mahasiswa

type MahasiswaFormatter struct {
	ID      int                     `json:"id`
	Nim     string                  `json:"nim"`
	Nama    string                  `json:"nama"`
	ProdiID int                     `json:"prodi_id"`
	Prodi   MahasiswaProdiFormatter `json:"prodi"`
}
type MahasiswaProdiFormatter struct {
	Nama string `json:"nama"`
}

func FormatMahasiswa(mahasiswa Mahasiswa) MahasiswaFormatter {
	mahasiswaFormatter := MahasiswaFormatter{}
	mahasiswaFormatter.ID = mahasiswa.ID
	mahasiswaFormatter.Nim = mahasiswa.Nim
	mahasiswaFormatter.Nama = mahasiswa.Nama
	mahasiswaFormatter.ProdiID = mahasiswa.ProdiID

	prodi := mahasiswa.Prodi
	mahasiswasProdiFormatter := MahasiswaProdiFormatter{}
	mahasiswasProdiFormatter.Nama = prodi.Nama

	mahasiswaFormatter.Prodi = mahasiswasProdiFormatter

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
