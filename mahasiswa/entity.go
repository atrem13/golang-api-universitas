package mahasiswa

import (
	"time"

	"github.com/atrem13/golang-api-universitas/prodi"
)

type Mahasiswa struct {
	ID        int
	Nim       string
	Nama      string
	ProdiID   int `gorm:"foreignKey:ID"`
	Prodi     prodi.Prodi
	CreatedAt time.Time
	UpdatedAt time.Time
}
