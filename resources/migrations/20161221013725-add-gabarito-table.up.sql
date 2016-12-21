CREATE TABLE gabarito
(
    cd_gabarito serial primary key,
    cd_caderno int references caderno(cd_caderno),
    id_questao int not null,
    resposta char(1) not null,
    ativo boolean default true
);