package models

type Regiao struct{
	Id      int
	Regiao  string
	Estado  string 
	Ativo   bool
}

func GetRegioes()([]*Regiao, error){
	linhas, err := db.Query("select id, regiao, estado, ativo from regioes where ativo = True")
	
    if err != nil {
        return nil, err
    }

	defer linhas.Close()
	
	regioes := make([]*Regiao, 0)
	
	for linhas.Next(){
		obj_regioes := new(Regiao)
		err := linhas.Scan(&obj_regioes.Id,&obj_regioes.Regiao,&obj_regioes.Estado,&obj_regioes.Ativo)	

	    if err != nil {
	        return nil, err
	    }

		regioes = append(regioes, obj_regioes)
	}
	
	if err = linhas.Err(); err != nil {
		return nil, err
	}
		
		return regioes, nil
}