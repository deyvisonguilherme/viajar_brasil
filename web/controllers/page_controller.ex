defmodule Senem.PageController do
  use Senem.Web, :controller

  def index(conn, _params) do
      render conn, "index.html", page_title: "Início"
  end

  def sobre(conn, _params) do
    render conn, "sobre.html", page_title: "Sobre"
  end

  def simulado(conn, _params) do
    render conn, "simulado.html", page_title: "Simulado"
  end

  def caderno(conn, _params) do
    render conn, "caderno.html", page_title: "Caderno"
  end

  def questoes(conn, _params) do
    render conn, "questoes.html", page_title: "Questões"
  end

  def respostas(conn, _params) do
    render conn, "respostas.html", page_title: "Respostas"
  end

  def usuarios(conn, _params) do
    render conn, "usuario.html", page_title: "Usuários"
  end

  def ranker(conn, _params) do
    render conn, "ranker.html", page_title: "Ranker"
  end

  def faleconosco(conn, _params) do
    render conn, "faleconosco.html", page_title: "Fale conosco"
  end

end
