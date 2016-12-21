CREATE TABLE usuario
(
    cd_usuario serial primary key,
    nm_completo varchar(45) not null,
    nm_curto varchar(15) not null,
    email varchar(65) not null,
    senha varchar(128) not null,
    ativo boolean default true
);