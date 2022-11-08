package prodi

type Service interface {
	GetProdiByID(input GetProdiDetailInput) (Prodi, error)
	CreateProdi(input CreateProdiInput) (Prodi, error)
	UpdateProdi(inputID GetProdiDetailInput, input CreateProdiInput) (Prodi, error)
}

type srevice struct {
	repository Repository
}
