package prodi

type Service interface {
	GetProdis() ([]Prodi, error)
	GetProdiByID(input GetProdiDetailInput) (Prodi, error)
	CreateProdi(input CreateProdiInput) (Prodi, error)
	UpdateProdi(inputID GetProdiDetailInput, input CreateProdiInput) (Prodi, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetProdis() ([]Prodi, error) {
	prodis, err := s.repository.FindAll()
	if err != nil {
		return prodis, err
	}
	return prodis, nil
}

func (s *service) GetProdiByID(input GetProdiDetailInput) (Prodi, error) {
	prodi, err := s.repository.FindByID(input.ID)
	if err != nil {
		return prodi, err
	}
	return prodi, nil
}

func (s *service) CreateProdi(input CreateProdiInput) (Prodi, error) {
	prodi := Prodi{}
	prodi.Nama = input.Nama

	newProdi, err := s.repository.Save(prodi)
	if err != nil {
		return newProdi, err
	}
	return newProdi, nil
}

func (s *service) UpdateProdi(inputID GetProdiDetailInput, inputData CreateProdiInput) (Prodi, error) {
	prodi, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return prodi, err
	}
	prodi.Name = inputData.Name

	updateProdi, err := s.repository.Save(prodi)
	if err != nil {
		return updateProdi, err
	}
	return updateProdi, nil

}
