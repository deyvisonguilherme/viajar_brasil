package models

// Votar : grava a votação do usuário.
func Votar(vusuario int, vlocal int, vvotacao int) {
	votar, err := db.Query("")
	if err != nil {
	}
	defer votar.Close()
}

// Votacao : lista a votação do usuário
func Votacao() (int, error) {
	votacao, err := db.Query("")
	if err != nil {
		return 0, err
	}

}

func Comentar() {

}

func Comentarios() {

}

func Page() {

}
