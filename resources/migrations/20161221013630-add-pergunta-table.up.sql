CREATE TABLE perguntas
(
    cd_pergunta serial primary key,
    cd_caderno int references caderno(cd_caderno),
    id_questao int not null,
    questao text not null,
    alternativa text[5],
    ativo boolean default true
);