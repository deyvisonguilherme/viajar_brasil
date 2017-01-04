defmodule Senem.Router do
  use Senem.Web, :router

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_flash
    plug :protect_from_forgery
    plug :put_secure_browser_headers
  end

  pipeline :api do
    plug :accepts, ["json"]
  end

# router pagina home
  scope "/", Senem do
    pipe_through :browser # Use the default browser stack

    get "/", PageController, :index
    get "/sobre", PageController, :sobre
    get "/simulado", PageController, :simulado
    get "/caderno", PageController, :caderno
    get "/questoes", PageController, :questoes
    get "/respostas", PageController, :respostas
    get "/usuario", PageController, :usuarios
    get "/ranker", PageController, :ranker
    get "/faleconosco", PageController, :faleconosco
  end

  # Other scopes may use custom stacks.
  # scope "/api", Senem do
  #   pipe_through :api
  # end
end
