package prodi

type ProdiFormatter struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

func FormatProdi(prodi Prodi) ProdiFormatter {
	prodiFormatter := ProdiFormatter{}
	prodiFormatter.Nama = prodi.Nama

	return prodiFormatter
}

func FormatProdis(prodis []Prodi) []ProdiFormatter {
	prodisFormatter := []ProdiFormatter{}
	for _, prodi := range prodis {
		prodiFormatter := FormatProdi(prodi)
		prodisFormatter = append(prodisFormatter, prodiFormatter)
	}
	return prodisFormatter
}

// type ProdiDetailFormatter struct {
// 	ID   int    `json:"id"`
// 	Nama string `json:"nama"`
// }

// func FormatProdiDetail(prodi Prodi) ProdiDetailFormatter {
// 	prodiDetailFormatter := ProdiDetailFormatter{}
// 	prodiDetailFormatter.Nama = prodi.Nama

// 	return prodiDetailFormatter
// }
