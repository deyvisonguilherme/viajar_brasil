package models

// MDLPesquisa : representa uma estrurua de retorno de pesquisa.
type MDLPesquisa struct {
	MDLCategoria []*MDLCategoria
	MDLRegiao    []*MDLRegiao
}

// MDLCategoria : representa uma estrutura de categoria.
type MDLCategoria struct {
	Codigo    int    `json:"id"`
	Categoria string `json:"categoria"`
}

// MDLRegiao :  representa uma estrutura de regiao.
type MDLRegiao struct {
	Codigo  int    `json:"codigo"`
	Regiao  string `json:"regiao"`
	Estados *string
}

// lstCategoria : retorna uma lista de categorias
func lstCategoria() ([]*MDLCategoria, error) {
	linhas, err := db.Query("select codigo, categoria from vwCategorias")
	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	categorias := make([]*MDLCategoria, 0)

	for linhas.Next() {
		objCategoria := new(MDLCategoria)

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

// lstRegiao : return função para busca de regioes.
func lstRegiao() ([]*MDLRegiao, error) {
	linhas, err := db.Query("select codigo, regiao, estados from vwRegioes")

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	regioes := make([]*MDLRegiao, 0)

	for linhas.Next() {
		objRegioes := new(MDLRegiao)
		err := linhas.Scan(&objRegioes.Codigo, &objRegioes.Regiao, &objRegioes.Estados)

		if err != nil {
			return nil, err
		}

		regioes = append(regioes, objRegioes)
	}

	if err = linhas.Err(); err != nil {
		return nil, err
	}

	return regioes, nil
}

// Pesquisa : retorna um conjunto de categorias e regioes.
func Pesquisa() (*MDLPesquisa, error) {
	categorias, err := lstCategoria()
	if err != nil {
		return nil, err
	}

	regioes, err := lstRegiao()
	if err != nil {
		return nil, err
	}

	pesquisa := new(MDLPesquisa)
	pesquisa.MDLRegiao = regioes
	pesquisa.MDLCategoria = categorias

	return pesquisa, nil
}
