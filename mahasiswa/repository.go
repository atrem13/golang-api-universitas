package mahasiswa

import "github.com/jinzhu/gorm"

type Repository interface {
	FindAll() ([]Mahasiswa, error)
	FindByID(ID int) (Mahasiswa, error)
	Save(mahasisw Mahasiswa) (Mahasiswa, error)
	Update(mahasiswa Mahasiswa) (Mahasiswa, error)
	Delete(mahasiswa Mahasiswa) (Mahasiswa, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Mahasiswa, error) {
	var mahasiswas []Mahasiswa
	err := r.db.Preload("Prodi").Find(&mahasiswas).Error
	if err != nil {
		return mahasiswas, err
	}
	return mahasiswas, nil
}

func (r *repository) FindByID(ID int) (Mahasiswa, error) {
	var mahasiswa Mahasiswa
	err := r.db.Where("id = ?", ID).Preload("Prodi").Find(&mahasiswa).Error
	if err != nil {
		return mahasiswa, err
	}
	return mahasiswa, nil
}

func (r *repository) Save(mahasiswa Mahasiswa) (Mahasiswa, error) {
	err := r.db.Create(&mahasiswa).Error
	if err != nil {
		return mahasiswa, err
	}
	return mahasiswa, nil
}

func (r *repository) Update(mahasiswa Mahasiswa) (Mahasiswa, error) {
	err := r.db.Save(&mahasiswa).Error
	if err != nil {
		return mahasiswa, err
	}
	return mahasiswa, nil
}

func (r *repository) Delete(mahasiswa Mahasiswa) (Mahasiswa, error) {
	err := r.db.Delete(&mahasiswa).Error
	if err != nil {
		return mahasiswa, err
	}
	return mahasiswa, nil
}
