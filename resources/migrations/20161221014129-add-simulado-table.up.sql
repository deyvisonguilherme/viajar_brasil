CREATE TABLE simulado
(
    cd_simulado serial primary key,
    cd_usuario int references usuario(cd_usuario),
    cd_caderno int references caderno(cd_caderno),
    dt_simulado date not null,
    questao int not null,
    resposta int not null
);