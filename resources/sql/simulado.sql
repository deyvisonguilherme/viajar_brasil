-- :name create-simulated! :! :n
-- :doc create a new Simulated.
INSERT INTO simulado
	    (cd_simulado, cd_usuario, cd_caderno, dt_simulado, questao, resposta)
VALUES  (nextval("sequence"), :cd_usuario, :cd_caderno, :dt_simulado, :questao, :resposta)

-- name: create-gabarito! :! :n
-- :doc create a new gabarito.
INSERT INTO simulado --< alter gabarito
	    (cd_simulado, cd_usuario, cd_caderno, dt_simulado, questao, resposta)
VALUES  (nextval("sequence"), :cd_usuario, :cd_caderno, :dt_simulado, :questao, :resposta)

-- :name update-simulated: :! :
-- :doc update an existing Simulated record.
UPDATE simulado --< alter gabarito
SET cd_usuario = :cd_usuario, cd_caderno = :cd_caderno, dt_simulado = :dt_simulado, questao = :questao, resposta = : resposta)
WHERE cd_simulado = :cd_simulado

-- :name update-simulated: :! :
-- :doc update an existing Simulated record.
UPDATE simulado
SET cd_usuario = :cd_usuario, cd_caderno = :cd_caderno, dt_simulado = :dt_simulado, questao = :questao, resposta = : resposta)
WHERE cd_simulado = :cd_simulado

-- :name get-simulated! :? :1
-- :doc retrieve a Simulated given the id.
SELECT *  --< alter gabarito
FROM silumado
WHERE cd_simulado = :cd_simulado

-- :name get-simulated! :? :1
-- :doc retrieve a Simulated given the id.
SELECT * 
FROM silumado
WHERE cd_simulado = :cd_simulado

-- :name delete-simulated! :! :n
-- :doc delete an Simulated given the id.
DELETE 
FROM simulado
WHERE cd_simulado = :cd_simulado