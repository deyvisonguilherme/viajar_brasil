package models

// Local representa um local buscado pelo usuário.
type Local struct {
	Foto      string
	Regiao    string
	Categoria string
	Cidade    string
}

// GetLocais função responsável por buscar locais no banco de dados.
func GetLocais(vregiao int, vcategoria int) ([]*Local, error) {
	linhas, err := db.Query("select * from spt_select_locais($1,$2,True)", vregiao, vcategoria)

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	locais := make([]*Local, 0)

	for linhas.Next() {
		objLocais := new(Local)
		err := linhas.Scan(&objLocais.Foto, &objLocais.Regiao, &objLocais.Categoria, &objLocais.Cidade)

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
