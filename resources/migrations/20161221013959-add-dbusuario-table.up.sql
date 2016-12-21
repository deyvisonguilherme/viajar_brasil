CREATE TABLE dbusuario
(
    cd_dbusuario serial primary key,
    cd_usuario int references usuario(cd_usuario),
    escolaridade varchar(15) not null,
    en_medio char(7) check(en_medio='Público' or en_medio='privado'),
    dt_nascimento date not null,
    cidade varchar(50) not null,
    estado char(2) not null
);