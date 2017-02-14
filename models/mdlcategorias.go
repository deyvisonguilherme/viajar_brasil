package models

// Categoria representa o organograma de categoria.
type Categoria struct {
	ID           int
	DSCcategoria string
	Ativo        bool
}

// GetCategorias  função que pega todas as categorias cadastradas.
func GetCategorias() ([]*Categoria, error) {
	linhas, err := db.Query("select id, ds_categoria, ativo from categorias where ativo=True")

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	categorias := make([]*Categoria, 0)
	for linhas.Next() {
		objCategoria := new(Categoria)

		err := linhas.Scan(&objCategoria.ID, &objCategoria.DSCcategoria, &objCategoria.Ativo)

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
