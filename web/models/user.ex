defmodule Senem.User do
  use Senem.Web, :model

  schema "users" do
    field :username, :string
    field :email, :string
    field :encrypted_password, :string
    field :password, :string ,virtual: true
    field :password_confirmation, :string, virtual: true

    timestamps()
  end

  @required_fields ~w(username email password password_confirmation)
  @optional_fields ~w()

  @doc """
  Builds a changeset based on the `struct` and `params`.
  """
  # def changeset(struct, params \\ %{}) do
  #   struct
  #   |> cast(params, [:username, :email, :encrypted_password])
  #   |> validate_required([:username, :email, :encrypted_password])
  # end

  def changeset(model, params \\ %{}) do
    model
    |> cast(params, @required_fields, @optional_fields)
    |> unique_constraint(:username, on: Senem.Repo, downcase: true)
    |> validate_length(:password, min: 1)
    |> validate_length(:password_confirmation, min: 1)
    |> validate_confirmation(:password)
  end
end
