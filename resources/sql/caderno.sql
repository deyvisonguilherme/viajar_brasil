-- :name create-book! :! :n
-- :doc create a new book.
INSERT INTO users
(id, first_name, last_name, email, pass)
VALUES (:id, :first_name, :last_name, :email, :pass)

-- :name update-book: :! :
-- :doc update an existing book record.
UPDATE users
SET first_name = :first_name, last_name = :last_name, email = :email
WHERE id = :id

-- :name get-book! :? :1
-- :doc retrieve a book given the id.
SELECT * FROM users
WHERE id = :id

-- :name delete-book! :! :n
-- :doc delete an book given the id.
DELETE FROM users
WHERE id = :id