-- :name create-Simulated! :! :n
-- :doc create a new Simulated.
INSERT INTO users
(id, first_name, last_name, email, pass)
VALUES (:id, :first_name, :last_name, :email, :pass)

-- :name update-Simulated: :! :
-- :doc update an existing Simulated record.
UPDATE users
SET first_name = :first_name, last_name = :last_name, email = :email
WHERE id = :id

-- :name get-Simulated! :? :1
-- :doc retrieve a Simulated given the id.
SELECT * FROM users
WHERE id = :id

-- :name delete-Simulated! :! :n
-- :doc delete an Simulated given the id.
DELETE FROM users
WHERE id = :id