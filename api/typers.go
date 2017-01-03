package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Caderno struct {
	Cd_caderno   int64  `gorm:"primary_key"`
	Cor          string `gorm:"type:varchar(10); not null"`
	Titulo       string `gorm:"type:varchar(30); not null"`
	Dt_aplicacao `gorm:"type:timestamp; not null"`
}

type Perguntas struct {
	Cd_pergunta int64     `gorm:"primary_key`
	Cd_caderno  []Caderno `gorm:"not null"`
	Id_questao  int64     `gorm:"not null"`
	Questao     text      `gorm:"not null"`
	// alternativa
	Ativo bool
}

type Gabarito struct {
	Cd_gabarito int64 `gorm:"primary_key"`
	Cd_caderno  []Caderno
	Id_questao  int64 `gorm:"not null"`
	Resposta    `gorm:"type:char(1);not null"`
	Ativo       bool
}

type Usuario struct {
	Cd_usuario  int64  `gorm:"primary_key"`
	Nm_completo string `gorm:"type:varchar(45); not null"`
	Nm_curto    string `gorm:"type:varchar(15); not null"`
	Email       string `gorm:"type:varchar(65); not null"`
	Senha       string `gorm:"type:varchar(128);not null"`
	Ativo       bool
}

type dbUsuario struct {
	Cd_dbusuario  int64 `gorm:"primary_key"`
	Cd_usuario    []Usuario
	Escolaridade  string `gorm:"type:varchar(15), not null"`
	En_medio      `gorm:"type:char(7);check(En_medio=Público or En_medio=Privado)"`
	Dt_nascimento `gorm:"type:timestamp; not null"`
	Cidade        string `gorm:"type:varchar(50); not null"`
	Estado        `gorm:"type:char(2); not null"`
}

type Simulado struct {
	Cd_simulado int64 `gorm:"primary_key"`
	Cd_usuario  []Usuario
	Cd_caderno  []Caderno
	Dt_simulado `gorm:"type:timestamp; not null"`
	Questao     int64 `gorm:"not null"`
	Resposta    int64 `gorm:"not null"`
}
