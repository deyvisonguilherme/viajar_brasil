defmodule Senem.PageController do
  use Senem.Web, :controller

  def index(conn, _params) do
    render conn, "index.html"
  end

  def sobre(conn, _params) do
    render conn, "sobre.html"
  end

  def simulado(conn, _params) do
    render conn, "simulado.html"
  end

  def caderno(conn, _params) do
    render conn, "caderno.html"
  end

  def questoes(conn, _params) do
    render conn, "questoes.html"
  end

  def respostas(conn, _params) do
    render conn, "respostas.html"
  end

  def usuarios(conn, _params) do
    render conn, "usuario.html"
  end

  def ranker(conn, _params) do
    render conn, "ranker.html"
  end

  def faleconosco(conn, _params) do
    render conn, "faleconosco.html"
  end

end
