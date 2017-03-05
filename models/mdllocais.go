package models

// MDLLocal :  representa um local buscado pelo usuário.
type MDLLocal struct {
	Codigo    int    `json:"codigo"`
	Categoria string `json:"categoria"`
	Regiao    string `json:"regiao"`
	Foto      string `json:"foto"`
	Cidade    string `json:"cidade"`
}

// Locais : função responsável por buscar locais no banco de dados.
func Locais(vregiao int, vcategoria int) ([]*MDLLocal, error) {
	linhas, err := db.Query("select codigo, categoria, regiao, foto, cidade from spt_select_locais($1,$2)", vregiao, vcategoria)
	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	locais := make([]*MDLLocal, 0)

	for linhas.Next() {
		objLocais := new(MDLLocal)
		err := linhas.Scan(&objLocais.Codigo, &objLocais.Categoria, &objLocais.Regiao, &objLocais.Foto, &objLocais.Cidade)

		if err != nil {
			return nil, err
		}

		locais = append(locais, objLocais)
	}
	if err = linhas.Err(); err != nil {
		return nil, err
	}

	return locais, nil
}
