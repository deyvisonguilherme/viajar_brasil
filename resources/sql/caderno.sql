-- :name create-book! :! :n
-- :doc create a new book.
INSERT INTO caderno 
       (cd_caderno, cor, titulo, dt_aplicacao, ativo)
VALUES (nextval('sequence'), :cor, :titulo, :dt_aplicacao, :ativo)

-- :name update-book: :! :
-- :doc update an existing book record.
UPDATE caderno
SET cor = :cor, titulo = :titulo, dt_aplicacao = :dt_aplicacao, ativo = : ativo
WHERE cd_caderno = :cd_caderno

-- :name get-book! :? :1
-- :doc retrieve a book given the id.
SELECT * 
FROM caderno
WHERE cd_caderno = :cd_caderno

-- :name delete-book! :! :n
-- :doc delete an book given the id.
DELETE 
FROM caderno
WHERE cd_caderno = :cd_caderno