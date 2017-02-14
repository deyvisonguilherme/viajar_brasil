package models

// mdlcateogira representa o organograma de categoria.
type mdlcateogira struct {
	Codigo    int    `json:"id"`
	Categoria string `json:"categoria"`
}

// ListCategoria retona uma lista de categorias.
func ListCategoria() ([]*mdlcateogira, error) {
	linhas, err := db.Query("select codigo, categoria from vwCategorias")
	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	categorias := make([]*mdlcateogira, 0)

	for linhas.Next() {
		objCategoria := new(mdlcateogira)

		err := linhas.Scan(&objCategoria.Codigo, &objCategoria.Categoria)

		if err != nil {
			return nil, err
		}

		categorias = append(categorias, objCategoria)
	}
	if err = linhas.Err(); err != nil {
		return nil, err
	}

	return categorias, nil
}
