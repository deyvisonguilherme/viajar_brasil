package models

import "fmt"

// Regioes representa um objeto Regioes.
type mdlregiao struct {
	Codigo  int    `json:"codigo"`
	Regiao  string `json:"regiao"`
	Estados *string
}

// ListRegioes return função para busca de regioes.
func ListRegioes() ([]*mdlregiao, error) {
	linhas, err := db.Query("select codigo, regiao, estados from vwRegioes")

	if err != nil {
		return nil, err
	}

	defer linhas.Close()

	regioes := make([]*mdlregiao, 0)

	for linhas.Next() {
		objRegioes := new(mdlregiao)
		err := linhas.Scan(&objRegioes.Codigo, &objRegioes.Regiao, &objRegioes.Estados)

		if err != nil {
			return nil, err
		}

		regioes = append(regioes, objRegioes)
	}
	fmt.Println(regioes)
	if err = linhas.Err(); err != nil {
		return nil, err
	}

	return regioes, nil
}
