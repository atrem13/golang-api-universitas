package prodi

type GetProdiDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateProdiInput struct {
	Nama string `json:"name" binding:"required"`
}
