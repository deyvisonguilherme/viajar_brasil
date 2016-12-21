CREATE TABLE caderno
(
    cd_caderno serial primary key,
    cor varchar(10) not null,
    titulo varchar(30) not null,
    dt_aplicacao date,
    ativo boolean default true
);
