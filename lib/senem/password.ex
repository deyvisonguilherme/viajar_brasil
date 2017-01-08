defmodule Senem.Password do
  alias Senem.Repo
  import Ecto.Changeset, only: [put_change: 3]
  import Comeonin.Bcrypt, only: [hashpwsalt: 1]

  def generate_password(changeset) do
    put_change(changeset, :encrypted_password, hashpwsalt(changeset.params["password"]))
  end

  def generate_password_and_store_user(changeset) do
    changeset
    |> generate_password
    |> Repo.insert
  end
end
