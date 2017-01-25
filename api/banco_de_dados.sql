CREATE TABLE caderno
(
    cd_caderno serial primary key,
    cor varchar(10) not null,
    titulo varchar(30) not null,
    dt_aplicacao date,
    ativo boolean default true
);

CREATE TABLE perguntas
(
    cd_pergunta serial primary key,
    cd_caderno int references caderno(cd_caderno),
    id_questao int not null,
    questao text not null,
    alternativa text[5],
    ativo boolean default true
);

CREATE TABLE gabarito
(
    cd_gabarito serial primary key,
    cd_caderno int references caderno(cd_caderno),
    id_questao int not null,
    resposta char(1) not null,
    ativo boolean default true
);

CREATE TABLE usuario
(
    cd_usuario serial primary key,
    nm_completo varchar(45) not null,
    nm_curto varchar(15) not null,
    email varchar(65) not null,
    senha varchar(128) not null,
    ativo boolean default true
);

CREATE TABLE db_usuario
(
    cd_dbusuario serial primary key,
    cd_usuario int references usuario(cd_usuario),
    escolaridade varchar(15) not null,
    en_medio char(7) check(en_medio='Público' or en_medio='Privado'),
    dt_nascimento date not null,
    cidade varchar(50) not null,
    estado char(2) not null
);

CREATE TABLE simulado
(
    cd_simulado serial primary key,
    cd_usuario int references usuario(cd_usuario),
    cd_caderno int references caderno(cd_caderno),
    dt_simulado date not null,
    questao int not null,
    resposta int not null
);

CREATE TABLE ranker
(
    cd_ranker serial primary key,
    cd_simulado int references simulado(cd_simulado)
);