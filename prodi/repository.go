package prodi

import "github.com/jinzhu/gorm"

type Repository interface {
	FindAll() ([]Prodi, error)
	FindByID(ID int) (Prodi, error)
	Save(prodi Prodi) (Prodi, error)
	Update(prodi Prodi) (Prodi, error)
	Delete(prodi Prodi) (Prodi, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Prodi, error) {
	var prodis []Prodi
	err := r.db.Find(&prodis).Error
	if err != nil {
		return prodis, err
	}
	return prodis, nil
}

func (r *repository) FindByID(ID int) (Prodi, error) {
	var prodi Prodi
	err := r.db.Where("id = ?", ID).Find(&prodi).Error
	if err != nil {
		return prodi, err
	}
	return prodi, nil
}

func (r *repository) Save(prodi Prodi) (Prodi, error) {
	err := r.db.Create(&prodi).Error
	if err != nil {
		return prodi, err
	}
	return prodi, nil
}

func (r *repository) Update(prodi Prodi) (Prodi, error) {
	err := r.db.Save(&prodi).Error
	if err != nil {
		return prodi, err
	}
	return prodi, nil
}

func (r *repository) Delete(prodi Prodi) (Prodi, error) {
	err := r.db.Delete(&prodi).Error
	if err != nil {
		return prodi, err
	}
	return prodi, nil
}
