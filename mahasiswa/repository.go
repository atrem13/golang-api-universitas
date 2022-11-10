package mahasiswa

import "gorm.io/gorm"

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
	err := r.db.Preload("Prodi").Create(&mahasiswa).Error
	if err != nil {
		return mahasiswa, err
	}
	return mahasiswa, nil
}

func (r *repository) Update(mahasiswa Mahasiswa) (Mahasiswa, error) {
	// err := r.db.Save(&mahasiswa).Error
	err := r.db.Model(&mahasiswa).Where("id = ?", mahasiswa.ID).Updates(map[string]interface{}{"nim": mahasiswa.Nim, "nama": mahasiswa.Nama, "prodi_id": 3}).Error
	// err := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&mahasiswa).Error
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
