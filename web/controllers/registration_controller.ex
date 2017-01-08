defmodule Senem.RegistrationController do
  use Senem.Web, :controller

  alias Senem.Password
  plug :scrub_params, "user" when action in [:create]
  plug :action

  def create(conn, %{"user" => user_params}) do
    changeset = User.changeset(%User{}, user_params)
    if changeset.valid? do
      new_user = Password.generate_password_and_store_user(changeset)

      conn
      |> put_flash(:info, "Sucesso, usuário criado com sucesso")
      |> put_session(:current_user, new_user)
      |> redirect(to: page_path(conn, :index))
    else
      render conn, "new.html", changeset: changeset
    end
  end

  def new(conn, _params) do
    changeset = User.changeset(%User{})
    render conn, changeset: changeset
  end
end
