package models

// Regioes representa um objeto Regioes.
type Regioes struct {
	ID     int
	Estado string
	Ativo  bool
}

// GetRegioes função para busca de regioes.
func GetRegioes() ([]*Regioes, error) {
	linhas, err := db.Query("select id, estado, ativo from regioes where ativo = True order by estado")

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	regioes := make([]*Regioes, 0)

	for linhas.Next() {
		objRegioes := new(Regioes)
		err := linhas.Scan(&objRegioes.ID, &objRegioes.Estado, &objRegioes.Ativo)

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
