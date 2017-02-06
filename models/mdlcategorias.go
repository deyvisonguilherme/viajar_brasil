package models

type Categoria struct {
    Id           int
    Ds_categoria string
    Ativo        bool
}

func GetCategorias() ([]*Categoria, error) {
    linhas, err := db.Query("select id, ds_categoria, ativo from categorias where ativo=True")

    if err != nil {
        return nil, err
    }

    defer linhas.Close()

    categorias := make([]*Categoria, 0)
    for linhas.Next() {
        obj_categoria := new(Categoria)

        err := linhas.Scan(&obj_categoria.Id, &obj_categoria.Ds_categoria, &obj_categoria.Ativo)

        if err != nil {
            return nil, err
        }

        categorias = append(categorias, obj_categoria)
    }
    if err = linhas.Err(); err != nil {
        return nil, err
    }

    return categorias, nil
}
