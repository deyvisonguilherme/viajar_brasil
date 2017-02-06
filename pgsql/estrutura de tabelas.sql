create table if not exists categorias
(
	id serial primary key,
	ds_categoria varchar(10) not null,
	ativo boolean default True
);

create table if not exists regioes
(
	id serial primary key,
	regiao varchar(8) not null,
	estado char(2) not null,
	ativo boolean default True
);

create table if not exists fotos
(
	id serial primary key,
	legenda varchar(20) not null,
	ativo boolean default True	
);

create table if not exists locais
(
	id serial primary key,
	id_fotos int references fotos(id),
	id_regiao int references regioes (id),
	id_categorias int references categorias (id),
	cidade varchar(20) not null,
	endereco varchar(25) not null,
	ativo boolean default True		
);

create table if not exists usuarios
(
	id serial primary key,
	email varchar(25) not null,
	senha varchar(128) not null,
	ativo boolean default True	
);

create table if not exists dados_usuarios
(
	id serial primary key,
	id_usuario int references usuarios(id),
	nome varchar(40) not null,
	sobrenome varchar(40) not null,
	cidade varchar(30) not null,
	estado char(2) not null,
	nascimento date not null,
	ativo boolean default True	
);

create table if not exists votacao
(
	id serial primary key,
	id_usuario int references usuarios(id),
	id_locais int references locais(id),
	voto int not null,
	ativo boolean default True
);