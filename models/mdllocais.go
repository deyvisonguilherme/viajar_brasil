package models

// mdllocal representa um local buscado pelo usuário.
type mdllocal struct {
	Codigo    int    `json:"codigo"`
	Categoria string `json:"categoria"`
	Regiao    string `json:"regiao"`
	Foto      string `json:"foto"`
	Cidade    string `json:"cidade"`
}

// ListLocais função responsável por buscar locais no banco de dados.
func ListLocais(vregiao int, vcategoria int) ([]*mdllocal, error) {
	linhas, err := db.Query("select codigo, categoria, regiao, foto, cidade from spt_select_locais($1,$2)", vregiao, vcategoria)
	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	locais := make([]*mdllocal, 0)

	for linhas.Next() {
		objLocais := new(mdllocal)
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
