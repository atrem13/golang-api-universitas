package mahasiswa

type Service interface {
	GetMahasiswas() ([]Mahasiswa, error)
	GetMahasiswaByID(input GetMahasiswaDetailInput) (Mahasiswa, error)
	CreateMahasiswa(input CreateMahasiswaInput) (Mahasiswa, error)
	UpdateMahasiswa(inputID GetMahasiswaDetailInput, inputData CreateMahasiswaInput) (Mahasiswa, error)
	DeleteMahasiswa(inputID GetMahasiswaDetailInput) (Mahasiswa, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetMahasiswas() ([]Mahasiswa, error) {
	mahasiswas, err := s.repository.FindAll()
	if err != nil {
		return mahasiswas, err
	}
	return mahasiswas, nil
}

func (s *service) GetMahasiswaByID(input GetMahasiswaDetailInput) (Mahasiswa, error) {
	mahasiswa, err := s.repository.FindByID(input.ID)
	if err != nil {
		return mahasiswa, err
	}
	return mahasiswa, nil
}

func (s *service) CreateMahasiswa(input CreateMahasiswaInput) (Mahasiswa, error) {
	mahasiswa := Mahasiswa{}
	mahasiswa.Nim = input.Nim
	mahasiswa.Nama = input.Nama
	mahasiswa.ProdiID = input.ProdiID

	newMahasiswa, err := s.repository.Save(mahasiswa)
	if err != nil {
		return newMahasiswa, err
	}
	return newMahasiswa, nil
}

func (s *service) UpdateMahasiswa(inputID GetMahasiswaDetailInput, inputData CreateMahasiswaInput) (Mahasiswa, error) {
	mahasiswa, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return mahasiswa, err
	}

	mahasiswa.Nim = inputData.Nim
	mahasiswa.Nama = inputData.Nama
	mahasiswa.ProdiID = inputData.ProdiID

	updateMahasiswa, err := s.repository.Update(mahasiswa)
	if err != nil {
		return updateMahasiswa, err
	}
	return updateMahasiswa, nil
}

func (s *service) DeleteMahasiswa(inputID GetMahasiswaDetailInput, inputData CreateMahasiswaInput) (Mahasiswa, error) {
	mahasiswa, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return mahasiswa, err
	}

	deleteMahasiswa, err := s.repository.Delete(mahasiswa)
	if err != nil {
		return deleteMahasiswa, err
	}
	return deleteMahasiswa, nil
}
