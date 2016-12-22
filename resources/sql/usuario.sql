-- :name create-user! :! :n
-- :doc create a new user.
INSERT INTO users
	   (cd_usuario, nm_completo, nm_curto, email, senha, ativo)
VALUES (nextval("sequence"), :nm_completo, :nm_curto, :email, :senha, :ativo)

-- :name create-dbuser! :! :n
-- :doc create a new dbusuario to user.
INSERT INTO dbusuario
	   (cd_dbusuario, cd_usuario, escolaridade, en_medio, dt_nascimento, cidade, estado)
VALUES (NEXTVAL("SEQUENCE"), :cd_usuario, :escolaridade, :en_medio, :dt_nascimento, :cidade, :estado)

-- :name update-dbuser! :! :name
-- :doc update meta info from usuário.
UPDATE dbusuario
SET cd_usuario = :cd_usuario, escolaridade = :escolaridade, en_medio = :en_medio, dt_nascimento = :dt_nascimento, cidade = :cidade, estado = : estado
WHERE cd_dbusuario = :cd_dbusuario

-- :name update-user: :! :
-- :doc update an existing user record.
UPDATE users
SET nm_completo = :nm_completo, nm_curto = :nm_curto, email = :email, senha = :senha, ativo = :ativo
WHERE cd_usuario = :cd_usuario

-- :name get-user! :? :1
-- :doc retrieve a user given the id.
SELECT * 
FROM users
WHERE cd_usuario = :cd_usuario

-- :name get-user! :? :1
-- :doc retrieve a user given the id.
SELECT * 
FROM users
WHERE cd_dbusuario = :cd_dbusuario

-- :name delete-user! :! :n
-- :doc delete an user given the id.
DELETE 
FROM users
WHERE cd_usuario = :cd_usuario